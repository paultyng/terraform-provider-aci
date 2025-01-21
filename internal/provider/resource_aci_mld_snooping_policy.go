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
var _ resource.Resource = &MldSnoopPolResource{}
var _ resource.ResourceWithImportState = &MldSnoopPolResource{}

func NewMldSnoopPolResource() resource.Resource {
	return &MldSnoopPolResource{}
}

// MldSnoopPolResource defines the resource implementation.
type MldSnoopPolResource struct {
	client *client.Client
}

// MldSnoopPolResourceModel describes the resource data model.
type MldSnoopPolResourceModel struct {
	Id              types.String `tfsdk:"id"`
	ParentDn        types.String `tfsdk:"parent_dn"`
	AdminSt         types.String `tfsdk:"admin_state"`
	Annotation      types.String `tfsdk:"annotation"`
	Ctrl            types.Set    `tfsdk:"control"`
	Descr           types.String `tfsdk:"description"`
	LastMbrIntvl    types.String `tfsdk:"last_member_interval"`
	Name            types.String `tfsdk:"name"`
	NameAlias       types.String `tfsdk:"name_alias"`
	OwnerKey        types.String `tfsdk:"owner_key"`
	OwnerTag        types.String `tfsdk:"owner_tag"`
	QueryIntvl      types.String `tfsdk:"query_interval"`
	RspIntvl        types.String `tfsdk:"response_interval"`
	StartQueryCnt   types.String `tfsdk:"start_query_count"`
	StartQueryIntvl types.String `tfsdk:"start_query_interval"`
	Ver             types.String `tfsdk:"version"`
	TagAnnotation   types.Set    `tfsdk:"annotations"`
	TagTag          types.Set    `tfsdk:"tags"`
}

func getEmptyMldSnoopPolResourceModel() *MldSnoopPolResourceModel {
	return &MldSnoopPolResourceModel{
		Id:              basetypes.NewStringNull(),
		ParentDn:        basetypes.NewStringNull(),
		AdminSt:         basetypes.NewStringNull(),
		Annotation:      basetypes.NewStringNull(),
		Ctrl:            types.SetNull(types.StringType),
		Descr:           basetypes.NewStringNull(),
		LastMbrIntvl:    basetypes.NewStringNull(),
		Name:            basetypes.NewStringNull(),
		NameAlias:       basetypes.NewStringNull(),
		OwnerKey:        basetypes.NewStringNull(),
		OwnerTag:        basetypes.NewStringNull(),
		QueryIntvl:      basetypes.NewStringNull(),
		RspIntvl:        basetypes.NewStringNull(),
		StartQueryCnt:   basetypes.NewStringNull(),
		StartQueryIntvl: basetypes.NewStringNull(),
		Ver:             basetypes.NewStringNull(),
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

// TagAnnotationMldSnoopPolResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationMldSnoopPolResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func getEmptyTagAnnotationMldSnoopPolResourceModel() TagAnnotationMldSnoopPolResourceModel {
	return TagAnnotationMldSnoopPolResourceModel{
		Key:   basetypes.NewStringNull(),
		Value: basetypes.NewStringNull(),
	}
}

var TagAnnotationMldSnoopPolType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"key":   types.StringType,
		"value": types.StringType,
	},
}

// TagTagMldSnoopPolResourceModel describes the resource data model for the children without relation ships.
type TagTagMldSnoopPolResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func getEmptyTagTagMldSnoopPolResourceModel() TagTagMldSnoopPolResourceModel {
	return TagTagMldSnoopPolResourceModel{
		Key:   basetypes.NewStringNull(),
		Value: basetypes.NewStringNull(),
	}
}

var TagTagMldSnoopPolType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"key":   types.StringType,
		"value": types.StringType,
	},
}

type MldSnoopPolIdentifier struct {
	Name types.String
}

func (r *MldSnoopPolResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if !req.Plan.Raw.IsNull() {
		var planData, stateData *MldSnoopPolResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

		if resp.Diagnostics.HasError() {
			return
		}

		if (planData.Id.IsUnknown() || planData.Id.IsNull()) && !planData.ParentDn.IsUnknown() && !planData.Name.IsUnknown() {
			setMldSnoopPolId(ctx, planData)
		}

		if stateData == nil && !globalAllowExistingOnCreate && !planData.Id.IsUnknown() && !planData.Id.IsNull() {
			CheckDn(ctx, &resp.Diagnostics, r.client, "mldSnoopPol", planData.Id.ValueString())
			if resp.Diagnostics.HasError() {
				return
			}
		}

		resp.Diagnostics.Append(resp.Plan.Set(ctx, &planData)...)
	}
}

