// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"context"
	"encoding/json"
	"fmt"

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
var _ resource.Resource = &L3extProvLblResource{}
var _ resource.ResourceWithImportState = &L3extProvLblResource{}

func NewL3extProvLblResource() resource.Resource {
	return &L3extProvLblResource{}
}

// L3extProvLblResource defines the resource implementation.
type L3extProvLblResource struct {
	client *client.Client
}

// L3extProvLblResourceModel describes the resource data model.
type L3extProvLblResourceModel struct {
	Id            types.String `tfsdk:"id"`
	ParentDn      types.String `tfsdk:"parent_dn"`
	Annotation    types.String `tfsdk:"annotation"`
	Descr         types.String `tfsdk:"description"`
	Name          types.String `tfsdk:"name"`
	NameAlias     types.String `tfsdk:"name_alias"`
	OwnerKey      types.String `tfsdk:"owner_key"`
	OwnerTag      types.String `tfsdk:"owner_tag"`
	Tag           types.String `tfsdk:"tag"`
	TagAnnotation types.Set    `tfsdk:"annotations"`
	TagTag        types.Set    `tfsdk:"tags"`
}

func getEmptyL3extProvLblResourceModel() *L3extProvLblResourceModel {
	return &L3extProvLblResourceModel{
		Id:         basetypes.NewStringNull(),
		ParentDn:   basetypes.NewStringNull(),
		Annotation: basetypes.NewStringNull(),
		Descr:      basetypes.NewStringNull(),
		Name:       basetypes.NewStringNull(),
		NameAlias:  basetypes.NewStringNull(),
		OwnerKey:   basetypes.NewStringNull(),
		OwnerTag:   basetypes.NewStringNull(),
		Tag:        basetypes.NewStringNull(),
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

// TagAnnotationL3extProvLblResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationL3extProvLblResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func getEmptyTagAnnotationL3extProvLblResourceModel() TagAnnotationL3extProvLblResourceModel {
	return TagAnnotationL3extProvLblResourceModel{
		Key:   basetypes.NewStringNull(),
		Value: basetypes.NewStringNull(),
	}
}

var TagAnnotationL3extProvLblType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"key":   types.StringType,
		"value": types.StringType,
	},
}

// TagTagL3extProvLblResourceModel describes the resource data model for the children without relation ships.
type TagTagL3extProvLblResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func getEmptyTagTagL3extProvLblResourceModel() TagTagL3extProvLblResourceModel {
	return TagTagL3extProvLblResourceModel{
		Key:   basetypes.NewStringNull(),
		Value: basetypes.NewStringNull(),
	}
}

var TagTagL3extProvLblType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"key":   types.StringType,
		"value": types.StringType,
	},
}

type L3extProvLblIdentifier struct {
	Name types.String
}

func (r *L3extProvLblResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if !req.Plan.Raw.IsNull() {
		var planData, stateData *L3extProvLblResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

		if resp.Diagnostics.HasError() {
			return
		}

		if (planData.Id.IsUnknown() || planData.Id.IsNull()) && !planData.ParentDn.IsUnknown() && !planData.Name.IsUnknown() {
			setL3extProvLblId(ctx, planData)
		}

		if stateData == nil && !globalAllowExistingOnCreate && !planData.Id.IsUnknown() && !planData.Id.IsNull() {
			CheckDn(ctx, &resp.Diagnostics, r.client, "l3extProvLbl", planData.Id.ValueString())
			if resp.Diagnostics.HasError() {
				return
			}
		}

		resp.Diagnostics.Append(resp.Plan.Set(ctx, &planData)...)
	}
}

func (r *L3extProvLblResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: aci_l3out_provider_label")
	resp.TypeName = req.ProviderTypeName + "_l3out_provider_label"
	tflog.Debug(ctx, "End metadata of resource: aci_l3out_provider_label")
}

