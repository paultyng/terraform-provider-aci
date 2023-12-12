// Code generated by "gen/generator.go"; DO NOT EDIT.

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
var _ resource.Resource = &MgmtInstPResource{}
var _ resource.ResourceWithImportState = &MgmtInstPResource{}

func NewMgmtInstPResource() resource.Resource {
	return &MgmtInstPResource{}
}

// MgmtInstPResource defines the resource implementation.
type MgmtInstPResource struct {
	client *client.Client
}

// MgmtInstPResourceModel describes the resource data model.
type MgmtInstPResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Annotation    types.String `tfsdk:"annotation"`
	Descr         types.String `tfsdk:"description"`
	Name          types.String `tfsdk:"name"`
	NameAlias     types.String `tfsdk:"name_alias"`
	Prio          types.String `tfsdk:"priority"`
	MgmtRsOoBCons types.Set    `tfsdk:"relation_to_consumed_out_of_band_contracts"`
	TagAnnotation types.Set    `tfsdk:"annotations"`
}

// MgmtRsOoBConsMgmtInstPResourceModel describes the resource data model for the children without relation ships.
type MgmtRsOoBConsMgmtInstPResourceModel struct {
	Annotation      types.String `tfsdk:"annotation"`
	Prio            types.String `tfsdk:"priority"`
	TnVzOOBBrCPName types.String `tfsdk:"out_of_band_contract_name"`
}

// TagAnnotationMgmtInstPResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationMgmtInstPResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type MgmtInstPIdentifier struct {
	Name types.String
}

func (r *MgmtInstPResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	tflog.Debug(ctx, "Start plan modification of resource: aci_external_management_network_instance_profile")

	if !req.Plan.Raw.IsNull() {
		var planData *MgmtInstPResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)

		var MgmtRsOoBConsPlan, MgmtRsOoBConsUpdate []MgmtRsOoBConsResourceModel
		planData.MgmtRsOoBCons.ElementsAs(ctx, &MgmtRsOoBConsPlan, false)
		if len(MgmtRsOoBConsPlan) > 0 {
			for _, MgmtRsOoBCons := range MgmtRsOoBConsPlan {
				if r.client.ValidateRelationDn {
					CheckDn(ctx, r.client, MgmtRsOoBCons.TnVzOOBBrCPName.ValueString(), &resp.Diagnostics)
				}
				MgmtRsOoBCons.TnVzOOBBrCPName = basetypes.NewStringValue(GetMOName(MgmtRsOoBCons.TnVzOOBBrCPName.ValueString()))
				MgmtRsOoBConsUpdate = append(MgmtRsOoBConsUpdate, MgmtRsOoBCons)
			}
			MgmtRsOoBConsSet, _ := types.SetValueFrom(ctx, planData.MgmtRsOoBCons.ElementType(ctx), MgmtRsOoBConsUpdate)
			planData.MgmtRsOoBCons = MgmtRsOoBConsSet
		}

		resp.Plan.Set(ctx, planData)
	}
	tflog.Debug(ctx, "End plan modification of resource: aci_external_management_network_instance_profile")
}

func (r *MgmtInstPResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: aci_external_management_network_instance_profile")
	resp.TypeName = req.ProviderTypeName + "_external_management_network_instance_profile"
	tflog.Debug(ctx, "End metadata of resource: aci_external_management_network_instance_profile")
}

