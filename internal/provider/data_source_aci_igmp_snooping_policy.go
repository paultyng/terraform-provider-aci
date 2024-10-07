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
var _ datasource.DataSource = &IgmpSnoopPolDataSource{}

func NewIgmpSnoopPolDataSource() datasource.DataSource {
	return &IgmpSnoopPolDataSource{}
}

// IgmpSnoopPolDataSource defines the data source implementation.
type IgmpSnoopPolDataSource struct {
	client *client.Client
}

func (d *IgmpSnoopPolDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of datasource: aci_igmp_snooping_policy")
	resp.TypeName = req.ProviderTypeName + "_igmp_snooping_policy"
	tflog.Debug(ctx, "End metadata of datasource: aci_igmp_snooping_policy")
}

func (d *IgmpSnoopPolDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of datasource: aci_igmp_snooping_policy")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The igmp_snooping_policy datasource for the 'igmpSnoopPol' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the IGMP Snooping Policy object.",
			},
			"parent_dn": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The distinguished name (DN) of the parent object.",
			},
			"admin_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The administrative state of the IGMP Snooping Policy object.`,
			},
			"annotation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The annotation of the IGMP Snooping Policy object.`,
			},
			"control": schema.SetAttribute{
				Computed:            true,
				MarkdownDescription: `The controls for the IGMP Snooping Policy object.`,
				ElementType:         types.StringType,
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The description of the IGMP Snooping Policy object.`,
			},
			"last_member_interval": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The last member interval (seconds) of the IGMP Snooping Policy object. The group state is removed when no host responds before the timeout.`,
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: `The name of the IGMP Snooping Policy object.`,
			},
			"name_alias": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The name alias of the IGMP Snooping Policy object.`,
			},
			"owner_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The key for enabling clients to own their data for entity correlation.`,
			},
			"owner_tag": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `A tag for enabling clients to add their own data. For example, to indicate who created this object.`,
			},
			"query_interval": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The query interval (seconds) of the IGMP Snooping Policy object.`,
			},
			"response_interval": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The response interval (seconds) of the IGMP Snooping Policy object.`,
			},
			"start_query_count": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The start query count of the IGMP Snooping Policy object.`,
			},
			"start_query_interval": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The query interval (seconds) of the IGMP Snooping Policy object at start-up.`,
			},
			"querier_version": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `Version.`,
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
	tflog.Debug(ctx, "End schema of datasource: aci_igmp_snooping_policy")
}

func (d *IgmpSnoopPolDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of datasource: aci_igmp_snooping_policy")
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
	tflog.Debug(ctx, "End configure of datasource: aci_igmp_snooping_policy")
}

func (d *IgmpSnoopPolDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Start read of datasource: aci_igmp_snooping_policy")
	var data *IgmpSnoopPolResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setIgmpSnoopPolId(ctx, data)

	// Create a copy of the Id for when not found during getAndSetIgmpSnoopPolAttributes
	cachedId := data.Id.ValueString()

	tflog.Debug(ctx, fmt.Sprintf("Read of datasource aci_igmp_snooping_policy with id '%s'", data.Id.ValueString()))

	getAndSetIgmpSnoopPolAttributes(ctx, &resp.Diagnostics, d.client, data)

	if data.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Failed to read aci_igmp_snooping_policy data source",
			fmt.Sprintf("The aci_igmp_snooping_policy data source with id '%s' has not been found", cachedId),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End read of datasource aci_igmp_snooping_policy with id '%s'", data.Id.ValueString()))
}