func (r *L3extProvLblResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: aci_l3out_provider_label")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The l3out_provider_label resource for the 'l3extProvLbl' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the L3Out Provider Label object.",
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
				MarkdownDescription: `The annotation of the L3Out Provider Label object.`,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The description of the L3Out Provider Label object.`,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The name of the L3Out Provider Label object.`,
			},
			"name_alias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The name alias of the L3Out Provider Label object.`,
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
			"tag": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("alice-blue", "antique-white", "aqua", "aquamarine", "azure", "beige", "bisque", "black", "blanched-almond", "blue", "blue-violet", "brown", "burlywood", "cadet-blue", "chartreuse", "chocolate", "coral", "cornflower-blue", "cornsilk", "crimson", "cyan", "dark-blue", "dark-cyan", "dark-goldenrod", "dark-gray", "dark-green", "dark-khaki", "dark-magenta", "dark-olive-green", "dark-orange", "dark-orchid", "dark-red", "dark-salmon", "dark-sea-green", "dark-slate-blue", "dark-slate-gray", "dark-turquoise", "dark-violet", "deep-pink", "deep-sky-blue", "dim-gray", "dodger-blue", "fire-brick", "floral-white", "forest-green", "fuchsia", "gainsboro", "ghost-white", "gold", "goldenrod", "gray", "green", "green-yellow", "honeydew", "hot-pink", "indian-red", "indigo", "ivory", "khaki", "lavender", "lavender-blush", "lawn-green", "lemon-chiffon", "light-blue", "light-coral", "light-cyan", "light-goldenrod-yellow", "light-gray", "light-green", "light-pink", "light-salmon", "light-sea-green", "light-sky-blue", "light-slate-gray", "light-steel-blue", "light-yellow", "lime", "lime-green", "linen", "magenta", "maroon", "medium-aquamarine", "medium-blue", "medium-orchid", "medium-purple", "medium-sea-green", "medium-slate-blue", "medium-spring-green", "medium-turquoise", "medium-violet-red", "midnight-blue", "mint-cream", "misty-rose", "moccasin", "navajo-white", "navy", "old-lace", "olive", "olive-drab", "orange", "orange-red", "orchid", "pale-goldenrod", "pale-green", "pale-turquoise", "pale-violet-red", "papaya-whip", "peachpuff", "peru", "pink", "plum", "powder-blue", "purple", "red", "rosy-brown", "royal-blue", "saddle-brown", "salmon", "sandy-brown", "sea-green", "seashell", "sienna", "silver", "sky-blue", "slate-blue", "slate-gray", "snow", "spring-green", "steel-blue", "tan", "teal", "thistle", "tomato", "turquoise", "violet", "wheat", "white", "white-smoke", "yellow", "yellow-green"),
				},
				MarkdownDescription: `Specifies the color of a policy label.`,
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
	tflog.Debug(ctx, "End schema of resource: aci_l3out_provider_label")
}

func (r *L3extProvLblResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: aci_l3out_provider_label")
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
	tflog.Debug(ctx, "End configure of resource: aci_l3out_provider_label")
}

func (r *L3extProvLblResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: aci_l3out_provider_label")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *L3extProvLblResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	if stateData.Id.IsUnknown() || stateData.Id.IsNull() {
		setL3extProvLblId(ctx, stateData)
	}
	getAndSetL3extProvLblAttributes(ctx, &resp.Diagnostics, r.client, stateData)
	if !globalAllowExistingOnCreate && !stateData.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Object Already Exists",
			fmt.Sprintf("The l3extProvLbl object with DN '%s' already exists.", stateData.Id.ValueString()),
		)
		return
	}

	var data *L3extProvLblResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsUnknown() || data.Id.IsNull() {
		setL3extProvLblId(ctx, data)
	}

	tflog.Debug(ctx, fmt.Sprintf("Create of resource aci_l3out_provider_label with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationL3extProvLblResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagL3extProvLblResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getL3extProvLblCreateJsonPayload(ctx, &resp.Diagnostics, true, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetL3extProvLblAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End create of resource aci_l3out_provider_label with id '%s'", data.Id.ValueString()))
}

