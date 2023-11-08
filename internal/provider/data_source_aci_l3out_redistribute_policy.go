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
var _ datasource.DataSource = &L3extRsRedistributePolDataSource{}

func NewL3extRsRedistributePolDataSource() datasource.DataSource {
	return &L3extRsRedistributePolDataSource{}
}

// L3extRsRedistributePolDataSource defines the data source implementation.
type L3extRsRedistributePolDataSource struct {
	client *client.Client
}

func (d *L3extRsRedistributePolDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Trace(ctx, "Start metadata of datasource: aci_l3out_redistribute_policy")
	resp.TypeName = req.ProviderTypeName + "_l3out_redistribute_policy"
	tflog.Trace(ctx, "End metadata of datasource: aci_l3out_redistribute_policy")
}

func (d *L3extRsRedistributePolDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Trace(ctx, "Start schema of datasource: aci_l3out_redistribute_policy")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The l3out_redistribute_policy datasource for the 'l3extRsRedistributePol' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the L3out Redistribute Policy object.",
			},
			"parent_dn": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The distinguished name (DN) of the parent object.",
			},
			"annotation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The annotation of the L3out Redistribute Policy object.`,
			},
			"src": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: `The source IP address.`,
			},
			"route_control_profile_name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: `The name of the route profile associated with this object.`,
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
	tflog.Trace(ctx, "End schema of datasource: aci_l3out_redistribute_policy")
}

func (d *L3extRsRedistributePolDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Trace(ctx, "Start configure of datasource: aci_l3out_redistribute_policy")
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		resp.Diagnostics.AddError(
			"Provider has not been configured",
			"The req.ProviderData is nil. Please report this issue to the provider developers.",
		)
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
	tflog.Trace(ctx, "End configure of datasource: aci_l3out_redistribute_policy")
}

func (d *L3extRsRedistributePolDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Trace(ctx, "Start read of datasource: aci_l3out_redistribute_policy")
	var data *L3extRsRedistributePolResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setL3extRsRedistributePolId(ctx, data)

	tflog.Trace(ctx, fmt.Sprintf("Read of datasource aci_l3out_redistribute_policy with id '%s'", data.Id.ValueString()))

	setL3extRsRedistributePolAttributes(ctx, &resp.Diagnostics, d.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "End read of datasource: aci_l3out_redistribute_policy")
}
