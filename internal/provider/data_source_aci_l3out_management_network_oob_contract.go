// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MgmtRsOoBConsDataSource{}

func NewMgmtRsOoBConsDataSource() datasource.DataSource {
	return &MgmtRsOoBConsDataSource{}
}

// MgmtRsOoBConsDataSource defines the data source implementation.
type MgmtRsOoBConsDataSource struct {
	client *client.Client
}

func (d *MgmtRsOoBConsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Trace(ctx, "Start metadata of datasource: aci_l3out_management_network_oob_contract")
	resp.TypeName = req.ProviderTypeName + "_l3out_management_network_oob_contract"
	tflog.Trace(ctx, "End metadata of datasource: aci_l3out_management_network_oob_contract")
}

func (d *MgmtRsOoBConsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Trace(ctx, "Start schema of datasource: aci_l3out_management_network_oob_contract")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The l3out_management_network_oob_contract datasource for the 'mgmtRsOoBCons' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the L3out Management Network Oob Contract object.",
			},
			"parent_dn": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The distinguished name (DN) of the parent object.",
			},
			"annotation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The annotation of the L3out Management Network Oob Contract object.`,
			},
			"priority": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The Quality of service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.`,
			},
			"out_of_band_contract_name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: `An out-of-band management endpoint group contract consists of switches (leaves/spines) and APICs that are part of the associated out-of-band management zone. Each node in the group is assigned an IP address that is dynamically allocated from the address pool associated with the corresponding out-of-band management zone.`,
			},
			"annotations": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The key or password used to uniquely identify this configuration object.`,
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
	tflog.Trace(ctx, "End schema of datasource: aci_l3out_management_network_oob_contract")
}

func (d *MgmtRsOoBConsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Trace(ctx, "Start configure of datasource: aci_l3out_management_network_oob_contract")
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
	tflog.Trace(ctx, "End configure of datasource: aci_l3out_management_network_oob_contract")
}

func (d *MgmtRsOoBConsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Trace(ctx, "Start read of datasource: aci_l3out_management_network_oob_contract")
	var data *MgmtRsOoBConsResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setMgmtRsOoBConsId(ctx, data)

	tflog.Trace(ctx, fmt.Sprintf("Read of datasource aci_l3out_management_network_oob_contract with id '%s'", data.Id.ValueString()))

	messageMap := setMgmtRsOoBConsAttributes(ctx, d.client, data)
	if messageMap != nil {
		resp.Diagnostics.AddError(messageMap.(map[string]string)["message"], messageMap.(map[string]string)["messageDetail"])
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "End read of datasource: aci_l3out_management_network_oob_contract")
}
