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
var _ datasource.DataSource = &FhsBDPolDataSource{}

func NewFhsBDPolDataSource() datasource.DataSource {
	return &FhsBDPolDataSource{}
}

// FhsBDPolDataSource defines the data source implementation.
type FhsBDPolDataSource struct {
	client *client.Client
}

func (d *FhsBDPolDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of datasource: aci_first_hop_security_policy")
	resp.TypeName = req.ProviderTypeName + "_first_hop_security_policy"
	tflog.Debug(ctx, "End metadata of datasource: aci_first_hop_security_policy")
}

func (d *FhsBDPolDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of datasource: aci_first_hop_security_policy")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The first_hop_security_policy datasource for the 'fhsBDPol' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the First Hop Security Policy object.",
			},
			"parent_dn": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The distinguished name (DN) of the parent object.",
			},
			"annotation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The annotation of the First Hop Security Policy object.`,
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The description of the First Hop Security Policy object.`,
			},
			"ip_inspection": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The Inspection Status for IPv4 and IPv6 traffic of the First Hop Security Policy object.`,
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: `The name of the First Hop Security Policy object.`,
			},
			"name_alias": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The name alias of the First Hop Security Policy object.`,
			},
			"owner_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The key for enabling clients to own their data for entity correlation.`,
			},
			"owner_tag": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `A tag for enabling clients to add their own data. For example, to indicate who created this object.`,
			},
			"router_advertisement": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `Enable Router Advertisement Guard for the First Hop Security Policy object.`,
			},
			"source_guard": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The Source Guard Status for IPv4 and IPv6 traffic of the First Hop Security Policy object.`,
			},
			"route_advertisement_guard_policy": schema.SingleNestedAttribute{
				MarkdownDescription: ``,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"annotation": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The annotation of the Route Advertisement Guard Policy object.`,
					},
					"description": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The description of the Route Advertisement Guard Policy object.`,
					},
					"managed_config_check": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `Perform a managed configuration check for the Route Advertisement Guard Policy object.`,
					},
					"managed_config_flag": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The managed configuration flag setting for the Route Advertisement Guard Policy object.`,
					},
					"max_hop_limit": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The maximum hop limit for the Route Advertisement Guard Policy object.`,
					},
					"max_router_preference": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The allowed maximum router preference for the Route Advertisement Guard Policy object.`,
					},
					"min_hop_limit": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The minimum hop limit for the Route Advertisement Guard Policy object.`,
					},
					"name": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The name of the Route Advertisement Guard Policy object.`,
					},
					"name_alias": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The name alias of the Route Advertisement Guard Policy object.`,
					},
					"other_config_check": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `Perform other configuration checks for the Route Advertisement Guard Policy object.`,
					},
					"other_config_flag": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The other configuration flag setting for the Route Advertisement Guard Policy object.`,
					},
					"owner_key": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `The key for enabling clients to own their data for entity correlation.`,
					},
					"owner_tag": schema.StringAttribute{
						Computed:            true,
						MarkdownDescription: `A tag for enabling clients to add their own data. For example, to indicate who created this object.`,
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
	tflog.Debug(ctx, "End schema of datasource: aci_first_hop_security_policy")
}

func (d *FhsBDPolDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of datasource: aci_first_hop_security_policy")
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
	tflog.Debug(ctx, "End configure of datasource: aci_first_hop_security_policy")
}

func (d *FhsBDPolDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Start read of datasource: aci_first_hop_security_policy")
	var data *FhsBDPolResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setFhsBDPolId(ctx, data)

	// Create a copy of the Id for when not found during getAndSetFhsBDPolAttributes
	cachedId := data.Id.ValueString()

	tflog.Debug(ctx, fmt.Sprintf("Read of datasource aci_first_hop_security_policy with id '%s'", data.Id.ValueString()))

	getAndSetFhsBDPolAttributes(ctx, &resp.Diagnostics, d.client, data)

	if data.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Failed to read aci_first_hop_security_policy data source",
			fmt.Sprintf("The aci_first_hop_security_policy data source with id '%s' has not been found", cachedId),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End read of datasource aci_first_hop_security_policy with id '%s'", data.Id.ValueString()))
}
