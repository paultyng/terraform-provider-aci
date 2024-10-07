// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"context"
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NdIfPolDataSource{}

func NewNdIfPolDataSource() datasource.DataSource {
	return &NdIfPolDataSource{}
}

// NdIfPolDataSource defines the data source implementation.
type NdIfPolDataSource struct {
	client *client.Client
}

func (d *NdIfPolDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of datasource: aci_neighbor_discovery_interface_policy")
	resp.TypeName = req.ProviderTypeName + "_neighbor_discovery_interface_policy"
	tflog.Debug(ctx, "End metadata of datasource: aci_neighbor_discovery_interface_policy")
}

func (d *NdIfPolDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of datasource: aci_neighbor_discovery_interface_policy")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The neighbor_discovery_interface_policy datasource for the 'ndIfPol' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the Neighbor Discovery Interface Policy object.",
			},
			"parent_dn": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The distinguished name (DN) of the parent object.",
			},
			"annotation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The annotation of the Neighbor Discovery Interface Policy object.`,
			},
			"controller_state": schema.SetAttribute{
				Computed:            true,
				MarkdownDescription: `The controls for the Neighbor Discovery Interface Policy object.`,
				ElementType:         types.StringType,
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The description of the Neighbor Discovery Interface Policy object.`,
			},
			"hop_limit": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The hop limit of the Neighbor Discovery Interface Policy object.`,
			},
			"mtu": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The maximum transmission unit (MTU) of the Neighbor Discovery Interface Policy object.`,
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: `The name of the Neighbor Discovery Interface Policy object.`,
			},
			"name_alias": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The name alias of the Neighbor Discovery Interface Policy object.`,
			},
			"neighbor_solicitation_interval": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The interval (milliseconds) for sending neighbor solicitation (NS) messages.`,
			},
			"neighbor_solicitation_retries": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The retransmission count for sending neighbor solicitation (NS) messages.`,
			},
			"nud_retry_base": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The retransmission base value for neighbor unreachability detection (NUD).`,
			},
			"nud_retry_interval": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The retransmission interval (milliseconds) for neighbor unreachability detection (NUD).`,
			},
			"nud_retry_max_attempts": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The maximum number of retransmission attempts for neighbor unreachability detection (NUD).`,
			},
			"owner_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The key for enabling clients to own their data for entity correlation.`,
			},
			"owner_tag": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `A tag for enabling clients to add their own data. For example, to indicate who created this object.`,
			},
			"router_advertisement_interval": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The interval (seconds) for sending route advertisement messages.`,
			},
			"router_advertisement_lifetime": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The default router lifetime (seconds).`,
			},
			"reachable_time": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The time (milliseconds) for which a neighbor is considered reachable after receiving reachability confirmation.`,
			},
			"retransmit_timer": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The retransmission time (milliseconds) for sending neighbor solicitation messages.`,
			},
			"annotations": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
			"tags": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
		},
	}
	tflog.Debug(ctx, "End schema of datasource: aci_neighbor_discovery_interface_policy")
}

func (d *NdIfPolDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of datasource: aci_neighbor_discovery_interface_policy")
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
	tflog.Debug(ctx, "End configure of datasource: aci_neighbor_discovery_interface_policy")
}

func (d *NdIfPolDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Start read of datasource: aci_neighbor_discovery_interface_policy")
	var data *NdIfPolResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setNdIfPolId(ctx, data)

	// Create a copy of the Id for when not found during getAndSetNdIfPolAttributes
	cachedId := data.Id.ValueString()

	tflog.Debug(ctx, fmt.Sprintf("Read of datasource aci_neighbor_discovery_interface_policy with id '%s'", data.Id.ValueString()))

	getAndSetNdIfPolAttributes(ctx, &resp.Diagnostics, d.client, data)

	if data.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Failed to read aci_neighbor_discovery_interface_policy data source",
			fmt.Sprintf("The aci_neighbor_discovery_interface_policy data source with id '%s' has not been found", cachedId),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End read of datasource aci_neighbor_discovery_interface_policy with id '%s'", data.Id.ValueString()))
}
