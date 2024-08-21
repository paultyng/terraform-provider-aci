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
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &L3extRsLblToInstPDataSource{}

func NewL3extRsLblToInstPDataSource() datasource.DataSource {
	return &L3extRsLblToInstPDataSource{}
}

// L3extRsLblToInstPDataSource defines the data source implementation.
type L3extRsLblToInstPDataSource struct {
	client *client.Client
}

func (d *L3extRsLblToInstPDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of datasource: aci_relation_to_external_network_instance_profile")
	resp.TypeName = req.ProviderTypeName + "_relation_to_external_network_instance_profile"
	tflog.Debug(ctx, "End metadata of datasource: aci_relation_to_external_network_instance_profile")
}

func (d *L3extRsLblToInstPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of datasource: aci_relation_to_external_network_instance_profile")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The relation_to_external_network_instance_profile datasource for the 'l3extRsLblToInstP' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the Relation To External Network Instance Profile object.",
			},
			"parent_dn": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The distinguished name (DN) of the parent object.",
			},
			"annotation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The annotation of the Relation To External Network Instance Profile object.`,
			},
			"target_dn": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: `The distinguished name (DN) of the External Network Instance Profile object.`,
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
	tflog.Debug(ctx, "End schema of datasource: aci_relation_to_external_network_instance_profile")
}

func (d *L3extRsLblToInstPDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of datasource: aci_relation_to_external_network_instance_profile")
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
	tflog.Debug(ctx, "End configure of datasource: aci_relation_to_external_network_instance_profile")
}

func (d *L3extRsLblToInstPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Start read of datasource: aci_relation_to_external_network_instance_profile")
	var data *L3extRsLblToInstPResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setL3extRsLblToInstPId(ctx, data)

	// Create a copy of the Id for when not found during getAndSetL3extRsLblToInstPAttributes
	cachedId := data.Id.ValueString()

	tflog.Debug(ctx, fmt.Sprintf("Read of datasource aci_relation_to_external_network_instance_profile with id '%s'", data.Id.ValueString()))

	getAndSetL3extRsLblToInstPAttributes(ctx, &resp.Diagnostics, d.client, data)

	if data.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Failed to read aci_relation_to_external_network_instance_profile data source",
			fmt.Sprintf("The aci_relation_to_external_network_instance_profile data source with id '%s' has not been found", cachedId),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End read of datasource aci_relation_to_external_network_instance_profile with id '%s'", data.Id.ValueString()))
}
