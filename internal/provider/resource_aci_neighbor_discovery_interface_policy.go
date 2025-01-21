// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/container"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NdIfPolResource{}
var _ resource.ResourceWithImportState = &NdIfPolResource{}

func NewNdIfPolResource() resource.Resource {
	return &NdIfPolResource{}
}

// NdIfPolResource defines the resource implementation.
type NdIfPolResource struct {
	client *client.Client
}

// NdIfPolResourceModel describes the resource data model.
type NdIfPolResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	ParentDn            types.String `tfsdk:"parent_dn"`
	Annotation          types.String `tfsdk:"annotation"`
	Ctrl                types.Set    `tfsdk:"controller_state"`
	Descr               types.String `tfsdk:"description"`
	HopLimit            types.String `tfsdk:"hop_limit"`
	Mtu                 types.String `tfsdk:"mtu"`
	Name                types.String `tfsdk:"name"`
	NameAlias           types.String `tfsdk:"name_alias"`
	NsIntvl             types.String `tfsdk:"neighbor_solicitation_interval"`
	NsRetries           types.String `tfsdk:"neighbor_solicitation_retries"`
	NudRetryBase        types.String `tfsdk:"nud_retry_base"`
	NudRetryInterval    types.String `tfsdk:"nud_retry_interval"`
	NudRetryMaxAttempts types.String `tfsdk:"nud_retry_max_attempts"`
	OwnerKey            types.String `tfsdk:"owner_key"`
	OwnerTag            types.String `tfsdk:"owner_tag"`
	RaIntvl             types.String `tfsdk:"router_advertisement_interval"`
	RaLifetime          types.String `tfsdk:"router_advertisement_lifetime"`
	ReachableTime       types.String `tfsdk:"reachable_time"`
	RetransTimer        types.String `tfsdk:"retransmit_timer"`
	TagAnnotation       types.Set    `tfsdk:"annotations"`
	TagTag              types.Set    `tfsdk:"tags"`
}

func getEmptyNdIfPolResourceModel() *NdIfPolResourceModel {
	return &NdIfPolResourceModel{
		Id:                  basetypes.NewStringNull(),
		ParentDn:            basetypes.NewStringNull(),
		Annotation:          basetypes.NewStringNull(),
		Ctrl:                types.SetNull(types.StringType),
		Descr:               basetypes.NewStringNull(),
		HopLimit:            basetypes.NewStringNull(),
		Mtu:                 basetypes.NewStringNull(),
		Name:                basetypes.NewStringNull(),
		NameAlias:           basetypes.NewStringNull(),
		NsIntvl:             basetypes.NewStringNull(),
		NsRetries:           basetypes.NewStringNull(),
		NudRetryBase:        basetypes.NewStringNull(),
		NudRetryInterval:    basetypes.NewStringNull(),
		NudRetryMaxAttempts: basetypes.NewStringNull(),
		OwnerKey:            basetypes.NewStringNull(),
		OwnerTag:            basetypes.NewStringNull(),
		RaIntvl:             basetypes.NewStringNull(),
		RaLifetime:          basetypes.NewStringNull(),
		ReachableTime:       basetypes.NewStringNull(),
		RetransTimer:        basetypes.NewStringNull(),
		TagAnnotation: types.SetNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key":   types.StringType,
				"value": types.StringType,
			},
		}),
		TagTag: types.SetNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key":   types.StringType,
				"value": types.StringType,
			},
		}),
	}
}

// TagAnnotationNdIfPolResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationNdIfPolResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func getEmptyTagAnnotationNdIfPolResourceModel() TagAnnotationNdIfPolResourceModel {
	return TagAnnotationNdIfPolResourceModel{
		Key:   basetypes.NewStringNull(),
		Value: basetypes.NewStringNull(),
	}
}

var TagAnnotationNdIfPolType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"key":   types.StringType,
		"value": types.StringType,
	},
}

// TagTagNdIfPolResourceModel describes the resource data model for the children without relation ships.
type TagTagNdIfPolResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func getEmptyTagTagNdIfPolResourceModel() TagTagNdIfPolResourceModel {
	return TagTagNdIfPolResourceModel{
		Key:   basetypes.NewStringNull(),
		Value: basetypes.NewStringNull(),
	}
}

var TagTagNdIfPolType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"key":   types.StringType,
		"value": types.StringType,
	},
}

type NdIfPolIdentifier struct {
	Name types.String
}