func (r *MgmtInstPResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: aci_external_management_network_instance_profile")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The external_management_network_instance_profile resource for the 'mgmtInstP' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the External Management Network Instance Profile object.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"annotation": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Default:             stringdefault.StaticString(globalAnnotation),
				MarkdownDescription: `The annotation of the External Management Network Instance Profile object.`,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `The description of the External Management Network Instance Profile object.`,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The name of the External Management Network Instance Profile object.`,
			},
			"name_alias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `The name alias of the External Management Network Instance Profile object.`,
			},
			"priority": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("level1", "level2", "level3", "level4", "level5", "level6", "unspecified"),
				},
				MarkdownDescription: `The QoS priority class identifier.`,
			},
			"relation_to_consumed_out_of_band_contracts": schema.SetNestedAttribute{
				MarkdownDescription: `An external management entity instance profile to an out-of-band binary contract profile. The instance profiles of external management entities can communicate with nodes that are part of out-of-band management endpoint group. To enable this communication, a contract is required between the instance profile and the out-of-band management endpoint group.`,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"annotation": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `The annotation of the Relation To Consumed Out Of Band Contract object.`,
						},
						"priority": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf("level1", "level2", "level3", "level4", "level5", "level6", "unspecified"),
							},
							MarkdownDescription: `The Quality of service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.`,
						},
						"out_of_band_contract_name": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `The name of the Out Of Band Contract object.`,
						},
					},
				},
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
		},
	}
	tflog.Debug(ctx, "End schema of resource: aci_external_management_network_instance_profile")
}

func (r *MgmtInstPResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: aci_external_management_network_instance_profile")
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
	tflog.Debug(ctx, "End configure of resource: aci_external_management_network_instance_profile")
}

func (r *MgmtInstPResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: aci_external_management_network_instance_profile")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *MgmtInstPResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	setMgmtInstPId(ctx, stateData)
	setMgmtInstPAttributes(ctx, &resp.Diagnostics, r.client, stateData)

	var data *MgmtInstPResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setMgmtInstPId(ctx, data)

	tflog.Debug(ctx, fmt.Sprintf("Create of resource aci_external_management_network_instance_profile with id '%s'", data.Id.ValueString()))

	var mgmtRsOoBConsPlan, mgmtRsOoBConsState []MgmtRsOoBConsMgmtInstPResourceModel
	data.MgmtRsOoBCons.ElementsAs(ctx, &mgmtRsOoBConsPlan, false)
	stateData.MgmtRsOoBCons.ElementsAs(ctx, &mgmtRsOoBConsState, false)
	var tagAnnotationPlan, tagAnnotationState []TagAnnotationMgmtInstPResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	jsonPayload := getMgmtInstPCreateJsonPayload(ctx, &resp.Diagnostics, data, mgmtRsOoBConsPlan, mgmtRsOoBConsState, tagAnnotationPlan, tagAnnotationState)

	if resp.Diagnostics.HasError() {
		return
	}

	doMgmtInstPRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}

	setMgmtInstPAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, "End create of resource: aci_external_management_network_instance_profile")
}

func (r *MgmtInstPResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: aci_external_management_network_instance_profile")
	var data *MgmtInstPResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource aci_external_management_network_instance_profile with id '%s'", data.Id.ValueString()))

	setMgmtInstPAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *MgmtInstPResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, "End read of resource: aci_external_management_network_instance_profile")
}

func (r *MgmtInstPResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: aci_external_management_network_instance_profile")
	var data *MgmtInstPResourceModel
	var stateData *MgmtInstPResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource aci_external_management_network_instance_profile with id '%s'", data.Id.ValueString()))

	var mgmtRsOoBConsPlan, mgmtRsOoBConsState []MgmtRsOoBConsMgmtInstPResourceModel
	data.MgmtRsOoBCons.ElementsAs(ctx, &mgmtRsOoBConsPlan, false)
	stateData.MgmtRsOoBCons.ElementsAs(ctx, &mgmtRsOoBConsState, false)
	var tagAnnotationPlan, tagAnnotationState []TagAnnotationMgmtInstPResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	jsonPayload := getMgmtInstPCreateJsonPayload(ctx, &resp.Diagnostics, data, mgmtRsOoBConsPlan, mgmtRsOoBConsState, tagAnnotationPlan, tagAnnotationState)

	if resp.Diagnostics.HasError() {
		return
	}

	doMgmtInstPRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	setMgmtInstPAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, "End update of resource: aci_external_management_network_instance_profile")
}

func (r *MgmtInstPResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: aci_external_management_network_instance_profile")
	var data *MgmtInstPResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource aci_external_management_network_instance_profile with id '%s'", data.Id.ValueString()))
	jsonPayload := getMgmtInstPDeleteJsonPayload(ctx, &resp.Diagnostics, data)
	if resp.Diagnostics.HasError() {
		return
	}
	doMgmtInstPRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "End delete of resource: aci_external_management_network_instance_profile")
}

func (r *MgmtInstPResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: aci_external_management_network_instance_profile")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *MgmtInstPResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)
	tflog.Debug(ctx, fmt.Sprintf("Import state of resource aci_external_management_network_instance_profile with id '%s'", stateData.Id.ValueString()))

	tflog.Debug(ctx, "End import of state resource: aci_external_management_network_instance_profile")
}

func setMgmtInstPAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *MgmtInstPResourceModel) {
	requestData := doMgmtInstPRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=children&rsp-subtree-class=%s", data.Id.ValueString(), "mgmtInstP,mgmtRsOoBCons,tagAnnotation"), "GET", nil)

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("mgmtInstP").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("mgmtInstP").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					data.Id = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "annotation" {
					data.Annotation = basetypes.NewStringValue(attributeValue.(string))
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
				if attributeName == "prio" {
					data.Prio = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			MgmtRsOoBConsMgmtInstPList := make([]MgmtRsOoBConsMgmtInstPResourceModel, 0)
			TagAnnotationMgmtInstPList := make([]TagAnnotationMgmtInstPResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "mgmtRsOoBCons" {
							MgmtRsOoBConsMgmtInstP := MgmtRsOoBConsMgmtInstPResourceModel{}
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "annotation" {
									MgmtRsOoBConsMgmtInstP.Annotation = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "prio" {
									MgmtRsOoBConsMgmtInstP.Prio = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "tnVzOOBBrCPName" {
									MgmtRsOoBConsMgmtInstP.TnVzOOBBrCPName = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							MgmtRsOoBConsMgmtInstPList = append(MgmtRsOoBConsMgmtInstPList, MgmtRsOoBConsMgmtInstP)
						}
						if childClassName == "tagAnnotation" {
							TagAnnotationMgmtInstP := TagAnnotationMgmtInstPResourceModel{}
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationMgmtInstP.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationMgmtInstP.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagAnnotationMgmtInstPList = append(TagAnnotationMgmtInstPList, TagAnnotationMgmtInstP)
						}
					}
				}
			}
			if len(MgmtRsOoBConsMgmtInstPList) > 0 {
				mgmtRsOoBConsSet, _ := types.SetValueFrom(ctx, data.MgmtRsOoBCons.ElementType(ctx), MgmtRsOoBConsMgmtInstPList)
				data.MgmtRsOoBCons = mgmtRsOoBConsSet
			}
			if len(TagAnnotationMgmtInstPList) > 0 {
				tagAnnotationSet, _ := types.SetValueFrom(ctx, data.TagAnnotation.ElementType(ctx), TagAnnotationMgmtInstPList)
				data.TagAnnotation = tagAnnotationSet
			}
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'mgmtInstP'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	} else {
		data.Id = basetypes.NewStringNull()
	}
}

func getMgmtInstPRn(ctx context.Context, data *MgmtInstPResourceModel) string {
	rn := "tn-mgmt/extmgmt-default/instp-{name}"
	for _, identifier := range []string{"name"} {
		fieldName := fmt.Sprintf("%s%s", strings.ToUpper(identifier[:1]), identifier[1:])
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface().(basetypes.StringValue).ValueString()
		rn = strings.ReplaceAll(rn, fmt.Sprintf("{%s}", identifier), fieldValue)
	}
	return rn
}

func setMgmtInstPId(ctx context.Context, data *MgmtInstPResourceModel) {
	rn := getMgmtInstPRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", strings.Split([]string{"uni/tn-mgmt/extmgmt-default/instp-{name}"}[0], "/")[0], rn))
}

func getMgmtInstPMgmtRsOoBConsChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *MgmtInstPResourceModel, mgmtRsOoBConsPlan, mgmtRsOoBConsState []MgmtRsOoBConsMgmtInstPResourceModel) []map[string]interface{} {

	childPayloads := []map[string]interface{}{}
	if !data.MgmtRsOoBCons.IsUnknown() {
		mgmtRsOoBConsIdentifiers := []MgmtRsOoBConsIdentifier{}
		for _, mgmtRsOoBCons := range mgmtRsOoBConsPlan {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			if !mgmtRsOoBCons.Annotation.IsUnknown() {
				childMap["attributes"]["annotation"] = mgmtRsOoBCons.Annotation.ValueString()
			} else {
				childMap["attributes"]["annotation"] = globalAnnotation
			}
			if !mgmtRsOoBCons.Prio.IsUnknown() {
				childMap["attributes"]["prio"] = mgmtRsOoBCons.Prio.ValueString()
			}
			if !mgmtRsOoBCons.TnVzOOBBrCPName.IsUnknown() {
				childMap["attributes"]["tnVzOOBBrCPName"] = mgmtRsOoBCons.TnVzOOBBrCPName.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"mgmtRsOoBCons": childMap})
			mgmtRsOoBConsIdentifier := MgmtRsOoBConsIdentifier{}
			mgmtRsOoBConsIdentifier.TnVzOOBBrCPName = mgmtRsOoBCons.TnVzOOBBrCPName
			mgmtRsOoBConsIdentifiers = append(mgmtRsOoBConsIdentifiers, mgmtRsOoBConsIdentifier)
		}
		for _, mgmtRsOoBCons := range mgmtRsOoBConsState {
			delete := true
			for _, mgmtRsOoBConsIdentifier := range mgmtRsOoBConsIdentifiers {
				if mgmtRsOoBConsIdentifier.TnVzOOBBrCPName == mgmtRsOoBCons.TnVzOOBBrCPName {
					delete = false
					break
				}
			}
			if delete {
				childMap := map[string]map[string]interface{}{"attributes": {}}
				childMap["attributes"]["status"] = "deleted"
				childMap["attributes"]["tnVzOOBBrCPName"] = mgmtRsOoBCons.TnVzOOBBrCPName.ValueString()
				childPayloads = append(childPayloads, map[string]interface{}{"mgmtRsOoBCons": childMap})
			}
		}
	} else {
		data.MgmtRsOoBCons = types.SetNull(data.MgmtRsOoBCons.ElementType(ctx))
	}

	return childPayloads
}
func getMgmtInstPTagAnnotationChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *MgmtInstPResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationMgmtInstPResourceModel) []map[string]interface{} {

	childPayloads := []map[string]interface{}{}
	if !data.TagAnnotation.IsUnknown() {
		tagAnnotationIdentifiers := []TagAnnotationIdentifier{}
		for _, tagAnnotation := range tagAnnotationPlan {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			if !tagAnnotation.Key.IsUnknown() {
				childMap["attributes"]["key"] = tagAnnotation.Key.ValueString()
			}
			if !tagAnnotation.Value.IsUnknown() {
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

func getMgmtInstPCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, data *MgmtInstPResourceModel, mgmtRsOoBConsPlan, mgmtRsOoBConsState []MgmtRsOoBConsMgmtInstPResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationMgmtInstPResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}
	childPayloads := []map[string]interface{}{}

	MgmtRsOoBConschildPayloads := getMgmtInstPMgmtRsOoBConsChildPayloads(ctx, diags, data, mgmtRsOoBConsPlan, mgmtRsOoBConsState)
	if MgmtRsOoBConschildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, MgmtRsOoBConschildPayloads...)

	TagAnnotationchildPayloads := getMgmtInstPTagAnnotationChildPayloads(ctx, diags, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)

	payloadMap["children"] = childPayloads
	if !data.Annotation.IsNull() && !data.Annotation.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["annotation"] = data.Annotation.ValueString()
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
	if !data.Prio.IsNull() && !data.Prio.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["prio"] = data.Prio.ValueString()
	}

	payload, err := json.Marshal(map[string]interface{}{"mgmtInstP": payloadMap})
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

func getMgmtInstPDeleteJsonPayload(ctx context.Context, diags *diag.Diagnostics, data *MgmtInstPResourceModel) *container.Container {

	jsonString := fmt.Sprintf(`{"mgmtInstP":{"attributes":{"dn": "%s","status": "deleted"}}}`, data.Id.ValueString())
	jsonPayload, err := container.ParseJSON([]byte(jsonString))
	if err != nil {
		diags.AddError(
			"Construction of json payload failed",
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}
	return jsonPayload
}

func doMgmtInstPRequest(ctx context.Context, diags *diag.Diagnostics, client *client.Client, path, method string, payload *container.Container) *container.Container {

	restRequest, err := client.MakeRestRequest(method, path, payload, true)
	if err != nil {
		diags.AddError(
			"Creation of rest request failed",
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	cont, restResponse, err := client.Do(restRequest)

	if restResponse != nil && restResponse.StatusCode != 200 {
		diags.AddError(
			fmt.Sprintf("The %s rest request failed", strings.ToLower(method)),
			fmt.Sprintf("Response: %s, err: %s. Please report this issue to the provider developers.", cont.Data().(map[string]interface{})["imdata"], err),
		)
		return nil
	} else if err != nil {
		diags.AddError(
			fmt.Sprintf("The %s rest request failed", strings.ToLower(method)),
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	return cont
}
