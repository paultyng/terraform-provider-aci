// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/container"
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
var _ resource.Resource = &RtctrlProfileResource{}
var _ resource.ResourceWithImportState = &RtctrlProfileResource{}

func NewRtctrlProfileResource() resource.Resource {
	return &RtctrlProfileResource{}
}

// RtctrlProfileResource defines the resource implementation.
type RtctrlProfileResource struct {
	client *client.Client
}

// RtctrlProfileResourceModel describes the resource data model.
type RtctrlProfileResourceModel struct {
	Id            types.String `tfsdk:"id"`
	ParentDn      types.String `tfsdk:"parent_dn"`
	Annotation    types.String `tfsdk:"annotation"`
	AutoContinue  types.String `tfsdk:"route_map_continue"`
	Descr         types.String `tfsdk:"description"`
	Name          types.String `tfsdk:"name"`
	NameAlias     types.String `tfsdk:"name_alias"`
	OwnerKey      types.String `tfsdk:"owner_key"`
	OwnerTag      types.String `tfsdk:"owner_tag"`
	Type          types.String `tfsdk:"route_control_profile_type"`
	TagAnnotation types.Set    `tfsdk:"annotations"`
	TagTag        types.Set    `tfsdk:"tags"`
}

func getEmptyRtctrlProfileResourceModel() *RtctrlProfileResourceModel {
	return &RtctrlProfileResourceModel{
		Id:           basetypes.NewStringNull(),
		ParentDn:     basetypes.NewStringNull(),
		Annotation:   basetypes.NewStringNull(),
		AutoContinue: basetypes.NewStringNull(),
		Descr:        basetypes.NewStringNull(),
		Name:         basetypes.NewStringNull(),
		NameAlias:    basetypes.NewStringNull(),
		OwnerKey:     basetypes.NewStringNull(),
		OwnerTag:     basetypes.NewStringNull(),
		Type:         basetypes.NewStringNull(),
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

// TagAnnotationRtctrlProfileResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationRtctrlProfileResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func getEmptyTagAnnotationRtctrlProfileResourceModel() TagAnnotationRtctrlProfileResourceModel {
	return TagAnnotationRtctrlProfileResourceModel{
		Key:   basetypes.NewStringNull(),
		Value: basetypes.NewStringNull(),
	}
}

// TagTagRtctrlProfileResourceModel describes the resource data model for the children without relation ships.
type TagTagRtctrlProfileResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func getEmptyTagTagRtctrlProfileResourceModel() TagTagRtctrlProfileResourceModel {
	return TagTagRtctrlProfileResourceModel{
		Key:   basetypes.NewStringNull(),
		Value: basetypes.NewStringNull(),
	}
}

type RtctrlProfileIdentifier struct {
	Name types.String
}

func (r *RtctrlProfileResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if !req.Plan.Raw.IsNull() {
		var planData, stateData *RtctrlProfileResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

		if resp.Diagnostics.HasError() {
			return
		}

		if (planData.Id.IsUnknown() || planData.Id.IsNull()) && !planData.ParentDn.IsUnknown() && !planData.Name.IsUnknown() {
			SetRtctrlProfileId(ctx, planData)
		}

		if stateData == nil && !globalAllowExistingOnCreate && !planData.Id.IsUnknown() && !planData.Id.IsNull() {
			CheckDn(ctx, &resp.Diagnostics, r.client, "rtctrlProfile", planData.Id.ValueString())
			if resp.Diagnostics.HasError() {
				return
			}
		}

		resp.Diagnostics.Append(resp.Plan.Set(ctx, &planData)...)
	}
}

func (r *RtctrlProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: aci_route_control_profile")
	resp.TypeName = req.ProviderTypeName + "_route_control_profile"
	tflog.Debug(ctx, "End metadata of resource: aci_route_control_profile")
}

func (r *RtctrlProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: aci_route_control_profile")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The route_control_profile resource for the 'rtctrlProfile' class",
		Version:             1,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the Route Control Profile object.",
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
				MarkdownDescription: `The annotation of the Route Control Profile object.`,
			},
			"route_map_continue": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("no", "yes"),
				},
				MarkdownDescription: `The route_map_continue applies a continue statement for all user-configured sequences (contexts) in a given BGP route profile so that switches will continue to find matches with subsequent user-configured sequences in the route map. Without the continue statement, after a route matches one of the sequences in a route map, switches will not continue to check other sequences.`,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The description of the Route Control Profile object.`,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The name of the Route Control Profile object.`,
			},
			"name_alias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The name alias of the Route Control Profile object.`,
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
			"route_control_profile_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("combinable", "global"),
				},
				MarkdownDescription: `The type of the Route Control Profile object. Use combinable when pervasive subnets (fvSubnet) and external subnets (l3extSubnet) should be combined with a route profile and merged into a single route map or route map entry. Use global when the route profile is the only source of information to generate a route map, this will overwrite other policy attributes.`,
			},
			"annotations": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
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
	tflog.Debug(ctx, "End schema of resource: aci_route_control_profile")
}

