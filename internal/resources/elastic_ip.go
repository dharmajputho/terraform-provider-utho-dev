package resources

import (
	"context"
	"fmt"

	"github.com/dharmajputho/terraform-provider-utho/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ElasticIPResource struct{ client *client.Client }
type ElasticIPModel struct {
	ID           types.String `tfsdk:"id"`
	IP           types.String `tfsdk:"ip"`
	DCSlug       types.String `tfsdk:"dcslug"`
	BillingCycle types.String `tfsdk:"billingcycle"`
}

func NewElasticIPResource() resource.Resource { return &ElasticIPResource{} }

func (r *ElasticIPResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_elastic_ip"
}

func (r *ElasticIPResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Allocate and manage Elastic IPs on Utho Cloud.",
		Attributes: map[string]schema.Attribute{
			"id":           schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"ip":           schema.StringAttribute{Computed: true, Description: "Allocated elastic IP address."},
			"dcslug":       schema.StringAttribute{Required: true, Description: "Data center slug.", PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"billingcycle": schema.StringAttribute{Required: true, Description: "Billing cycle: monthly or hourly.", PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
		},
	}
}

func (r *ElasticIPResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Provider Data", "Expected *client.Client")
		return
	}
	r.client = c
}

func (r *ElasticIPResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ElasticIPModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	eip, err := r.client.AllocateElasticIP(&client.ElasticIPAllocateRequest{
		DCSlug: plan.DCSlug.ValueString(), BillingCycle: plan.BillingCycle.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error allocating Elastic IP", fmt.Sprintf("%s", err))
		return
	}

	plan.IP = types.StringValue(eip.IP)
	plan.ID = types.StringValue(eip.IP)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *ElasticIPResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *ElasticIPResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "Elastic IPs cannot be updated. Destroy and reallocate.")
}

func (r *ElasticIPResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ElasticIPModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.client.ReleaseElasticIP(state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Error releasing Elastic IP", fmt.Sprintf("%s", err))
	}
}
