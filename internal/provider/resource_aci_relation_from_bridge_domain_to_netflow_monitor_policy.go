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
var _ resource.Resource = &FvRsBDToNetflowMonitorPolResource{}
var _ resource.ResourceWithImportState = &FvRsBDToNetflowMonitorPolResource{}

func NewFvRsBDToNetflowMonitorPolResource() resource.Resource {
	return &FvRsBDToNetflowMonitorPolResource{}
}

// FvRsBDToNetflowMonitorPolResource defines the resource implementation.
type FvRsBDToNetflowMonitorPolResource struct {
	client *client.Client
}

// FvRsBDToNetflowMonitorPolResourceModel describes the resource data model.
type FvRsBDToNetflowMonitorPolResourceModel struct {
	Id                      types.String `tfsdk:"id"`
	ParentDn                types.String `tfsdk:"parent_dn"`
	Annotation              types.String `tfsdk:"annotation"`
	FltType                 types.String `tfsdk:"filter_type"`
	TnNetflowMonitorPolName types.String `tfsdk:"netflow_monitor_policy_name"`
	TagAnnotation           types.Set    `tfsdk:"annotations"`
	TagTag                  types.Set    `tfsdk:"tags"`
}

func getEmptyFvRsBDToNetflowMonitorPolResourceModel() *FvRsBDToNetflowMonitorPolResourceModel {
	return &FvRsBDToNetflowMonitorPolResourceModel{
		Id:                      basetypes.NewStringNull(),
		ParentDn:                basetypes.NewStringNull(),
		Annotation:              basetypes.NewStringNull(),
		FltType:                 basetypes.NewStringNull(),
		TnNetflowMonitorPolName: basetypes.NewStringNull(),
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

// TagAnnotationFvRsBDToNetflowMonitorPolResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationFvRsBDToNetflowMonitorPolResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func getEmptyTagAnnotationFvRsBDToNetflowMonitorPolResourceModel() TagAnnotationFvRsBDToNetflowMonitorPolResourceModel {
	return TagAnnotationFvRsBDToNetflowMonitorPolResourceModel{
		Key:   basetypes.NewStringNull(),
		Value: basetypes.NewStringNull(),
	}
}

// TagTagFvRsBDToNetflowMonitorPolResourceModel describes the resource data model for the children without relation ships.
type TagTagFvRsBDToNetflowMonitorPolResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func getEmptyTagTagFvRsBDToNetflowMonitorPolResourceModel() TagTagFvRsBDToNetflowMonitorPolResourceModel {
	return TagTagFvRsBDToNetflowMonitorPolResourceModel{
		Key:   basetypes.NewStringNull(),
		Value: basetypes.NewStringNull(),
	}
}

type FvRsBDToNetflowMonitorPolIdentifier struct {
	FltType                 types.String
	TnNetflowMonitorPolName types.String
}

func (r *FvRsBDToNetflowMonitorPolResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if !req.Plan.Raw.IsNull() {
		var planData, stateData *FvRsBDToNetflowMonitorPolResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

		if resp.Diagnostics.HasError() {
			return
		}

		if (planData.Id.IsUnknown() || planData.Id.IsNull()) && !planData.ParentDn.IsUnknown() && !planData.FltType.IsUnknown() && !planData.TnNetflowMonitorPolName.IsUnknown() {
			setFvRsBDToNetflowMonitorPolId(ctx, planData)
		}

		if stateData == nil && !globalAllowExistingOnCreate && !planData.Id.IsUnknown() && !planData.Id.IsNull() {
			CheckDn(ctx, &resp.Diagnostics, r.client, "fvRsBDToNetflowMonitorPol", planData.Id.ValueString())
			if resp.Diagnostics.HasError() {
				return
			}
		}

		resp.Diagnostics.Append(resp.Plan.Set(ctx, &planData)...)
	}
}

func (r *FvRsBDToNetflowMonitorPolResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: aci_relation_from_bridge_domain_to_netflow_monitor_policy")
	resp.TypeName = req.ProviderTypeName + "_relation_from_bridge_domain_to_netflow_monitor_policy"
	tflog.Debug(ctx, "End metadata of resource: aci_relation_from_bridge_domain_to_netflow_monitor_policy")
}

func (r *FvRsBDToNetflowMonitorPolResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: aci_relation_from_bridge_domain_to_netflow_monitor_policy")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The relation_from_bridge_domain_to_netflow_monitor_policy resource for the 'fvRsBDToNetflowMonitorPol' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the Relation From Bridge Domain To NetFlow Monitor Policy object.",
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
				MarkdownDescription: `The annotation of the Relation From Bridge Domain To NetFlow Monitor Policy object.`,
			},
			"filter_type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("ce", "ipv4", "ipv6", "unspecified"),
				},
				MarkdownDescription: `The filter type of the NetFlow Monitor Policy object.`,
			},
			"netflow_monitor_policy_name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The name of the NetFlow Monitor Policy object.`,
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
	tflog.Debug(ctx, "End schema of resource: aci_relation_from_bridge_domain_to_netflow_monitor_policy")
}

