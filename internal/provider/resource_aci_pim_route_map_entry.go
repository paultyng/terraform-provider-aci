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
var _ resource.Resource = &PimRouteMapEntryResource{}
var _ resource.ResourceWithImportState = &PimRouteMapEntryResource{}

func NewPimRouteMapEntryResource() resource.Resource {
	return &PimRouteMapEntryResource{}
}

// PimRouteMapEntryResource defines the resource implementation.
type PimRouteMapEntryResource struct {
	client *client.Client
}

// PimRouteMapEntryResourceModel describes the resource data model.
type PimRouteMapEntryResourceModel struct {
	Id            types.String `tfsdk:"id"`
	ParentDn      types.String `tfsdk:"parent_dn"`
	Action        types.String `tfsdk:"action"`
	Annotation    types.String `tfsdk:"annotation"`
	Descr         types.String `tfsdk:"description"`
	Grp           types.String `tfsdk:"group_ip"`
	Name          types.String `tfsdk:"name"`
	NameAlias     types.String `tfsdk:"name_alias"`
	Order         types.String `tfsdk:"order"`
	Rp            types.String `tfsdk:"rendezvous_point_ip"`
	Src           types.String `tfsdk:"source_ip"`
	TagAnnotation types.Set    `tfsdk:"annotations"`
	TagTag        types.Set    `tfsdk:"tags"`
}

func getEmptyPimRouteMapEntryResourceModel() *PimRouteMapEntryResourceModel {
	return &PimRouteMapEntryResourceModel{
		Id:         basetypes.NewStringNull(),
		ParentDn:   basetypes.NewStringNull(),
		Action:     basetypes.NewStringNull(),
		Annotation: basetypes.NewStringNull(),
		Descr:      basetypes.NewStringNull(),
		Grp:        basetypes.NewStringNull(),
		Name:       basetypes.NewStringNull(),
		NameAlias:  basetypes.NewStringNull(),
		Order:      basetypes.NewStringNull(),
		Rp:         basetypes.NewStringNull(),
		Src:        basetypes.NewStringNull(),
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

// TagAnnotationPimRouteMapEntryResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationPimRouteMapEntryResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func getEmptyTagAnnotationPimRouteMapEntryResourceModel() TagAnnotationPimRouteMapEntryResourceModel {
	return TagAnnotationPimRouteMapEntryResourceModel{
		Key:   basetypes.NewStringNull(),
		Value: basetypes.NewStringNull(),
	}
}

// TagTagPimRouteMapEntryResourceModel describes the resource data model for the children without relation ships.
type TagTagPimRouteMapEntryResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func getEmptyTagTagPimRouteMapEntryResourceModel() TagTagPimRouteMapEntryResourceModel {
	return TagTagPimRouteMapEntryResourceModel{
		Key:   basetypes.NewStringNull(),
		Value: basetypes.NewStringNull(),
	}
}

type PimRouteMapEntryIdentifier struct {
	Order types.String
}

func (r *PimRouteMapEntryResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if !req.Plan.Raw.IsNull() {
		var planData, stateData *PimRouteMapEntryResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

		if resp.Diagnostics.HasError() {
			return
		}

		if (planData.Id.IsUnknown() || planData.Id.IsNull()) && !planData.ParentDn.IsUnknown() && !planData.Order.IsUnknown() {
			setPimRouteMapEntryId(ctx, planData)
		}

		if stateData == nil && !globalAllowExistingOnCreate && !planData.Id.IsUnknown() && !planData.Id.IsNull() {
			CheckDn(ctx, &resp.Diagnostics, r.client, "pimRouteMapEntry", planData.Id.ValueString())
			if resp.Diagnostics.HasError() {
				return
			}
		}

		resp.Diagnostics.Append(resp.Plan.Set(ctx, &planData)...)
	}
}

func (r *PimRouteMapEntryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: aci_pim_route_map_entry")
	resp.TypeName = req.ProviderTypeName + "_pim_route_map_entry"
	tflog.Debug(ctx, "End metadata of resource: aci_pim_route_map_entry")
}

func (r *PimRouteMapEntryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: aci_pim_route_map_entry")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The pim_route_map_entry resource for the 'pimRouteMapEntry' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the Pim Route Map Entry object.",
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
			"action": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("deny", "permit"),
				},
				MarkdownDescription: `The route action of the Pim Route Map Entry object.`,
			},
			"annotation": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Default:             stringdefault.StaticString(globalAnnotation),
				MarkdownDescription: `The annotation of the Pim Route Map Entry object.`,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The description of the Pim Route Map Entry object.`,
			},
			"group_ip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The group ip of the Pim Route Map Entry object.`,
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The name of the Pim Route Map Entry object.`,
			},
			"name_alias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The name alias of the Pim Route Map Entry object.`,
			},
			"order": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `PIM route map entry order.`,
			},
			"rendezvous_point_ip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The rendezvous point ip of the Pim Route Map Entry object.`,
			},
			"source_ip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The source ip of the Pim Route Map Entry object.`,
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
	tflog.Debug(ctx, "End schema of resource: aci_pim_route_map_entry")
}