func (r *MldSnoopPolResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: aci_mld_snooping_policy")
	resp.TypeName = req.ProviderTypeName + "_mld_snooping_policy"
	tflog.Debug(ctx, "End metadata of resource: aci_mld_snooping_policy")
}

func (r *MldSnoopPolResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: aci_mld_snooping_policy")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The mld_snooping_policy resource for the 'mldSnoopPol' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the MLD Snooping Policy object.",
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
			"admin_state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("disabled", "enabled"),
				},
				MarkdownDescription: `The administrative state of the MLD Snooping Policy object.`,
			},
			"annotation": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Default:             stringdefault.StaticString(globalAnnotation),
				MarkdownDescription: `The annotation of the MLD Snooping Policy object.`,
			},
			"control": schema.SetAttribute{
				MarkdownDescription: `The controls for the MLD Snooping Policy object.`,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
					SetToSetNullWhenStateIsNullPlanIsUnknownDuringUpdate(nil),
				},
				Validators: []validator.Set{
					setvalidator.SizeAtMost(4),
					setvalidator.ValueStringsAre(
						stringvalidator.OneOf("fast-leave", "opt-flood", "querier", "routing"),
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
				MarkdownDescription: `The description of the MLD Snooping Policy object.`,
			},
			"last_member_interval": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The last member interval (seconds) of the MLD Snooping Policy object. The group state is removed when no host responds before the timeout.`,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The name of the MLD Snooping Policy object.`,
			},
			"name_alias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The name alias of the MLD Snooping Policy object.`,
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
			"query_interval": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The query interval (seconds) of the MLD Snooping Policy object.`,
			},
			"response_interval": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The response interval (seconds) of the MLD Snooping Policy object.`,
			},
			"start_query_count": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The start query count of the MLD Snooping Policy object.`,
			},
			"start_query_interval": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The query interval (seconds) of the MLD Snooping Policy object at start-up.`,
			},
			"version": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("unspecified", "v1", "v2"),
				},
				MarkdownDescription: `The MLD version.`,
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
	tflog.Debug(ctx, "End schema of resource: aci_mld_snooping_policy")
}

func (r *MldSnoopPolResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: aci_mld_snooping_policy")
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
	tflog.Debug(ctx, "End configure of resource: aci_mld_snooping_policy")
}

func (r *MldSnoopPolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: aci_mld_snooping_policy")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *MldSnoopPolResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	if stateData.Id.IsUnknown() || stateData.Id.IsNull() {
		setMldSnoopPolId(ctx, stateData)
	}
	getAndSetMldSnoopPolAttributes(ctx, &resp.Diagnostics, r.client, stateData)
	if !globalAllowExistingOnCreate && !stateData.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Object Already Exists",
			fmt.Sprintf("The mldSnoopPol object with DN '%s' already exists.", stateData.Id.ValueString()),
		)
		return
	}

	var data *MldSnoopPolResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsUnknown() || data.Id.IsNull() {
		setMldSnoopPolId(ctx, data)
	}

	tflog.Debug(ctx, fmt.Sprintf("Create of resource aci_mld_snooping_policy with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationMldSnoopPolResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagMldSnoopPolResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getMldSnoopPolCreateJsonPayload(ctx, &resp.Diagnostics, true, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetMldSnoopPolAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End create of resource aci_mld_snooping_policy with id '%s'", data.Id.ValueString()))
}

func (r *MldSnoopPolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: aci_mld_snooping_policy")
	var data *MldSnoopPolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource aci_mld_snooping_policy with id '%s'", data.Id.ValueString()))

	getAndSetMldSnoopPolAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *MldSnoopPolResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, fmt.Sprintf("End read of resource aci_mld_snooping_policy with id '%s'", data.Id.ValueString()))
}

