// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"context"
	"fmt"

	customTypes "github.com/CiscoDevNet/terraform-provider-aci/v2/internal/custom_types"
	"github.com/ciscoecosystem/aci-go-client/v2/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &QosDot1PClassDataSource{}

func NewQosDot1PClassDataSource() datasource.DataSource {
	return &QosDot1PClassDataSource{}
}

// QosDot1PClassDataSource defines the data source implementation.
type QosDot1PClassDataSource struct {
	client *client.Client
}

func (d *QosDot1PClassDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of datasource: aci_dot1p_classifier")
	resp.TypeName = req.ProviderTypeName + "_dot1p_classifier"
	tflog.Debug(ctx, "End metadata of datasource: aci_dot1p_classifier")
}

func (d *QosDot1PClassDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of datasource: aci_dot1p_classifier")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The dot1p_classifier datasource for the 'qosDot1PClass' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the Dot1p Classifier object.",
			},
			"parent_dn": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The distinguished name (DN) of the parent object.",
			},
			"annotation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The annotation of the Dot1p Classifier object.`,
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The description of the Dot1p Classifier object.`,
			},
			"from": schema.StringAttribute{
				CustomType:          customTypes.QosDot1PClassFromStringType{},
				Required:            true,
				MarkdownDescription: `The Dot1p priority range starting value.`,
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The name of the Dot1p Classifier object.`,
			},
			"name_alias": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The name alias of the Dot1p Classifier object.`,
			},
			"priority": schema.StringAttribute{
				CustomType:          customTypes.QosDot1PClassPrioStringType{},
				Computed:            true,
				MarkdownDescription: `The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.`,
			},
			"target": schema.StringAttribute{
				CustomType:          customTypes.QosDot1PClassTargetStringType{},
				Computed:            true,
				MarkdownDescription: `The target of the Dot1p Classifier object. This Fabric only supports DSCP mutation, Dot1P mutation is not supported.`,
			},
			"target_cos": schema.StringAttribute{
				CustomType:          customTypes.QosDot1PClassTargetCosStringType{},
				Computed:            true,
				MarkdownDescription: `Target COS to be driven based on the range of input values of DSCP coming into the fabric.`,
			},
			"to": schema.StringAttribute{
				CustomType:          customTypes.QosDot1PClassToStringType{},
				Required:            true,
				MarkdownDescription: `The Dot1p priority range ending value.`,
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
	tflog.Debug(ctx, "End schema of datasource: aci_dot1p_classifier")
}

func (d *QosDot1PClassDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of datasource: aci_dot1p_classifier")
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
	tflog.Debug(ctx, "End configure of datasource: aci_dot1p_classifier")
}

func (d *QosDot1PClassDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Start read of datasource: aci_dot1p_classifier")
	var data *QosDot1PClassResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setQosDot1PClassId(ctx, data)

	// Create a copy of the Id for when not found during getAndSetQosDot1PClassAttributes
	cachedId := data.Id.ValueString()

	tflog.Debug(ctx, fmt.Sprintf("Read of datasource aci_dot1p_classifier with id '%s'", data.Id.ValueString()))

	getAndSetQosDot1PClassAttributes(ctx, &resp.Diagnostics, d.client, data)

	if data.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Failed to read aci_dot1p_classifier data source",
			fmt.Sprintf("The aci_dot1p_classifier data source with id '%s' has not been found", cachedId),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End read of datasource aci_dot1p_classifier with id '%s'", data.Id.ValueString()))
}