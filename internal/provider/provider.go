package provider

import (
	"context"
	"os"

	"github.com/dharmajputho/terraform-provider-utho/internal/client"
	"github.com/dharmajputho/terraform-provider-utho/internal/datasources"
	"github.com/dharmajputho/terraform-provider-utho/internal/resources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type uthoProvider struct{}

type uthoProviderModel struct {
	APIKey types.String `tfsdk:"api_key"`
}

func New() func() provider.Provider {
	return func() provider.Provider {
		return &uthoProvider{}
	}
}

func (p *uthoProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "utho"
	resp.Version = "1.0.0"
}

func (p *uthoProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with Utho Cloud resources.",
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				Required:    true,
				Sensitive:   true,
				Description: "Utho API key. Can also be set via UTHO_API_KEY environment variable.",
			},
		},
	}
}

func (p *uthoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config uthoProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiKey := os.Getenv("UTHO_API_KEY")
	if !config.APIKey.IsNull() && config.APIKey.ValueString() != "" {
		apiKey = config.APIKey.ValueString()
	}

	if apiKey == "" {
		resp.Diagnostics.AddError(
			"Missing API Key",
			"Set api_key in provider block or export UTHO_API_KEY=your-key",
		)
		return
	}

	uthoClient := client.NewClient(apiKey)
	resp.DataSourceData = uthoClient
	resp.ResourceData = uthoClient
}

func (p *uthoProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewCloudResource,         // utho_cloud
		resources.NewCloudPowerResource,    // utho_cloud_power
		resources.NewCloudFirewallResource, //utho_cloud_firewall
		resources.NewCloudStorageResource,  // utho_cloud_storage
		resources.NewCloudEBSResource,      // utho_cloud_ebs
		resources.NewCloudSnapshotResource, // utho_cloud_snapshot
		resources.NewCloudISOResource,      // utho_cloud_iso
		resources.NewCloudResizeResource,   // utho_cloud_resize
		resources.NewCloudPublicIPResource, // utho_cloud_public_ip
		resources.NewCloudVPCResource,      // utho_cloud_vpc
		resources.NewVPCResource,           // utho_vpc
		resources.NewSubnetResource,        // utho_subnet
		resources.NewNATGatewayResource,    // utho_nat_gateway
		resources.NewRouteTableResource,    // utho_route_table
		resources.NewRouteResource,         // utho_route
		resources.NewElasticIPResource,     // utho_elastic_ip
		resources.NewVPCPeeringResource,    // utho_vpc_peering

	}
}

func (p *uthoProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datasources.NewCloudsDataSource, // data.utho_clouds
	}
}