func (r *MldSnoopPolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: aci_mld_snooping_policy")
	var data *MldSnoopPolResourceModel
	var stateData *MldSnoopPolResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource aci_mld_snooping_policy with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationMldSnoopPolResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagMldSnoopPolResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getMldSnoopPolCreateJsonPayload(ctx, &resp.Diagnostics, false, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetMldSnoopPolAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End update of resource aci_mld_snooping_policy with id '%s'", data.Id.ValueString()))
}

func (r *MldSnoopPolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: aci_mld_snooping_policy")
	var data *MldSnoopPolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource aci_mld_snooping_policy with id '%s'", data.Id.ValueString()))
	jsonPayload := GetDeleteJsonPayload(ctx, &resp.Diagnostics, "mldSnoopPol", data.Id.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}
	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("End delete of resource aci_mld_snooping_policy with id '%s'", data.Id.ValueString()))
}

func (r *MldSnoopPolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: aci_mld_snooping_policy")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *MldSnoopPolResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)
	tflog.Debug(ctx, fmt.Sprintf("Import state of resource aci_mld_snooping_policy with id '%s'", stateData.Id.ValueString()))

	tflog.Debug(ctx, "End import of state resource: aci_mld_snooping_policy")
}

func getAndSetMldSnoopPolAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *MldSnoopPolResourceModel) {
	requestData := DoRestRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=full&rsp-subtree-class=%s", data.Id.ValueString(), "mldSnoopPol,tagAnnotation,tagTag"), "GET", nil)

	readData := getEmptyMldSnoopPolResourceModel()

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("mldSnoopPol").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("mldSnoopPol").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					readData.Id = basetypes.NewStringValue(attributeValue.(string))
					setMldSnoopPolParentDn(ctx, attributeValue.(string), readData)
				}
				if attributeName == "adminSt" {
					readData.AdminSt = basetypes.NewStringValue(attributeValue.(string))
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
				if attributeName == "lastMbrIntvl" {
					readData.LastMbrIntvl = basetypes.NewStringValue(attributeValue.(string))
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
				if attributeName == "queryIntvl" {
					readData.QueryIntvl = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "rspIntvl" {
					readData.RspIntvl = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "startQueryCnt" {
					readData.StartQueryCnt = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "startQueryIntvl" {
					readData.StartQueryIntvl = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ver" {
					readData.Ver = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			TagAnnotationMldSnoopPolList := make([]TagAnnotationMldSnoopPolResourceModel, 0)
			TagTagMldSnoopPolList := make([]TagTagMldSnoopPolResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "tagAnnotation" {
							TagAnnotationMldSnoopPol := getEmptyTagAnnotationMldSnoopPolResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationMldSnoopPol.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationMldSnoopPol.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}

							}
							TagAnnotationMldSnoopPolList = append(TagAnnotationMldSnoopPolList, TagAnnotationMldSnoopPol)
						}
						if childClassName == "tagTag" {
							TagTagMldSnoopPol := getEmptyTagTagMldSnoopPolResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagTagMldSnoopPol.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagTagMldSnoopPol.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}

							}
							TagTagMldSnoopPolList = append(TagTagMldSnoopPolList, TagTagMldSnoopPol)
						}
					}
				}
			}
			tagAnnotationSet, _ := types.SetValueFrom(ctx, readData.TagAnnotation.ElementType(ctx), TagAnnotationMldSnoopPolList)
			readData.TagAnnotation = tagAnnotationSet
			tagTagSet, _ := types.SetValueFrom(ctx, readData.TagTag.ElementType(ctx), TagTagMldSnoopPolList)
			readData.TagTag = tagTagSet
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'mldSnoopPol'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	} else {
		readData.Id = basetypes.NewStringNull()
	}
	*data = *readData
}

func getMldSnoopPolRn(ctx context.Context, data *MldSnoopPolResourceModel) string {
	return fmt.Sprintf("mldsnoopPol-%s", data.Name.ValueString())
}

