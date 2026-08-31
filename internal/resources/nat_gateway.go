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

type NATGatewayResource struct{ client *client.Client }
type NATGatewayModel struct {
	ID       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	SubnetID types.String `tfsdk:"subnet_id"`
	PublicIP types.String `tfsdk:"public_ip"`
	DCSlug   types.String `tfsdk:"dcslug"`
}

func NewNATGatewayResource() resource.Resource { return &NATGatewayResource{} }

func (r *NATGatewayResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nat_gateway"
}

func (r *NATGatewayResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Create and manage NAT Gateways for Utho VPC subnets.",
		Attributes: map[string]schema.Attribute{
			"id":        schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":      schema.StringAttribute{Required: true, Description: "NAT Gateway name.", PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"subnet_id": schema.StringAttribute{Required: true, Description: "Subnet ID to attach the NAT Gateway to.", PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"public_ip": schema.StringAttribute{Required: true, Description: "Public IP address for the NAT Gateway.", PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"dcslug":    schema.StringAttribute{Required: true, Description: "Data center slug.", PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
		},
	}
}

func (r *NATGatewayResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *NATGatewayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NATGatewayModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id, err := r.client.CreateNATGateway(&client.NATGatewayCreateRequest{
		Name: plan.Name.ValueString(), Subnet: plan.SubnetID.ValueString(),
		PublicIP: plan.PublicIP.ValueString(), DCSlug: plan.DCSlug.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error creating NAT Gateway", fmt.Sprintf("%s", err))
		return
	}

	plan.ID = types.StringValue(id)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *NATGatewayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *NATGatewayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "NAT Gateways cannot be updated in place.")
}

func (r *NATGatewayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NATGatewayModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Must detach before deleting
	if err := r.client.DetachNATGateway(state.ID.ValueString(), state.SubnetID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Error detaching NAT Gateway", fmt.Sprintf("%s", err))
		return
	}
	if err := r.client.DeleteNATGateway(state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Error deleting NAT Gateway", fmt.Sprintf("%s", err))
	}
}