func (r *RtctrlProfileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: aci_route_control_profile")
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
	tflog.Debug(ctx, "End configure of resource: aci_route_control_profile")
}

func (r *RtctrlProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: aci_route_control_profile")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *RtctrlProfileResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	if stateData.Id.IsUnknown() || stateData.Id.IsNull() {
		SetRtctrlProfileId(ctx, stateData)
	}
	getAndSetRtctrlProfileAttributes(ctx, &resp.Diagnostics, r.client, stateData)
	if !globalAllowExistingOnCreate && !stateData.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Object Already Exists",
			fmt.Sprintf("The rtctrlProfile object with DN '%s' already exists.", stateData.Id.ValueString()),
		)
		return
	}

	var data *RtctrlProfileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsUnknown() || data.Id.IsNull() {
		SetRtctrlProfileId(ctx, data)
	}

	tflog.Debug(ctx, fmt.Sprintf("Create of resource aci_route_control_profile with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationRtctrlProfileResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagRtctrlProfileResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := GetRtctrlProfileCreateJsonPayload(ctx, &resp.Diagnostics, true, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetRtctrlProfileAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End create of resource aci_route_control_profile with id '%s'", data.Id.ValueString()))
}

func (r *RtctrlProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: aci_route_control_profile")
	var data *RtctrlProfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource aci_route_control_profile with id '%s'", data.Id.ValueString()))

	getAndSetRtctrlProfileAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *RtctrlProfileResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, fmt.Sprintf("End read of resource aci_route_control_profile with id '%s'", data.Id.ValueString()))
}

func (r *RtctrlProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: aci_route_control_profile")
	var data *RtctrlProfileResourceModel
	var stateData *RtctrlProfileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource aci_route_control_profile with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationRtctrlProfileResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagRtctrlProfileResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := GetRtctrlProfileCreateJsonPayload(ctx, &resp.Diagnostics, false, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetRtctrlProfileAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End update of resource aci_route_control_profile with id '%s'", data.Id.ValueString()))
}

func (r *RtctrlProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: aci_route_control_profile")
	var data *RtctrlProfileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource aci_route_control_profile with id '%s'", data.Id.ValueString()))
	jsonPayload := GetDeleteJsonPayload(ctx, &resp.Diagnostics, "rtctrlProfile", data.Id.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}
	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("End delete of resource aci_route_control_profile with id '%s'", data.Id.ValueString()))
}

func (r *RtctrlProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: aci_route_control_profile")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *RtctrlProfileResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)
	tflog.Debug(ctx, fmt.Sprintf("Import state of resource aci_route_control_profile with id '%s'", stateData.Id.ValueString()))

	tflog.Debug(ctx, "End import of state resource: aci_route_control_profile")
}

func getAndSetRtctrlProfileAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *RtctrlProfileResourceModel) {
	requestData := DoRestRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=children&rsp-subtree-class=%s", data.Id.ValueString(), "rtctrlProfile,tagAnnotation,tagTag"), "GET", nil)

	*data = *getEmptyRtctrlProfileResourceModel()

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("rtctrlProfile").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("rtctrlProfile").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					data.Id = basetypes.NewStringValue(attributeValue.(string))
					setRtctrlProfileParentDn(ctx, attributeValue.(string), data)
				}
				if attributeName == "annotation" {
					data.Annotation = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "autoContinue" {
					data.AutoContinue = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "descr" {
					data.Descr = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "name" {
					data.Name = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "nameAlias" {
					data.NameAlias = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ownerKey" {
					data.OwnerKey = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ownerTag" {
					data.OwnerTag = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "type" {
					data.Type = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			TagAnnotationRtctrlProfileList := make([]TagAnnotationRtctrlProfileResourceModel, 0)
			TagTagRtctrlProfileList := make([]TagTagRtctrlProfileResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "tagAnnotation" {
							TagAnnotationRtctrlProfile := getEmptyTagAnnotationRtctrlProfileResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationRtctrlProfile.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationRtctrlProfile.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagAnnotationRtctrlProfileList = append(TagAnnotationRtctrlProfileList, TagAnnotationRtctrlProfile)
						}
						if childClassName == "tagTag" {
							TagTagRtctrlProfile := getEmptyTagTagRtctrlProfileResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagTagRtctrlProfile.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagTagRtctrlProfile.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagTagRtctrlProfileList = append(TagTagRtctrlProfileList, TagTagRtctrlProfile)
						}
					}
				}
			}
			tagAnnotationSet, _ := types.SetValueFrom(ctx, data.TagAnnotation.ElementType(ctx), TagAnnotationRtctrlProfileList)
			data.TagAnnotation = tagAnnotationSet
			tagTagSet, _ := types.SetValueFrom(ctx, data.TagTag.ElementType(ctx), TagTagRtctrlProfileList)
			data.TagTag = tagTagSet
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'rtctrlProfile'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	} else {
		data.Id = basetypes.NewStringNull()
	}
}

func getRtctrlProfileRn(ctx context.Context, data *RtctrlProfileResourceModel) string {
	rn := "prof-{name}"
	for _, identifier := range []string{"name"} {
		fieldName := fmt.Sprintf("%s%s", strings.ToUpper(identifier[:1]), identifier[1:])
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface().(basetypes.StringValue).ValueString()
		rn = strings.ReplaceAll(rn, fmt.Sprintf("{%s}", identifier), fieldValue)
	}
	return rn
}

func setRtctrlProfileParentDn(ctx context.Context, dn string, data *RtctrlProfileResourceModel) {
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

func SetRtctrlProfileId(ctx context.Context, data *RtctrlProfileResourceModel) {
	rn := getRtctrlProfileRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
}

func getRtctrlProfileTagAnnotationChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *RtctrlProfileResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationRtctrlProfileResourceModel) []map[string]interface{} {

	childPayloads := []map[string]interface{}{}
	if !data.TagAnnotation.IsUnknown() {
		tagAnnotationIdentifiers := []TagAnnotationIdentifier{}
		for _, tagAnnotation := range tagAnnotationPlan {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			if !tagAnnotation.Key.IsUnknown() && !tagAnnotation.Key.IsNull() {
				childMap["attributes"]["key"] = tagAnnotation.Key.ValueString()
			}
			if !tagAnnotation.Value.IsUnknown() && !tagAnnotation.Value.IsNull() {
				childMap["attributes"]["value"] = tagAnnotation.Value.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"tagAnnotation": childMap})
			tagAnnotationIdentifier := TagAnnotationIdentifier{}
			tagAnnotationIdentifier.Key = tagAnnotation.Key
			tagAnnotationIdentifiers = append(tagAnnotationIdentifiers, tagAnnotationIdentifier)
		}
		for _, tagAnnotation := range tagAnnotationState {
			delete := true
			for _, tagAnnotationIdentifier := range tagAnnotationIdentifiers {
				if tagAnnotationIdentifier.Key == tagAnnotation.Key {
					delete = false
					break
				}
			}
			if delete {
				childMap := map[string]map[string]interface{}{"attributes": {}}
				childMap["attributes"]["status"] = "deleted"
				childMap["attributes"]["key"] = tagAnnotation.Key.ValueString()
				childPayloads = append(childPayloads, map[string]interface{}{"tagAnnotation": childMap})
			}
		}
	} else {
		data.TagAnnotation = types.SetNull(data.TagAnnotation.ElementType(ctx))
	}

	return childPayloads
}
func getRtctrlProfileTagTagChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *RtctrlProfileResourceModel, tagTagPlan, tagTagState []TagTagRtctrlProfileResourceModel) []map[string]interface{} {

	childPayloads := []map[string]interface{}{}
	if !data.TagTag.IsUnknown() {
		tagTagIdentifiers := []TagTagIdentifier{}
		for _, tagTag := range tagTagPlan {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			if !tagTag.Key.IsUnknown() && !tagTag.Key.IsNull() {
				childMap["attributes"]["key"] = tagTag.Key.ValueString()
			}
			if !tagTag.Value.IsUnknown() && !tagTag.Value.IsNull() {
				childMap["attributes"]["value"] = tagTag.Value.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"tagTag": childMap})
			tagTagIdentifier := TagTagIdentifier{}
			tagTagIdentifier.Key = tagTag.Key
			tagTagIdentifiers = append(tagTagIdentifiers, tagTagIdentifier)
		}
		for _, tagTag := range tagTagState {
			delete := true
			for _, tagTagIdentifier := range tagTagIdentifiers {
				if tagTagIdentifier.Key == tagTag.Key {
					delete = false
					break
				}
			}
			if delete {
				childMap := map[string]map[string]interface{}{"attributes": {}}
				childMap["attributes"]["status"] = "deleted"
				childMap["attributes"]["key"] = tagTag.Key.ValueString()
				childPayloads = append(childPayloads, map[string]interface{}{"tagTag": childMap})
			}
		}
	} else {
		data.TagTag = types.SetNull(data.TagTag.ElementType(ctx))
	}

	return childPayloads
}

func GetRtctrlProfileCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, createType bool, data *RtctrlProfileResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationRtctrlProfileResourceModel, tagTagPlan, tagTagState []TagTagRtctrlProfileResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}

	if createType && !globalAllowExistingOnCreate {
		payloadMap["attributes"].(map[string]string)["status"] = "created"
	}
	childPayloads := []map[string]interface{}{}

	TagAnnotationchildPayloads := getRtctrlProfileTagAnnotationChildPayloads(ctx, diags, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)

	TagTagchildPayloads := getRtctrlProfileTagTagChildPayloads(ctx, diags, data, tagTagPlan, tagTagState)
	if TagTagchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagTagchildPayloads...)

	payloadMap["children"] = childPayloads
	if !data.Annotation.IsNull() && !data.Annotation.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["annotation"] = data.Annotation.ValueString()
	}
	if !data.AutoContinue.IsNull() && !data.AutoContinue.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["autoContinue"] = data.AutoContinue.ValueString()
	}
	if !data.Descr.IsNull() && !data.Descr.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["descr"] = data.Descr.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["name"] = data.Name.ValueString()
	}
	if !data.NameAlias.IsNull() && !data.NameAlias.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["nameAlias"] = data.NameAlias.ValueString()
	}
	if !data.OwnerKey.IsNull() && !data.OwnerKey.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["ownerKey"] = data.OwnerKey.ValueString()
	}
	if !data.OwnerTag.IsNull() && !data.OwnerTag.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["ownerTag"] = data.OwnerTag.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["type"] = data.Type.ValueString()
	}

	payload, err := json.Marshal(map[string]interface{}{"rtctrlProfile": payloadMap})
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
