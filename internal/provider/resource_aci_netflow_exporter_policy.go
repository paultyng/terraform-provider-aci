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

	"github.com/CiscoDevNet/terraform-provider-aci/v2/internal/validators"
	"github.com/CiscoDevNet/terraform-provider-aci/v2/internal/customtypes"
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
var _ resource.Resource = &NetflowExporterPolResource{}
var _ resource.ResourceWithImportState = &NetflowExporterPolResource{}

func NewNetflowExporterPolResource() resource.Resource {
	return &NetflowExporterPolResource{}
}

// NetflowExporterPolResource defines the resource implementation.
type NetflowExporterPolResource struct {
	client *client.Client
}

// NetflowExporterPolResourceModel describes the resource data model.
type NetflowExporterPolResourceModel struct {
	Id            types.String `tfsdk:"id"`
	ParentDn      types.String `tfsdk:"parent_dn"`
	Annotation    types.String `tfsdk:"annotation"`
	Descr         types.String `tfsdk:"description"`
	Dscp          types.String `tfsdk:"dscp"`
	DstAddr       types.String `tfsdk:"destination_address"`
	DstPort       customtypes.DstPortStringValue `tfsdk:"destination_port"`
	Name          types.String `tfsdk:"name"`
	NameAlias     types.String `tfsdk:"name_alias"`
	OwnerKey      types.String `tfsdk:"owner_key"`
	OwnerTag      types.String `tfsdk:"owner_tag"`
	SourceIpType  types.String `tfsdk:"source_ip_type"`
	SrcAddr       types.String `tfsdk:"source_address"`
	Ver           types.String `tfsdk:"version"`
	TagAnnotation types.Set    `tfsdk:"annotations"`
	TagTag        types.Set    `tfsdk:"tags"`
}

// TagAnnotationNetflowExporterPolResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationNetflowExporterPolResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

// TagTagNetflowExporterPolResourceModel describes the resource data model for the children without relation ships.
type TagTagNetflowExporterPolResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type NetflowExporterPolIdentifier struct {
	Name types.String
}

func (r *NetflowExporterPolResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if !req.Plan.Raw.IsNull() {
		var planData, stateData *NetflowExporterPolResourceModel
		// var DstPortValue string
		// matchMappingDstPort := map[string]string{"443": "https"}
		// req.Plan.GetAttribute(ctx, path.Root("DstPort"), DstPortValue)
		// if value, ok := matchMappingDstPort[DstPortValue]; ok {
		// 	req.Plan.SetAttribute(ctx, path.Root("DstPort"), value)
		// }
		resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

		if resp.Diagnostics.HasError() {
			return
		}

		if (planData.Id.IsUnknown() || planData.Id.IsNull()) && !planData.ParentDn.IsUnknown() && !planData.Name.IsUnknown() {
			setNetflowExporterPolId(ctx, planData)
		}

		if stateData == nil && !globalAllowExistingOnCreate && !planData.Id.IsUnknown() && !planData.Id.IsNull() {
			CheckDn(ctx, &resp.Diagnostics, r.client, "netflowExporterPol", planData.Id.ValueString())
			if resp.Diagnostics.HasError() {
				return
			}
		}

		resp.Diagnostics.Append(resp.Plan.Set(ctx, &planData)...)
	}
}

func (r *NetflowExporterPolResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: aci_netflow_exporter_policy")
	resp.TypeName = req.ProviderTypeName + "_netflow_exporter_policy"
	tflog.Debug(ctx, "End metadata of resource: aci_netflow_exporter_policy")
}