func (r *NdIfPolResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if !req.Plan.Raw.IsNull() {
		var planData, stateData *NdIfPolResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

		if resp.Diagnostics.HasError() {
			return
		}

		if (planData.Id.IsUnknown() || planData.Id.IsNull()) && !planData.ParentDn.IsUnknown() && !planData.Name.IsUnknown() {
			setNdIfPolId(ctx, planData)
		}

		if stateData == nil && !globalAllowExistingOnCreate && !planData.Id.IsUnknown() && !planData.Id.IsNull() {
			CheckDn(ctx, &resp.Diagnostics, r.client, "ndIfPol", planData.Id.ValueString())
			if resp.Diagnostics.HasError() {
				return
			}
		}

		resp.Diagnostics.Append(resp.Plan.Set(ctx, &planData)...)
	}
}

func (r *NdIfPolResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: aci_neighbor_discovery_interface_policy")
	resp.TypeName = req.ProviderTypeName + "_neighbor_discovery_interface_policy"
	tflog.Debug(ctx, "End metadata of resource: aci_neighbor_discovery_interface_policy")
}

func (r *NdIfPolResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: aci_neighbor_discovery_interface_policy")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The neighbor_discovery_interface_policy resource for the 'ndIfPol' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the Neighbor Discovery Interface Policy object.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"parent_dn": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The distinguished name (DN) of the parent object.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
			},
			"annotation": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Default:             stringdefault.StaticString(globalAnnotation),
				MarkdownDescription: `The annotation of the Neighbor Discovery Interface Policy object.`,
			},
			"controller_state": schema.SetAttribute{
				MarkdownDescription: `The controls for the Neighbor Discovery Interface Policy object.`,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
					SetToSetNullWhenStateIsNullPlanIsUnknownDuringUpdate(nil),
				},
				Validators: []validator.Set{
					setvalidator.SizeAtMost(6),
					setvalidator.ValueStringsAre(
						stringvalidator.OneOf("managed-cfg", "other-cfg", "suppress-ra", "suppress-ra-mtu", "unsolicit-na-glean", "unspecified"),
					),
				},
				ElementType: types.StringType,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The description of the Neighbor Discovery Interface Policy object.`,
			},
			"hop_limit": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The hop limit of the Neighbor Discovery Interface Policy object.`,
			},
			"mtu": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The maximum transmission unit (MTU) of the Neighbor Discovery Interface Policy object.`,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The name of the Neighbor Discovery Interface Policy object.`,
			},
			"name_alias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The name alias of the Neighbor Discovery Interface Policy object.`,
			},
			"neighbor_solicitation_interval": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The interval (milliseconds) for sending neighbor solicitation (NS) messages.`,
			},
			"neighbor_solicitation_retries": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The retransmission count for sending neighbor solicitation (NS) messages.`,
			},
			"nud_retry_base": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The retransmission base value for neighbor unreachability detection (NUD).`,
			},
			"nud_retry_interval": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The retransmission interval (milliseconds) for neighbor unreachability detection (NUD).`,
			},
			"nud_retry_max_attempts": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The maximum number of retransmission attempts for neighbor unreachability detection (NUD).`,
			},
			"owner_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The key for enabling clients to own their data for entity correlation.`,
			},
			"owner_tag": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `A tag for enabling clients to add their own data. For example, to indicate who created this object.`,
			},
			"router_advertisement_interval": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The interval (seconds) for sending route advertisement messages.`,
			},
			"router_advertisement_lifetime": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The default router lifetime (seconds).`,
			},
			"reachable_time": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The time (milliseconds) for which a neighbor is considered reachable after receiving reachability confirmation.`,
			},
			"retransmit_timer": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The retransmission time (milliseconds) for sending neighbor solicitation messages.`,
			},
			"annotations": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
					SetToSetNullWhenStateIsNullPlanIsUnknownDuringUpdate(nil),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
			"tags": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
					SetToSetNullWhenStateIsNullPlanIsUnknownDuringUpdate(nil),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
		},
	}
	tflog.Debug(ctx, "End schema of resource: aci_neighbor_discovery_interface_policy")
}

func (r *NdIfPolResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: aci_neighbor_discovery_interface_policy")
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
	tflog.Debug(ctx, "End configure of resource: aci_neighbor_discovery_interface_policy")
}

func (r *NdIfPolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: aci_neighbor_discovery_interface_policy")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *NdIfPolResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	if stateData.Id.IsUnknown() || stateData.Id.IsNull() {
		setNdIfPolId(ctx, stateData)
	}
	getAndSetNdIfPolAttributes(ctx, &resp.Diagnostics, r.client, stateData)
	if !globalAllowExistingOnCreate && !stateData.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Object Already Exists",
			fmt.Sprintf("The ndIfPol object with DN '%s' already exists.", stateData.Id.ValueString()),
		)
		return
	}

	var data *NdIfPolResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsUnknown() || data.Id.IsNull() {
		setNdIfPolId(ctx, data)
	}

	tflog.Debug(ctx, fmt.Sprintf("Create of resource aci_neighbor_discovery_interface_policy with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationNdIfPolResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagNdIfPolResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getNdIfPolCreateJsonPayload(ctx, &resp.Diagnostics, true, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetNdIfPolAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End create of resource aci_neighbor_discovery_interface_policy with id '%s'", data.Id.ValueString()))
}