func (r *L3extProvLblResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: aci_l3out_provider_label")
	var data *L3extProvLblResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource aci_l3out_provider_label with id '%s'", data.Id.ValueString()))

	getAndSetL3extProvLblAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *L3extProvLblResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, fmt.Sprintf("End read of resource aci_l3out_provider_label with id '%s'", data.Id.ValueString()))
}

func (r *L3extProvLblResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: aci_l3out_provider_label")
	var data *L3extProvLblResourceModel
	var stateData *L3extProvLblResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource aci_l3out_provider_label with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationL3extProvLblResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagL3extProvLblResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getL3extProvLblCreateJsonPayload(ctx, &resp.Diagnostics, false, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetL3extProvLblAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End update of resource aci_l3out_provider_label with id '%s'", data.Id.ValueString()))
}

func (r *L3extProvLblResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: aci_l3out_provider_label")
	var data *L3extProvLblResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource aci_l3out_provider_label with id '%s'", data.Id.ValueString()))
	jsonPayload := GetDeleteJsonPayload(ctx, &resp.Diagnostics, "l3extProvLbl", data.Id.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}
	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("End delete of resource aci_l3out_provider_label with id '%s'", data.Id.ValueString()))
}

func (r *L3extProvLblResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: aci_l3out_provider_label")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *L3extProvLblResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)
	tflog.Debug(ctx, fmt.Sprintf("Import state of resource aci_l3out_provider_label with id '%s'", stateData.Id.ValueString()))

	tflog.Debug(ctx, "End import of state resource: aci_l3out_provider_label")
}

func getAndSetL3extProvLblAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *L3extProvLblResourceModel) {
	requestData := DoRestRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=full&rsp-subtree-class=%s", data.Id.ValueString(), "l3extProvLbl,tagAnnotation,tagTag"), "GET", nil)

	readData := getEmptyL3extProvLblResourceModel()

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("l3extProvLbl").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("l3extProvLbl").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					readData.Id = basetypes.NewStringValue(attributeValue.(string))
					setL3extProvLblParentDn(ctx, attributeValue.(string), readData)
				}
				if attributeName == "annotation" {
					readData.Annotation = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "descr" {
					readData.Descr = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "name" {
					readData.Name = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "nameAlias" {
					readData.NameAlias = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ownerKey" {
					readData.OwnerKey = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ownerTag" {
					readData.OwnerTag = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "tag" {
					readData.Tag = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			TagAnnotationL3extProvLblList := make([]TagAnnotationL3extProvLblResourceModel, 0)
			TagTagL3extProvLblList := make([]TagTagL3extProvLblResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "tagAnnotation" {
							TagAnnotationL3extProvLbl := getEmptyTagAnnotationL3extProvLblResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationL3extProvLbl.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationL3extProvLbl.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}

							}
							TagAnnotationL3extProvLblList = append(TagAnnotationL3extProvLblList, TagAnnotationL3extProvLbl)
						}
						if childClassName == "tagTag" {
							TagTagL3extProvLbl := getEmptyTagTagL3extProvLblResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagTagL3extProvLbl.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagTagL3extProvLbl.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}

							}
							TagTagL3extProvLblList = append(TagTagL3extProvLblList, TagTagL3extProvLbl)
						}
					}
				}
			}
			tagAnnotationSet, _ := types.SetValueFrom(ctx, readData.TagAnnotation.ElementType(ctx), TagAnnotationL3extProvLblList)
			readData.TagAnnotation = tagAnnotationSet
			tagTagSet, _ := types.SetValueFrom(ctx, readData.TagTag.ElementType(ctx), TagTagL3extProvLblList)
			readData.TagTag = tagTagSet
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'l3extProvLbl'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	} else {
		readData.Id = basetypes.NewStringNull()
	}
	*data = *readData
}

func getL3extProvLblRn(ctx context.Context, data *L3extProvLblResourceModel) string {
	return fmt.Sprintf("provlbl-%s", data.Name.ValueString())
}