func (r *FvRsBDToNetflowMonitorPolResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: aci_relation_from_bridge_domain_to_netflow_monitor_policy")
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
	tflog.Debug(ctx, "End configure of resource: aci_relation_from_bridge_domain_to_netflow_monitor_policy")
}

func (r *FvRsBDToNetflowMonitorPolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: aci_relation_from_bridge_domain_to_netflow_monitor_policy")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *FvRsBDToNetflowMonitorPolResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	if stateData.Id.IsUnknown() || stateData.Id.IsNull() {
		setFvRsBDToNetflowMonitorPolId(ctx, stateData)
	}
	getAndSetFvRsBDToNetflowMonitorPolAttributes(ctx, &resp.Diagnostics, r.client, stateData)
	if !globalAllowExistingOnCreate && !stateData.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Object Already Exists",
			fmt.Sprintf("The fvRsBDToNetflowMonitorPol object with DN '%s' already exists.", stateData.Id.ValueString()),
		)
		return
	}

	var data *FvRsBDToNetflowMonitorPolResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsUnknown() || data.Id.IsNull() {
		setFvRsBDToNetflowMonitorPolId(ctx, data)
	}

	tflog.Debug(ctx, fmt.Sprintf("Create of resource aci_relation_from_bridge_domain_to_netflow_monitor_policy with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationFvRsBDToNetflowMonitorPolResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagFvRsBDToNetflowMonitorPolResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getFvRsBDToNetflowMonitorPolCreateJsonPayload(ctx, &resp.Diagnostics, true, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetFvRsBDToNetflowMonitorPolAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End create of resource aci_relation_from_bridge_domain_to_netflow_monitor_policy with id '%s'", data.Id.ValueString()))
}

func (r *FvRsBDToNetflowMonitorPolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: aci_relation_from_bridge_domain_to_netflow_monitor_policy")
	var data *FvRsBDToNetflowMonitorPolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource aci_relation_from_bridge_domain_to_netflow_monitor_policy with id '%s'", data.Id.ValueString()))

	getAndSetFvRsBDToNetflowMonitorPolAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *FvRsBDToNetflowMonitorPolResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, fmt.Sprintf("End read of resource aci_relation_from_bridge_domain_to_netflow_monitor_policy with id '%s'", data.Id.ValueString()))
}

func (r *FvRsBDToNetflowMonitorPolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: aci_relation_from_bridge_domain_to_netflow_monitor_policy")
	var data *FvRsBDToNetflowMonitorPolResourceModel
	var stateData *FvRsBDToNetflowMonitorPolResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource aci_relation_from_bridge_domain_to_netflow_monitor_policy with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationFvRsBDToNetflowMonitorPolResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagFvRsBDToNetflowMonitorPolResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getFvRsBDToNetflowMonitorPolCreateJsonPayload(ctx, &resp.Diagnostics, false, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetFvRsBDToNetflowMonitorPolAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End update of resource aci_relation_from_bridge_domain_to_netflow_monitor_policy with id '%s'", data.Id.ValueString()))
}

func (r *FvRsBDToNetflowMonitorPolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: aci_relation_from_bridge_domain_to_netflow_monitor_policy")
	var data *FvRsBDToNetflowMonitorPolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource aci_relation_from_bridge_domain_to_netflow_monitor_policy with id '%s'", data.Id.ValueString()))
	jsonPayload := GetDeleteJsonPayload(ctx, &resp.Diagnostics, "fvRsBDToNetflowMonitorPol", data.Id.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}
	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("End delete of resource aci_relation_from_bridge_domain_to_netflow_monitor_policy with id '%s'", data.Id.ValueString()))
}

func (r *FvRsBDToNetflowMonitorPolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: aci_relation_from_bridge_domain_to_netflow_monitor_policy")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *FvRsBDToNetflowMonitorPolResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)
	tflog.Debug(ctx, fmt.Sprintf("Import state of resource aci_relation_from_bridge_domain_to_netflow_monitor_policy with id '%s'", stateData.Id.ValueString()))

	tflog.Debug(ctx, "End import of state resource: aci_relation_from_bridge_domain_to_netflow_monitor_policy")
}