func (r *PimRouteMapEntryResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: aci_pim_route_map_entry")
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
	tflog.Debug(ctx, "End configure of resource: aci_pim_route_map_entry")
}

func (r *PimRouteMapEntryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: aci_pim_route_map_entry")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *PimRouteMapEntryResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	if stateData.Id.IsUnknown() || stateData.Id.IsNull() {
		setPimRouteMapEntryId(ctx, stateData)
	}
	getAndSetPimRouteMapEntryAttributes(ctx, &resp.Diagnostics, r.client, stateData)
	if !globalAllowExistingOnCreate && !stateData.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Object Already Exists",
			fmt.Sprintf("The pimRouteMapEntry object with DN '%s' already exists.", stateData.Id.ValueString()),
		)
		return
	}

	var data *PimRouteMapEntryResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsUnknown() || data.Id.IsNull() {
		setPimRouteMapEntryId(ctx, data)
	}

	tflog.Debug(ctx, fmt.Sprintf("Create of resource aci_pim_route_map_entry with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationPimRouteMapEntryResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagPimRouteMapEntryResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getPimRouteMapEntryCreateJsonPayload(ctx, &resp.Diagnostics, true, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetPimRouteMapEntryAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End create of resource aci_pim_route_map_entry with id '%s'", data.Id.ValueString()))
}

func (r *PimRouteMapEntryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: aci_pim_route_map_entry")
	var data *PimRouteMapEntryResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource aci_pim_route_map_entry with id '%s'", data.Id.ValueString()))

	getAndSetPimRouteMapEntryAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *PimRouteMapEntryResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, fmt.Sprintf("End read of resource aci_pim_route_map_entry with id '%s'", data.Id.ValueString()))
}

func (r *PimRouteMapEntryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: aci_pim_route_map_entry")
	var data *PimRouteMapEntryResourceModel
	var stateData *PimRouteMapEntryResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource aci_pim_route_map_entry with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationPimRouteMapEntryResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagPimRouteMapEntryResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getPimRouteMapEntryCreateJsonPayload(ctx, &resp.Diagnostics, false, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetPimRouteMapEntryAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End update of resource aci_pim_route_map_entry with id '%s'", data.Id.ValueString()))
}

func (r *PimRouteMapEntryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: aci_pim_route_map_entry")
	var data *PimRouteMapEntryResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource aci_pim_route_map_entry with id '%s'", data.Id.ValueString()))
	jsonPayload := GetDeleteJsonPayload(ctx, &resp.Diagnostics, "pimRouteMapEntry", data.Id.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}
	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("End delete of resource aci_pim_route_map_entry with id '%s'", data.Id.ValueString()))
}

func (r *PimRouteMapEntryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: aci_pim_route_map_entry")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *PimRouteMapEntryResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)
	tflog.Debug(ctx, fmt.Sprintf("Import state of resource aci_pim_route_map_entry with id '%s'", stateData.Id.ValueString()))

	tflog.Debug(ctx, "End import of state resource: aci_pim_route_map_entry")
}

func getAndSetPimRouteMapEntryAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *PimRouteMapEntryResourceModel) {
	requestData := DoRestRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=children&rsp-subtree-class=%s", data.Id.ValueString(), "pimRouteMapEntry,tagAnnotation,tagTag"), "GET", nil)

	readData := getEmptyPimRouteMapEntryResourceModel()

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("pimRouteMapEntry").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("pimRouteMapEntry").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					readData.Id = basetypes.NewStringValue(attributeValue.(string))
					setPimRouteMapEntryParentDn(ctx, attributeValue.(string), readData)
				}
				if attributeName == "action" {
					readData.Action = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "annotation" {
					readData.Annotation = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "descr" {
					readData.Descr = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "grp" {
					readData.Grp = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "name" {
					readData.Name = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "nameAlias" {
					readData.NameAlias = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "order" {
					readData.Order = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "rp" {
					readData.Rp = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "src" {
					readData.Src = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			TagAnnotationPimRouteMapEntryList := make([]TagAnnotationPimRouteMapEntryResourceModel, 0)
			TagTagPimRouteMapEntryList := make([]TagTagPimRouteMapEntryResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "tagAnnotation" {
							TagAnnotationPimRouteMapEntry := getEmptyTagAnnotationPimRouteMapEntryResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationPimRouteMapEntry.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationPimRouteMapEntry.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagAnnotationPimRouteMapEntryList = append(TagAnnotationPimRouteMapEntryList, TagAnnotationPimRouteMapEntry)
						}
						if childClassName == "tagTag" {
							TagTagPimRouteMapEntry := getEmptyTagTagPimRouteMapEntryResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagTagPimRouteMapEntry.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagTagPimRouteMapEntry.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagTagPimRouteMapEntryList = append(TagTagPimRouteMapEntryList, TagTagPimRouteMapEntry)
						}
					}
				}
			}
			tagAnnotationSet, _ := types.SetValueFrom(ctx, readData.TagAnnotation.ElementType(ctx), TagAnnotationPimRouteMapEntryList)
			readData.TagAnnotation = tagAnnotationSet
			tagTagSet, _ := types.SetValueFrom(ctx, readData.TagTag.ElementType(ctx), TagTagPimRouteMapEntryList)
			readData.TagTag = tagTagSet
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'pimRouteMapEntry'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	} else {
		readData.Id = basetypes.NewStringNull()
	}
	*data = *readData
}

func getPimRouteMapEntryRn(ctx context.Context, data *PimRouteMapEntryResourceModel) string {
	rn := "rtmapentry-{order}"
	for _, identifier := range []string{"order"} {
		fieldName := fmt.Sprintf("%s%s", strings.ToUpper(identifier[:1]), identifier[1:])
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface().(basetypes.StringValue).ValueString()
		rn = strings.ReplaceAll(rn, fmt.Sprintf("{%s}", identifier), fieldValue)
	}
	return rn
}

func setPimRouteMapEntryParentDn(ctx context.Context, dn string, data *PimRouteMapEntryResourceModel) {
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

func setPimRouteMapEntryId(ctx context.Context, data *PimRouteMapEntryResourceModel) {
	rn := getPimRouteMapEntryRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
}

func getPimRouteMapEntryTagAnnotationChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *PimRouteMapEntryResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationPimRouteMapEntryResourceModel) []map[string]interface{} {

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
func getPimRouteMapEntryTagTagChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *PimRouteMapEntryResourceModel, tagTagPlan, tagTagState []TagTagPimRouteMapEntryResourceModel) []map[string]interface{} {

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

func getPimRouteMapEntryCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, createType bool, data *PimRouteMapEntryResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationPimRouteMapEntryResourceModel, tagTagPlan, tagTagState []TagTagPimRouteMapEntryResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}

	if createType && !globalAllowExistingOnCreate {
		payloadMap["attributes"].(map[string]string)["status"] = "created"
	}
	childPayloads := []map[string]interface{}{}

	TagAnnotationchildPayloads := getPimRouteMapEntryTagAnnotationChildPayloads(ctx, diags, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)

	TagTagchildPayloads := getPimRouteMapEntryTagTagChildPayloads(ctx, diags, data, tagTagPlan, tagTagState)
	if TagTagchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagTagchildPayloads...)

	payloadMap["children"] = childPayloads
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["action"] = data.Action.ValueString()
	}
	if !data.Annotation.IsNull() && !data.Annotation.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["annotation"] = data.Annotation.ValueString()
	}
	if !data.Descr.IsNull() && !data.Descr.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["descr"] = data.Descr.ValueString()
	}
	if !data.Grp.IsNull() && !data.Grp.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["grp"] = data.Grp.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["name"] = data.Name.ValueString()
	}
	if !data.NameAlias.IsNull() && !data.NameAlias.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["nameAlias"] = data.NameAlias.ValueString()
	}
	if !data.Order.IsNull() && !data.Order.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["order"] = data.Order.ValueString()
	}
	if !data.Rp.IsNull() && !data.Rp.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["rp"] = data.Rp.ValueString()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["src"] = data.Src.ValueString()
	}
	payload, err := json.Marshal(map[string]interface{}{"pimRouteMapEntry": payloadMap})
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