func setL3extProvLblParentDn(ctx context.Context, dn string, data *L3extProvLblResourceModel) {
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

func setL3extProvLblId(ctx context.Context, data *L3extProvLblResourceModel) {
	rn := getL3extProvLblRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
}

func getL3extProvLblTagAnnotationChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *L3extProvLblResourceModel, tagAnnotationL3extProvLblPlan, tagAnnotationL3extProvLblState []TagAnnotationL3extProvLblResourceModel) []map[string]interface{} {
	childPayloads := []map[string]interface{}{}
	if !data.TagAnnotation.IsNull() && !data.TagAnnotation.IsUnknown() {
		tagAnnotationIdentifiers := []TagAnnotationIdentifier{}
		for _, tagAnnotationL3extProvLbl := range tagAnnotationL3extProvLblPlan {
			childMap := NewAciObject()
			if !tagAnnotationL3extProvLbl.Key.IsNull() && !tagAnnotationL3extProvLbl.Key.IsUnknown() {
				childMap.Attributes["key"] = tagAnnotationL3extProvLbl.Key.ValueString()
			}
			if !tagAnnotationL3extProvLbl.Value.IsNull() && !tagAnnotationL3extProvLbl.Value.IsUnknown() {
				childMap.Attributes["value"] = tagAnnotationL3extProvLbl.Value.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"tagAnnotation": childMap})
			tagAnnotationIdentifier := TagAnnotationIdentifier{}
			tagAnnotationIdentifier.Key = tagAnnotationL3extProvLbl.Key
			tagAnnotationIdentifiers = append(tagAnnotationIdentifiers, tagAnnotationIdentifier)
		}
		for _, tagAnnotation := range tagAnnotationL3extProvLblState {
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

func getL3extProvLblTagTagChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *L3extProvLblResourceModel, tagTagL3extProvLblPlan, tagTagL3extProvLblState []TagTagL3extProvLblResourceModel) []map[string]interface{} {
	childPayloads := []map[string]interface{}{}
	if !data.TagTag.IsNull() && !data.TagTag.IsUnknown() {
		tagTagIdentifiers := []TagTagIdentifier{}
		for _, tagTagL3extProvLbl := range tagTagL3extProvLblPlan {
			childMap := NewAciObject()
			if !tagTagL3extProvLbl.Key.IsNull() && !tagTagL3extProvLbl.Key.IsUnknown() {
				childMap.Attributes["key"] = tagTagL3extProvLbl.Key.ValueString()
			}
			if !tagTagL3extProvLbl.Value.IsNull() && !tagTagL3extProvLbl.Value.IsUnknown() {
				childMap.Attributes["value"] = tagTagL3extProvLbl.Value.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"tagTag": childMap})
			tagTagIdentifier := TagTagIdentifier{}
			tagTagIdentifier.Key = tagTagL3extProvLbl.Key
			tagTagIdentifiers = append(tagTagIdentifiers, tagTagIdentifier)
		}
		for _, tagTag := range tagTagL3extProvLblState {
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

func getL3extProvLblCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, createType bool, data *L3extProvLblResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationL3extProvLblResourceModel, tagTagPlan, tagTagState []TagTagL3extProvLblResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}

	if createType && !globalAllowExistingOnCreate {
		payloadMap["attributes"].(map[string]string)["status"] = "created"
	}
	childPayloads := []map[string]interface{}{}

	TagAnnotationchildPayloads := getL3extProvLblTagAnnotationChildPayloads(ctx, diags, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)

	TagTagchildPayloads := getL3extProvLblTagTagChildPayloads(ctx, diags, data, tagTagPlan, tagTagState)
	if TagTagchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagTagchildPayloads...)

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
	if !data.OwnerKey.IsNull() && !data.OwnerKey.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["ownerKey"] = data.OwnerKey.ValueString()
	}
	if !data.OwnerTag.IsNull() && !data.OwnerTag.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["ownerTag"] = data.OwnerTag.ValueString()
	}
	if !data.Tag.IsNull() && !data.Tag.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["tag"] = data.Tag.ValueString()
	}
	payload, err := json.Marshal(map[string]interface{}{"l3extProvLbl": payloadMap})
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