func getAndSetFvRsBDToNetflowMonitorPolAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *FvRsBDToNetflowMonitorPolResourceModel) {
	requestData := DoRestRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=children&rsp-subtree-class=%s", data.Id.ValueString(), "fvRsBDToNetflowMonitorPol,tagAnnotation,tagTag"), "GET", nil)

	readData := getEmptyFvRsBDToNetflowMonitorPolResourceModel()

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("fvRsBDToNetflowMonitorPol").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("fvRsBDToNetflowMonitorPol").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					readData.Id = basetypes.NewStringValue(attributeValue.(string))
					setFvRsBDToNetflowMonitorPolParentDn(ctx, attributeValue.(string), readData)
				}
				if attributeName == "annotation" {
					readData.Annotation = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "fltType" {
					readData.FltType = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "tnNetflowMonitorPolName" {
					readData.TnNetflowMonitorPolName = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			TagAnnotationFvRsBDToNetflowMonitorPolList := make([]TagAnnotationFvRsBDToNetflowMonitorPolResourceModel, 0)
			TagTagFvRsBDToNetflowMonitorPolList := make([]TagTagFvRsBDToNetflowMonitorPolResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "tagAnnotation" {
							TagAnnotationFvRsBDToNetflowMonitorPol := getEmptyTagAnnotationFvRsBDToNetflowMonitorPolResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationFvRsBDToNetflowMonitorPol.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationFvRsBDToNetflowMonitorPol.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagAnnotationFvRsBDToNetflowMonitorPolList = append(TagAnnotationFvRsBDToNetflowMonitorPolList, TagAnnotationFvRsBDToNetflowMonitorPol)
						}
						if childClassName == "tagTag" {
							TagTagFvRsBDToNetflowMonitorPol := getEmptyTagTagFvRsBDToNetflowMonitorPolResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagTagFvRsBDToNetflowMonitorPol.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagTagFvRsBDToNetflowMonitorPol.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagTagFvRsBDToNetflowMonitorPolList = append(TagTagFvRsBDToNetflowMonitorPolList, TagTagFvRsBDToNetflowMonitorPol)
						}
					}
				}
			}
			tagAnnotationSet, _ := types.SetValueFrom(ctx, readData.TagAnnotation.ElementType(ctx), TagAnnotationFvRsBDToNetflowMonitorPolList)
			readData.TagAnnotation = tagAnnotationSet
			tagTagSet, _ := types.SetValueFrom(ctx, readData.TagTag.ElementType(ctx), TagTagFvRsBDToNetflowMonitorPolList)
			readData.TagTag = tagTagSet
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'fvRsBDToNetflowMonitorPol'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	} else {
		readData.Id = basetypes.NewStringNull()
	}
	*data = *readData
}

func getFvRsBDToNetflowMonitorPolRn(ctx context.Context, data *FvRsBDToNetflowMonitorPolResourceModel) string {
	rn := "rsBDToNetflowMonitorPol-[{tnNetflowMonitorPolName}]-{fltType}"
	for _, identifier := range []string{"tnNetflowMonitorPolName", "fltType"} {
		fieldName := fmt.Sprintf("%s%s", strings.ToUpper(identifier[:1]), identifier[1:])
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface().(basetypes.StringValue).ValueString()
		rn = strings.ReplaceAll(rn, fmt.Sprintf("{%s}", identifier), fieldValue)
	}
	return rn
}

func setFvRsBDToNetflowMonitorPolParentDn(ctx context.Context, dn string, data *FvRsBDToNetflowMonitorPolResourceModel) {
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

func setFvRsBDToNetflowMonitorPolId(ctx context.Context, data *FvRsBDToNetflowMonitorPolResourceModel) {
	rn := getFvRsBDToNetflowMonitorPolRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
}

func getFvRsBDToNetflowMonitorPolTagAnnotationChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *FvRsBDToNetflowMonitorPolResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationFvRsBDToNetflowMonitorPolResourceModel) []map[string]interface{} {

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
func getFvRsBDToNetflowMonitorPolTagTagChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *FvRsBDToNetflowMonitorPolResourceModel, tagTagPlan, tagTagState []TagTagFvRsBDToNetflowMonitorPolResourceModel) []map[string]interface{} {

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

func getFvRsBDToNetflowMonitorPolCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, createType bool, data *FvRsBDToNetflowMonitorPolResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationFvRsBDToNetflowMonitorPolResourceModel, tagTagPlan, tagTagState []TagTagFvRsBDToNetflowMonitorPolResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}

	if createType && !globalAllowExistingOnCreate {
		payloadMap["attributes"].(map[string]string)["status"] = "created"
	}
	childPayloads := []map[string]interface{}{}

	TagAnnotationchildPayloads := getFvRsBDToNetflowMonitorPolTagAnnotationChildPayloads(ctx, diags, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)

	TagTagchildPayloads := getFvRsBDToNetflowMonitorPolTagTagChildPayloads(ctx, diags, data, tagTagPlan, tagTagState)
	if TagTagchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagTagchildPayloads...)

	payloadMap["children"] = childPayloads
	if !data.Annotation.IsNull() && !data.Annotation.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["annotation"] = data.Annotation.ValueString()
	}
	if !data.FltType.IsNull() && !data.FltType.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["fltType"] = data.FltType.ValueString()
	}
	if !data.TnNetflowMonitorPolName.IsNull() && !data.TnNetflowMonitorPolName.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["tnNetflowMonitorPolName"] = data.TnNetflowMonitorPolName.ValueString()
	}
	payload, err := json.Marshal(map[string]interface{}{"fvRsBDToNetflowMonitorPol": payloadMap})
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