func (r *NdIfPolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: aci_neighbor_discovery_interface_policy")
	var data *NdIfPolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource aci_neighbor_discovery_interface_policy with id '%s'", data.Id.ValueString()))

	getAndSetNdIfPolAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *NdIfPolResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, fmt.Sprintf("End read of resource aci_neighbor_discovery_interface_policy with id '%s'", data.Id.ValueString()))
}

func (r *NdIfPolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: aci_neighbor_discovery_interface_policy")
	var data *NdIfPolResourceModel
	var stateData *NdIfPolResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource aci_neighbor_discovery_interface_policy with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationNdIfPolResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagNdIfPolResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getNdIfPolCreateJsonPayload(ctx, &resp.Diagnostics, false, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetNdIfPolAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End update of resource aci_neighbor_discovery_interface_policy with id '%s'", data.Id.ValueString()))
}

func (r *NdIfPolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: aci_neighbor_discovery_interface_policy")
	var data *NdIfPolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource aci_neighbor_discovery_interface_policy with id '%s'", data.Id.ValueString()))
	jsonPayload := GetDeleteJsonPayload(ctx, &resp.Diagnostics, "ndIfPol", data.Id.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}
	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("End delete of resource aci_neighbor_discovery_interface_policy with id '%s'", data.Id.ValueString()))
}

func (r *NdIfPolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: aci_neighbor_discovery_interface_policy")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *NdIfPolResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)
	tflog.Debug(ctx, fmt.Sprintf("Import state of resource aci_neighbor_discovery_interface_policy with id '%s'", stateData.Id.ValueString()))

	tflog.Debug(ctx, "End import of state resource: aci_neighbor_discovery_interface_policy")
}

func getAndSetNdIfPolAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *NdIfPolResourceModel) {
	requestData := DoRestRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=full&rsp-subtree-class=%s", data.Id.ValueString(), "ndIfPol,tagAnnotation,tagTag"), "GET", nil)

	readData := getEmptyNdIfPolResourceModel()

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("ndIfPol").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("ndIfPol").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					readData.Id = basetypes.NewStringValue(attributeValue.(string))
					setNdIfPolParentDn(ctx, attributeValue.(string), readData)
				}
				if attributeName == "annotation" {
					readData.Annotation = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ctrl" {
					ctrlList := make([]string, 0)
					if attributeValue.(string) != "" {
						ctrlList = strings.Split(attributeValue.(string), ",")
					}
					ctrlSet, _ := types.SetValueFrom(ctx, readData.Ctrl.ElementType(ctx), ctrlList)
					readData.Ctrl = ctrlSet
				}
				if attributeName == "descr" {
					readData.Descr = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "hopLimit" {
					readData.HopLimit = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "mtu" {
					readData.Mtu = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "name" {
					readData.Name = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "nameAlias" {
					readData.NameAlias = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "nsIntvl" {
					readData.NsIntvl = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "nsRetries" {
					readData.NsRetries = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "nudRetryBase" {
					readData.NudRetryBase = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "nudRetryInterval" {
					readData.NudRetryInterval = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "nudRetryMaxAttempts" {
					readData.NudRetryMaxAttempts = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ownerKey" {
					readData.OwnerKey = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ownerTag" {
					readData.OwnerTag = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "raIntvl" {
					readData.RaIntvl = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "raLifetime" {
					readData.RaLifetime = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "reachableTime" {
					readData.ReachableTime = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "retransTimer" {
					readData.RetransTimer = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			TagAnnotationNdIfPolList := make([]TagAnnotationNdIfPolResourceModel, 0)
			TagTagNdIfPolList := make([]TagTagNdIfPolResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "tagAnnotation" {
							TagAnnotationNdIfPol := getEmptyTagAnnotationNdIfPolResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationNdIfPol.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationNdIfPol.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}

							}
							TagAnnotationNdIfPolList = append(TagAnnotationNdIfPolList, TagAnnotationNdIfPol)
						}
						if childClassName == "tagTag" {
							TagTagNdIfPol := getEmptyTagTagNdIfPolResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagTagNdIfPol.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagTagNdIfPol.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}

							}
							TagTagNdIfPolList = append(TagTagNdIfPolList, TagTagNdIfPol)
						}
					}
				}
			}
			tagAnnotationSet, _ := types.SetValueFrom(ctx, readData.TagAnnotation.ElementType(ctx), TagAnnotationNdIfPolList)
			readData.TagAnnotation = tagAnnotationSet
			tagTagSet, _ := types.SetValueFrom(ctx, readData.TagTag.ElementType(ctx), TagTagNdIfPolList)
			readData.TagTag = tagTagSet
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'ndIfPol'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	} else {
		readData.Id = basetypes.NewStringNull()
	}
	*data = *readData
}

func getNdIfPolRn(ctx context.Context, data *NdIfPolResourceModel) string {
	return fmt.Sprintf("ndifpol-%s", data.Name.ValueString())
}

func setNdIfPolParentDn(ctx context.Context, dn string, data *NdIfPolResourceModel) {
	bracketIndex := 0
	rnIndex := 0
	for i := len(dn) - 1; i >= 0; i-- {
		if string(dn[i]) == "]" {
			bracketIndex = bracketIndex + 1
		} else if string(dn[i]) == "[" {
			bracketIndex = bracketIndex - 1
		} else if string(dn[i]) == "/" && bracketIndex == 0 {
			rnIndex = i
			break
		}
	}
	data.ParentDn = basetypes.NewStringValue(dn[:rnIndex])
}

func setNdIfPolId(ctx context.Context, data *NdIfPolResourceModel) {
	rn := getNdIfPolRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
}

func getNdIfPolTagAnnotationChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *NdIfPolResourceModel, tagAnnotationNdIfPolPlan, tagAnnotationNdIfPolState []TagAnnotationNdIfPolResourceModel) []map[string]interface{} {
	childPayloads := []map[string]interface{}{}
	if !data.TagAnnotation.IsNull() && !data.TagAnnotation.IsUnknown() {
		tagAnnotationIdentifiers := []TagAnnotationIdentifier{}
		for _, tagAnnotationNdIfPol := range tagAnnotationNdIfPolPlan {
			childMap := NewAciObject()
			if !tagAnnotationNdIfPol.Key.IsNull() && !tagAnnotationNdIfPol.Key.IsUnknown() {
				childMap.Attributes["key"] = tagAnnotationNdIfPol.Key.ValueString()
			}
			if !tagAnnotationNdIfPol.Value.IsNull() && !tagAnnotationNdIfPol.Value.IsUnknown() {
				childMap.Attributes["value"] = tagAnnotationNdIfPol.Value.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"tagAnnotation": childMap})
			tagAnnotationIdentifier := TagAnnotationIdentifier{}
			tagAnnotationIdentifier.Key = tagAnnotationNdIfPol.Key
			tagAnnotationIdentifiers = append(tagAnnotationIdentifiers, tagAnnotationIdentifier)
		}
		for _, tagAnnotation := range tagAnnotationNdIfPolState {
			delete := true
			for _, tagAnnotationIdentifier := range tagAnnotationIdentifiers {
				if tagAnnotationIdentifier.Key == tagAnnotation.Key {
					delete = false
					break
				}
			}
			if delete {
				tagAnnotationChildMapForDelete := NewAciObject()
				tagAnnotationChildMapForDelete.Attributes["status"] = "deleted"
				tagAnnotationChildMapForDelete.Attributes["key"] = tagAnnotation.Key.ValueString()
				childPayloads = append(childPayloads, map[string]interface{}{"tagAnnotation": tagAnnotationChildMapForDelete})
			}
		}
	} else {
		data.TagAnnotation = types.SetNull(data.TagAnnotation.ElementType(ctx))
	}

	return childPayloads
}

func getNdIfPolTagTagChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *NdIfPolResourceModel, tagTagNdIfPolPlan, tagTagNdIfPolState []TagTagNdIfPolResourceModel) []map[string]interface{} {
	childPayloads := []map[string]interface{}{}
	if !data.TagTag.IsNull() && !data.TagTag.IsUnknown() {
		tagTagIdentifiers := []TagTagIdentifier{}
		for _, tagTagNdIfPol := range tagTagNdIfPolPlan {
			childMap := NewAciObject()
			if !tagTagNdIfPol.Key.IsNull() && !tagTagNdIfPol.Key.IsUnknown() {
				childMap.Attributes["key"] = tagTagNdIfPol.Key.ValueString()
			}
			if !tagTagNdIfPol.Value.IsNull() && !tagTagNdIfPol.Value.IsUnknown() {
				childMap.Attributes["value"] = tagTagNdIfPol.Value.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"tagTag": childMap})
			tagTagIdentifier := TagTagIdentifier{}
			tagTagIdentifier.Key = tagTagNdIfPol.Key
			tagTagIdentifiers = append(tagTagIdentifiers, tagTagIdentifier)
		}
		for _, tagTag := range tagTagNdIfPolState {
			delete := true
			for _, tagTagIdentifier := range tagTagIdentifiers {
				if tagTagIdentifier.Key == tagTag.Key {
					delete = false
					break
				}
			}
			if delete {
				tagTagChildMapForDelete := NewAciObject()
				tagTagChildMapForDelete.Attributes["status"] = "deleted"
				tagTagChildMapForDelete.Attributes["key"] = tagTag.Key.ValueString()
				childPayloads = append(childPayloads, map[string]interface{}{"tagTag": tagTagChildMapForDelete})
			}
		}
	} else {
		data.TagTag = types.SetNull(data.TagTag.ElementType(ctx))
	}

	return childPayloads
}

func getNdIfPolCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, createType bool, data *NdIfPolResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationNdIfPolResourceModel, tagTagPlan, tagTagState []TagTagNdIfPolResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}

	if createType && !globalAllowExistingOnCreate {
		payloadMap["attributes"].(map[string]string)["status"] = "created"
	}
	childPayloads := []map[string]interface{}{}

	TagAnnotationchildPayloads := getNdIfPolTagAnnotationChildPayloads(ctx, diags, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)

	TagTagchildPayloads := getNdIfPolTagTagChildPayloads(ctx, diags, data, tagTagPlan, tagTagState)
	if TagTagchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagTagchildPayloads...)

	payloadMap["children"] = childPayloads
	if !data.Annotation.IsNull() && !data.Annotation.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["annotation"] = data.Annotation.ValueString()
	}
	if !data.Ctrl.IsNull() && !data.Ctrl.IsUnknown() {
		var tmpCtrl []string
		data.Ctrl.ElementsAs(ctx, &tmpCtrl, false)
		payloadMap["attributes"].(map[string]string)["ctrl"] = strings.Join(tmpCtrl, ",")
	}
	if !data.Descr.IsNull() && !data.Descr.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["descr"] = data.Descr.ValueString()
	}
	if !data.HopLimit.IsNull() && !data.HopLimit.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["hopLimit"] = data.HopLimit.ValueString()
	}
	if !data.Mtu.IsNull() && !data.Mtu.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["mtu"] = data.Mtu.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["name"] = data.Name.ValueString()
	}
	if !data.NameAlias.IsNull() && !data.NameAlias.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["nameAlias"] = data.NameAlias.ValueString()
	}
	if !data.NsIntvl.IsNull() && !data.NsIntvl.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["nsIntvl"] = data.NsIntvl.ValueString()
	}
	if !data.NsRetries.IsNull() && !data.NsRetries.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["nsRetries"] = data.NsRetries.ValueString()
	}
	if !data.NudRetryBase.IsNull() && !data.NudRetryBase.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["nudRetryBase"] = data.NudRetryBase.ValueString()
	}
	if !data.NudRetryInterval.IsNull() && !data.NudRetryInterval.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["nudRetryInterval"] = data.NudRetryInterval.ValueString()
	}
	if !data.NudRetryMaxAttempts.IsNull() && !data.NudRetryMaxAttempts.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["nudRetryMaxAttempts"] = data.NudRetryMaxAttempts.ValueString()
	}
	if !data.OwnerKey.IsNull() && !data.OwnerKey.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["ownerKey"] = data.OwnerKey.ValueString()
	}
	if !data.OwnerTag.IsNull() && !data.OwnerTag.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["ownerTag"] = data.OwnerTag.ValueString()
	}
	if !data.RaIntvl.IsNull() && !data.RaIntvl.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["raIntvl"] = data.RaIntvl.ValueString()
	}
	if !data.RaLifetime.IsNull() && !data.RaLifetime.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["raLifetime"] = data.RaLifetime.ValueString()
	}
	if !data.ReachableTime.IsNull() && !data.ReachableTime.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["reachableTime"] = data.ReachableTime.ValueString()
	}
	if !data.RetransTimer.IsNull() && !data.RetransTimer.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["retransTimer"] = data.RetransTimer.ValueString()
	}
	payload, err := json.Marshal(map[string]interface{}{"ndIfPol": payloadMap})
	if err != nil {
		diags.AddError(
			"Marshalling of json payload failed",
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	jsonPayload, err := container.ParseJSON(payload)

	if err != nil {
		diags.AddError(
			"Construction of json payload failed",
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}
	return jsonPayload
}