func (r *NetflowExporterPolResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: aci_netflow_exporter_policy")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The netflow_exporter_policy resource for the 'netflowExporterPol' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the Netflow Exporter Policy object.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"parent_dn": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("uni/infra"),
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
				},
				Default:             stringdefault.StaticString(globalAnnotation),
				MarkdownDescription: `The annotation of the Netflow Exporter Policy object.`,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `The description of the Netflow Exporter Policy object.`,
			},
			"dscp": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("AF11", "AF12", "AF13", "AF21", "AF22", "AF23", "AF31", "AF32", "AF33", "AF41", "AF42", "AF43", "CS0", "CS1", "CS2", "CS3", "CS4", "CS5", "CS6", "CS7", "EF", "VA"),
				},
				MarkdownDescription: `The DSCP value of the Netflow Exporter Policy object.`,
			},
			"destination_address": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `The destination IP address of the remote node.`,
			},
			"destination_port": schema.StringAttribute{
				CustomType: customtypes.DstPortStringType{},
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.Any(
						stringvalidator.OneOf("dns", "ftpData", "http", "https", "pop3", "rtsp", "smtp", "ssh", "unspecified"),
						validators.InBetweenFromString(0, 65535),
					),
				},
				MarkdownDescription: `The destination port of the remote node.`,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The name of the Netflow Exporter Policy object.`,
			},
			"name_alias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `The name alias of the Netflow Exporter Policy object.`,
			},
			"owner_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `The key for enabling clients to own their data for entity correlation.`,
			},
			"owner_tag": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `A tag for enabling clients to add their own data. For example, to indicate who created this object.`,
			},
			"source_ip_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("custom-src-ip", "inband-mgmt-ip", "oob-mgmt-ip", "ptep"),
				},
				MarkdownDescription: `The type of the source IP address: It can be one of the available management IP address for a given leaf or a custom IP Address.`,
			},
			"source_address": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `The source IP address.`,
			},
			"version": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("cisco-v1", "v5", "v9"),
				},
				MarkdownDescription: `The collector version.`,
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
	tflog.Debug(ctx, "End schema of resource: aci_netflow_exporter_policy")
}

func (r *NetflowExporterPolResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: aci_netflow_exporter_policy")
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
	tflog.Debug(ctx, "End configure of resource: aci_netflow_exporter_policy")
}

func (r *NetflowExporterPolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: aci_netflow_exporter_policy")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *NetflowExporterPolResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	if stateData.Id.IsUnknown() || stateData.Id.IsNull() {
		setNetflowExporterPolId(ctx, stateData)
	}
	getAndSetNetflowExporterPolAttributes(ctx, &resp.Diagnostics, r.client, stateData)
	if !globalAllowExistingOnCreate && !stateData.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Object Already Exists",
			fmt.Sprintf("The netflowExporterPol object with DN '%s' already exists.", stateData.Id.ValueString()),
		)
		return
	}

	var data *NetflowExporterPolResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsUnknown() || data.Id.IsNull() {
		setNetflowExporterPolId(ctx, data)
	}

	tflog.Debug(ctx, fmt.Sprintf("Create of resource aci_netflow_exporter_policy with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationNetflowExporterPolResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagNetflowExporterPolResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getNetflowExporterPolCreateJsonPayload(ctx, &resp.Diagnostics, true, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetNetflowExporterPolAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End create of resource aci_netflow_exporter_policy with id '%s'", data.Id.ValueString()))
}

func (r *NetflowExporterPolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: aci_netflow_exporter_policy")
	var data *NetflowExporterPolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource aci_netflow_exporter_policy with id '%s'", data.Id.ValueString()))

	getAndSetNetflowExporterPolAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *NetflowExporterPolResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, fmt.Sprintf("End read of resource aci_netflow_exporter_policy with id '%s'", data.Id.ValueString()))
}

func (r *NetflowExporterPolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: aci_netflow_exporter_policy")
	var data *NetflowExporterPolResourceModel
	var stateData *NetflowExporterPolResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource aci_netflow_exporter_policy with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationNetflowExporterPolResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagNetflowExporterPolResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getNetflowExporterPolCreateJsonPayload(ctx, &resp.Diagnostics, false, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetNetflowExporterPolAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End update of resource aci_netflow_exporter_policy with id '%s'", data.Id.ValueString()))
}

func (r *NetflowExporterPolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: aci_netflow_exporter_policy")
	var data *NetflowExporterPolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource aci_netflow_exporter_policy with id '%s'", data.Id.ValueString()))
	jsonPayload := GetDeleteJsonPayload(ctx, &resp.Diagnostics, "netflowExporterPol", data.Id.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}
	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("End delete of resource aci_netflow_exporter_policy with id '%s'", data.Id.ValueString()))
}

func (r *NetflowExporterPolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: aci_netflow_exporter_policy")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *NetflowExporterPolResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)
	tflog.Debug(ctx, fmt.Sprintf("Import state of resource aci_netflow_exporter_policy with id '%s'", stateData.Id.ValueString()))

	tflog.Debug(ctx, "End import of state resource: aci_netflow_exporter_policy")
}

func getAndSetNetflowExporterPolAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *NetflowExporterPolResourceModel) {
	requestData := DoRestRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=children&rsp-subtree-class=%s", data.Id.ValueString(), "netflowExporterPol,tagAnnotation,tagTag"), "GET", nil)

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("netflowExporterPol").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("netflowExporterPol").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					data.Id = basetypes.NewStringValue(attributeValue.(string))
					setNetflowExporterPolParentDn(ctx, attributeValue.(string), data)
				}
				if attributeName == "annotation" {
					data.Annotation = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "descr" {
					data.Descr = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "dscp" {
					data.Dscp = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "dstAddr" {
					data.DstAddr = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "dstPort" {
					data.DstPort = customtypes.NewDstPortStringValue(attributeValue.(string))
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
				if attributeName == "sourceIpType" {
					data.SourceIpType = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "srcAddr" {
					data.SrcAddr = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ver" {
					data.Ver = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			TagAnnotationNetflowExporterPolList := make([]TagAnnotationNetflowExporterPolResourceModel, 0)
			TagTagNetflowExporterPolList := make([]TagTagNetflowExporterPolResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "tagAnnotation" {
							TagAnnotationNetflowExporterPol := TagAnnotationNetflowExporterPolResourceModel{}
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationNetflowExporterPol.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationNetflowExporterPol.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagAnnotationNetflowExporterPolList = append(TagAnnotationNetflowExporterPolList, TagAnnotationNetflowExporterPol)
						}
						if childClassName == "tagTag" {
							TagTagNetflowExporterPol := TagTagNetflowExporterPolResourceModel{}
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagTagNetflowExporterPol.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagTagNetflowExporterPol.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagTagNetflowExporterPolList = append(TagTagNetflowExporterPolList, TagTagNetflowExporterPol)
						}
					}
				}
			}
			tagAnnotationSet, _ := types.SetValueFrom(ctx, data.TagAnnotation.ElementType(ctx), TagAnnotationNetflowExporterPolList)
			data.TagAnnotation = tagAnnotationSet
			tagTagSet, _ := types.SetValueFrom(ctx, data.TagTag.ElementType(ctx), TagTagNetflowExporterPolList)
			data.TagTag = tagTagSet
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'netflowExporterPol'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	} else {
		data.Id = basetypes.NewStringNull()
	}
}

func getNetflowExporterPolRn(ctx context.Context, data *NetflowExporterPolResourceModel) string {
	rn := "exporterpol-{name}"
	for _, identifier := range []string{"name"} {
		fieldName := fmt.Sprintf("%s%s", strings.ToUpper(identifier[:1]), identifier[1:])
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface().(basetypes.StringValue).ValueString()
		rn = strings.ReplaceAll(rn, fmt.Sprintf("{%s}", identifier), fieldValue)
	}
	return rn
}

func setNetflowExporterPolParentDn(ctx context.Context, dn string, data *NetflowExporterPolResourceModel) {
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

func setNetflowExporterPolId(ctx context.Context, data *NetflowExporterPolResourceModel) {
	rn := getNetflowExporterPolRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
}

func getNetflowExporterPolTagAnnotationChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *NetflowExporterPolResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationNetflowExporterPolResourceModel) []map[string]interface{} {

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
func getNetflowExporterPolTagTagChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *NetflowExporterPolResourceModel, tagTagPlan, tagTagState []TagTagNetflowExporterPolResourceModel) []map[string]interface{} {

	childPayloads := []map[string]interface{}{}
	if !data.TagTag.IsUnknown() {
		tagTagIdentifiers := []TagTagIdentifier{}
		for _, tagTag := range tagTagPlan {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			if !tagTag.Key.IsUnknown() {
				childMap["attributes"]["key"] = tagTag.Key.ValueString()
			}
			if !tagTag.Value.IsUnknown() {
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

func getNetflowExporterPolCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, createType bool, data *NetflowExporterPolResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationNetflowExporterPolResourceModel, tagTagPlan, tagTagState []TagTagNetflowExporterPolResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}

	if createType && !globalAllowExistingOnCreate {
		payloadMap["attributes"].(map[string]string)["status"] = "created"
	}
	childPayloads := []map[string]interface{}{}

	TagAnnotationchildPayloads := getNetflowExporterPolTagAnnotationChildPayloads(ctx, diags, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)

	TagTagchildPayloads := getNetflowExporterPolTagTagChildPayloads(ctx, diags, data, tagTagPlan, tagTagState)
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
	if !data.Dscp.IsNull() && !data.Dscp.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["dscp"] = data.Dscp.ValueString()
	}
	if !data.DstAddr.IsNull() && !data.DstAddr.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["dstAddr"] = data.DstAddr.ValueString()
	}
	if !data.DstPort.IsNull() && !data.DstPort.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["dstPort"] = data.DstPort.ValueString()
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
	if !data.SourceIpType.IsNull() && !data.SourceIpType.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["sourceIpType"] = data.SourceIpType.ValueString()
	}
	if !data.SrcAddr.IsNull() && !data.SrcAddr.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["srcAddr"] = data.SrcAddr.ValueString()
	}
	if !data.Ver.IsNull() && !data.Ver.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["ver"] = data.Ver.ValueString()
	}

	payload, err := json.Marshal(map[string]interface{}{"netflowExporterPol": payloadMap})
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