func setMldSnoopPolParentDn(ctx context.Context, dn string, data *MldSnoopPolResourceModel) {
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

func setMldSnoopPolId(ctx context.Context, data *MldSnoopPolResourceModel) {
	rn := getMldSnoopPolRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
}

func getMldSnoopPolTagAnnotationChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *MldSnoopPolResourceModel, tagAnnotationMldSnoopPolPlan, tagAnnotationMldSnoopPolState []TagAnnotationMldSnoopPolResourceModel) []map[string]interface{} {
	childPayloads := []map[string]interface{}{}
	if !data.TagAnnotation.IsNull() && !data.TagAnnotation.IsUnknown() {
		tagAnnotationIdentifiers := []TagAnnotationIdentifier{}
		for _, tagAnnotationMldSnoopPol := range tagAnnotationMldSnoopPolPlan {
			childMap := NewAciObject()
			if !tagAnnotationMldSnoopPol.Key.IsNull() && !tagAnnotationMldSnoopPol.Key.IsUnknown() {
				childMap.Attributes["key"] = tagAnnotationMldSnoopPol.Key.ValueString()
			}
			if !tagAnnotationMldSnoopPol.Value.IsNull() && !tagAnnotationMldSnoopPol.Value.IsUnknown() {
				childMap.Attributes["value"] = tagAnnotationMldSnoopPol.Value.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"tagAnnotation": childMap})
			tagAnnotationIdentifier := TagAnnotationIdentifier{}
			tagAnnotationIdentifier.Key = tagAnnotationMldSnoopPol.Key
			tagAnnotationIdentifiers = append(tagAnnotationIdentifiers, tagAnnotationIdentifier)
		}
		for _, tagAnnotation := range tagAnnotationMldSnoopPolState {
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

func getMldSnoopPolTagTagChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *MldSnoopPolResourceModel, tagTagMldSnoopPolPlan, tagTagMldSnoopPolState []TagTagMldSnoopPolResourceModel) []map[string]interface{} {
	childPayloads := []map[string]interface{}{}
	if !data.TagTag.IsNull() && !data.TagTag.IsUnknown() {
		tagTagIdentifiers := []TagTagIdentifier{}
		for _, tagTagMldSnoopPol := range tagTagMldSnoopPolPlan {
			childMap := NewAciObject()
			if !tagTagMldSnoopPol.Key.IsNull() && !tagTagMldSnoopPol.Key.IsUnknown() {
				childMap.Attributes["key"] = tagTagMldSnoopPol.Key.ValueString()
			}
			if !tagTagMldSnoopPol.Value.IsNull() && !tagTagMldSnoopPol.Value.IsUnknown() {
				childMap.Attributes["value"] = tagTagMldSnoopPol.Value.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"tagTag": childMap})
			tagTagIdentifier := TagTagIdentifier{}
			tagTagIdentifier.Key = tagTagMldSnoopPol.Key
			tagTagIdentifiers = append(tagTagIdentifiers, tagTagIdentifier)
		}
		for _, tagTag := range tagTagMldSnoopPolState {
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

func getMldSnoopPolCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, createType bool, data *MldSnoopPolResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationMldSnoopPolResourceModel, tagTagPlan, tagTagState []TagTagMldSnoopPolResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}

	if createType && !globalAllowExistingOnCreate {
		payloadMap["attributes"].(map[string]string)["status"] = "created"
	}
	childPayloads := []map[string]interface{}{}

	TagAnnotationchildPayloads := getMldSnoopPolTagAnnotationChildPayloads(ctx, diags, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)

	TagTagchildPayloads := getMldSnoopPolTagTagChildPayloads(ctx, diags, data, tagTagPlan, tagTagState)
	if TagTagchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagTagchildPayloads...)

	payloadMap["children"] = childPayloads
	if !data.AdminSt.IsNull() && !data.AdminSt.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["adminSt"] = data.AdminSt.ValueString()
	}
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
	if !data.LastMbrIntvl.IsNull() && !data.LastMbrIntvl.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["lastMbrIntvl"] = data.LastMbrIntvl.ValueString()
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
	if !data.QueryIntvl.IsNull() && !data.QueryIntvl.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["queryIntvl"] = data.QueryIntvl.ValueString()
	}
	if !data.RspIntvl.IsNull() && !data.RspIntvl.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["rspIntvl"] = data.RspIntvl.ValueString()
	}
	if !data.StartQueryCnt.IsNull() && !data.StartQueryCnt.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["startQueryCnt"] = data.StartQueryCnt.ValueString()
	}
	if !data.StartQueryIntvl.IsNull() && !data.StartQueryIntvl.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["startQueryIntvl"] = data.StartQueryIntvl.ValueString()
	}
	if !data.Ver.IsNull() && !data.Ver.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["ver"] = data.Ver.ValueString()
	}
	payload, err := json.Marshal(map[string]interface{}{"mldSnoopPol": payloadMap})
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
