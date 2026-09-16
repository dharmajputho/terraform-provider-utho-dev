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

// ══════════════════════════════════════════════════════════════
// IPSEC TUNNEL — utho_ipsec
// ══════════════════════════════════════════════════════════════

type IPSecResource struct{ client *client.Client }

type IPSecModel struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	DCSlug       types.String `tfsdk:"dcslug"`
	VPC          types.String `tfsdk:"vpc"`
	CPUModel     types.String `tfsdk:"cpumodel"`
	BillingCycle types.String `tfsdk:"billingcycle"`
	PSK          types.String `tfsdk:"psk"`
	Status       types.String `tfsdk:"status"`
	CreatedAt    types.String `tfsdk:"created_at"`
}

func NewIPSecResource() resource.Resource { return &IPSecResource{} }

func (r *IPSecResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipsec"
}

func (r *IPSecResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Create and manage Utho IPSec site-to-site VPN tunnels.",
		Attributes: map[string]schema.Attribute{
			"id":           schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"psk":          schema.StringAttribute{Computed: true, Sensitive: true, Description: "Pre-shared key for the IPSec tunnel."},
			"status":       schema.StringAttribute{Computed: true, Description: "Tunnel status."},
			"created_at":   schema.StringAttribute{Computed: true, Description: "Creation timestamp."},
			"name":         schema.StringAttribute{Required: true, Description: "Tunnel name.", PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"dcslug":       schema.StringAttribute{Required: true, Description: "Data center slug.", PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"vpc":          schema.StringAttribute{Required: true, Description: "VPC subnet ID.", PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"billingcycle": schema.StringAttribute{Required: true, Description: "Billing cycle: monthly, 3month, 6month, 12month.", PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"cpumodel":     schema.StringAttribute{Optional: true, Description: "CPU model: amd or intel.", PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
		},
	}
}

func (r *IPSecResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *IPSecResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPSecModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id, err := r.client.CreateIPSec(&client.IPSecCreateRequest{
		Name: plan.Name.ValueString(), DCSlug: plan.DCSlug.ValueString(),
		VPC: plan.VPC.ValueString(), CPUModel: plan.CPUModel.ValueString(),
		BillingCycle: plan.BillingCycle.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error creating IPSec tunnel", fmt.Sprintf("%s", err))
		return
	}

	plan.ID = types.StringValue(id)
	plan.PSK = types.StringValue("")
	plan.Status = types.StringValue("active")
	plan.CreatedAt = types.StringValue("")
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *IPSecResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPSecModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ipsec, err := r.client.GetIPSec(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading IPSec tunnel", fmt.Sprintf("%s", err))
		return
	}
	if ipsec == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	state.Name = types.StringValue(ipsec.Name)
	state.DCSlug = types.StringValue(ipsec.DCSlug)
	state.PSK = types.StringValue(ipsec.PSK)
	state.Status = types.StringValue(ipsec.Status)
	state.BillingCycle = types.StringValue(ipsec.BillingCycle)
	state.CreatedAt = types.StringValue(ipsec.CreatedAt)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *IPSecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "IPSec tunnels cannot be updated. Destroy and recreate.")
}

func (r *IPSecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPSecModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.client.DeleteIPSec(state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Error deleting IPSec tunnel", fmt.Sprintf("%s", err))
	}
}

// ══════════════════════════════════════════════════════════════
// IPSEC CONNECTION — utho_ipsec_connection
// ══════════════════════════════════════════════════════════════

type IPSecConnectionResource struct{ client *client.Client }

type IPSecConnectionModel struct {
	ID               types.String `tfsdk:"id"`
	IPSecID          types.String `tfsdk:"ipsec_id"`
	Name             types.String `tfsdk:"name"`
	RemoteIP         types.String `tfsdk:"remote_ip"`
	RemoteLocalIP    types.String `tfsdk:"remote_local_ip"`
	LocalIP          types.String `tfsdk:"local_ip"`
	PSK              types.String `tfsdk:"psk"`
	Phase1Encryption types.String `tfsdk:"phase1_encryption"`
	Phase2Encryption types.String `tfsdk:"phase2_encryption"`
	Phase1Integrity  types.String `tfsdk:"phase1_integrity"`
	Phase2Integrity  types.String `tfsdk:"phase2_integrity"`
	Phase1DHGroup    types.String `tfsdk:"phase1_dh_group"`
	Phase2DHGroup    types.String `tfsdk:"phase2_dh_group"`
	IKEVersion       types.String `tfsdk:"ike_version"`
	Phase1Lifetime   types.String `tfsdk:"phase1_lifetime"`
	Phase2Lifetime   types.String `tfsdk:"phase2_lifetime"`
	RekeyMargin      types.String `tfsdk:"rekey_margin"`
	RekeyFuzz        types.String `tfsdk:"rekey_fuzz"`
	ReplayWindow     types.String `tfsdk:"replay_window"`
	DPDTimeout       types.String `tfsdk:"dpd_timeout"`
	DPDAction        types.String `tfsdk:"dpd_action"`
	StartupAction    types.String `tfsdk:"startup_action"`
	Status           types.String `tfsdk:"status"`
}

func NewIPSecConnectionResource() resource.Resource { return &IPSecConnectionResource{} }

func (r *IPSecConnectionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipsec_connection"
}

func (r *IPSecConnectionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Create and manage connections inside a Utho IPSec VPN tunnel.",
		Attributes: map[string]schema.Attribute{
			"id":                schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"status":            schema.StringAttribute{Computed: true, Description: "Connection status."},
			"ipsec_id":          schema.StringAttribute{Required: true, Description: "IPSec tunnel ID.", PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"name":              schema.StringAttribute{Required: true, Description: "Connection name."},
			"remote_ip":         schema.StringAttribute{Required: true, Description: "Remote peer public IP address."},
			"remote_local_ip":   schema.StringAttribute{Required: true, Description: "Remote network CIDR(s). Comma-separated for multiple."},
			"local_ip":          schema.StringAttribute{Required: true, Description: "Local network CIDR(s). Comma-separated for multiple."},
			"psk":               schema.StringAttribute{Required: true, Sensitive: true, Description: "Pre-shared key for this connection."},
			"phase1_encryption": schema.StringAttribute{Required: true, Description: "Phase 1 encryption algorithm(s) e.g. AES128, AES256. Comma-separated for multiple."},
			"phase2_encryption": schema.StringAttribute{Required: true, Description: "Phase 2 encryption algorithm(s). Comma-separated for multiple."},
			"phase1_integrity":  schema.StringAttribute{Required: true, Description: "Phase 1 integrity algorithm(s) e.g. SHA1, SHA2-256. Comma-separated for multiple."},
			"phase2_integrity":  schema.StringAttribute{Required: true, Description: "Phase 2 integrity algorithm(s). Comma-separated for multiple."},
			"phase1_dh_group":   schema.StringAttribute{Required: true, Description: "Phase 1 DH group(s) e.g. 14, 15. Comma-separated for multiple."},
			"phase2_dh_group":   schema.StringAttribute{Required: true, Description: "Phase 2 DH group(s). Comma-separated for multiple."},
			"ike_version":       schema.StringAttribute{Required: true, Description: "IKE version: ikev1, ikev2, or ikev2,ikev1 for both."},
			"phase1_lifetime":   schema.StringAttribute{Required: true, Description: "Phase 1 lifetime in seconds. Default: 28800."},
			"phase2_lifetime":   schema.StringAttribute{Required: true, Description: "Phase 2 lifetime in seconds. Default: 3600."},
			"rekey_margin":      schema.StringAttribute{Required: true, Description: "Rekey margin in seconds. Default: 270."},
			"rekey_fuzz":        schema.StringAttribute{Required: true, Description: "Rekey fuzz percentage. Default: 100."},
			"replay_window":     schema.StringAttribute{Required: true, Description: "Replay window size. Default: 1024."},
			"dpd_timeout":       schema.StringAttribute{Required: true, Description: "DPD timeout in seconds. Default: 30."},
			"dpd_action":        schema.StringAttribute{Required: true, Description: "DPD action: none, restart, or clear."},
			"startup_action":    schema.StringAttribute{Required: true, Description: "Startup action: start or add."},
		},
	}
}

func (r *IPSecConnectionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *IPSecConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPSecConnectionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	id, err := r.client.CreateIPSecConnection(&client.IPSecConnectionRequest{
		IPSecID: plan.IPSecID.ValueString(), Name: plan.Name.ValueString(),
		RemoteIP: plan.RemoteIP.ValueString(), RemoteLocalIP: plan.RemoteLocalIP.ValueString(),
		LocalIP: plan.LocalIP.ValueString(), PSK: plan.PSK.ValueString(),
		Phase1Encryption: plan.Phase1Encryption.ValueString(), Phase2Encryption: plan.Phase2Encryption.ValueString(),
		Phase1Integrity: plan.Phase1Integrity.ValueString(), Phase2Integrity: plan.Phase2Integrity.ValueString(),
		Phase1DHGroup: plan.Phase1DHGroup.ValueString(), Phase2DHGroup: plan.Phase2DHGroup.ValueString(),
		IKEVersion: plan.IKEVersion.ValueString(), Phase1Lifetime: plan.Phase1Lifetime.ValueString(),
		Phase2Lifetime: plan.Phase2Lifetime.ValueString(), RekeyMargin: plan.RekeyMargin.ValueString(),
		RekeyFuzz: plan.RekeyFuzz.ValueString(), ReplayWindow: plan.ReplayWindow.ValueString(),
		DPDTimeout: plan.DPDTimeout.ValueString(), DPDAction: plan.DPDAction.ValueString(),
		StartupAction: plan.StartupAction.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error creating IPSec connection", fmt.Sprintf("%s", err))
		return
	}

	plan.ID = types.StringValue(id)
	plan.Status = types.StringValue("inactive")
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *IPSecConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPSecConnectionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	connections, err := r.client.ListIPSecConnections(state.IPSecID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading IPSec connections", fmt.Sprintf("%s", err))
		return
	}

	for _, conn := range connections {
		if conn.ID == state.ID.ValueString() {
			state.Name = types.StringValue(conn.Name)
			state.RemoteIP = types.StringValue(conn.RemoteIP)
			state.RemoteLocalIP = types.StringValue(conn.RemoteLocalIP)
			state.LocalIP = types.StringValue(conn.LocalIP)
			state.Status = types.StringValue(conn.Status)
			resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			return
		}
	}
	resp.State.RemoveResource(ctx)
}

func (r *IPSecConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPSecConnectionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.UpdateIPSecConnection(&client.IPSecConnectionRequest{
		IPSecID: plan.IPSecID.ValueString(), ID: plan.ID.ValueString(),
		Name: plan.Name.ValueString(), RemoteIP: plan.RemoteIP.ValueString(),
		RemoteLocalIP: plan.RemoteLocalIP.ValueString(), LocalIP: plan.LocalIP.ValueString(),
		PSK: plan.PSK.ValueString(), Phase1Encryption: plan.Phase1Encryption.ValueString(),
		Phase2Encryption: plan.Phase2Encryption.ValueString(), Phase1Integrity: plan.Phase1Integrity.ValueString(),
		Phase2Integrity: plan.Phase2Integrity.ValueString(), Phase1DHGroup: plan.Phase1DHGroup.ValueString(),
		Phase2DHGroup: plan.Phase2DHGroup.ValueString(), IKEVersion: plan.IKEVersion.ValueString(),
		Phase1Lifetime: plan.Phase1Lifetime.ValueString(), Phase2Lifetime: plan.Phase2Lifetime.ValueString(),
		RekeyMargin: plan.RekeyMargin.ValueString(), RekeyFuzz: plan.RekeyFuzz.ValueString(),
		ReplayWindow: plan.ReplayWindow.ValueString(), DPDTimeout: plan.DPDTimeout.ValueString(),
		DPDAction: plan.DPDAction.ValueString(), StartupAction: plan.StartupAction.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error updating IPSec connection", fmt.Sprintf("%s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *IPSecConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPSecConnectionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.client.DeleteIPSecConnection(state.IPSecID.ValueString(), state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Error deleting IPSec connection", fmt.Sprintf("%s", err))
	}
}
