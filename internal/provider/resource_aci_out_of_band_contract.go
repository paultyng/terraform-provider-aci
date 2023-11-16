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
var _ resource.Resource = &VzOOBBrCPResource{}
var _ resource.ResourceWithImportState = &VzOOBBrCPResource{}

func NewVzOOBBrCPResource() resource.Resource {
	return &VzOOBBrCPResource{}
}

// VzOOBBrCPResource defines the resource implementation.
type VzOOBBrCPResource struct {
	client *client.Client
}

// VzOOBBrCPResourceModel describes the resource data model.
type VzOOBBrCPResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Annotation    types.String `tfsdk:"annotation"`
	Descr         types.String `tfsdk:"description"`
	Intent        types.String `tfsdk:"intent"`
	Name          types.String `tfsdk:"name"`
	NameAlias     types.String `tfsdk:"name_alias"`
	OwnerKey      types.String `tfsdk:"owner_key"`
	OwnerTag      types.String `tfsdk:"owner_tag"`
	Prio          types.String `tfsdk:"priority"`
	Scope         types.String `tfsdk:"scope"`
	TargetDscp    types.String `tfsdk:"target_dscp"`
	TagAnnotation types.Set    `tfsdk:"annotations"`
}

// TagAnnotationVzOOBBrCPResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationVzOOBBrCPResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type VzOOBBrCPIdentifier struct {
	Name types.String
}

func (r *VzOOBBrCPResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: aci_out_of_band_contract")
	resp.TypeName = req.ProviderTypeName + "_out_of_band_contract"
	tflog.Debug(ctx, "End metadata of resource: aci_out_of_band_contract")
}

func (r *VzOOBBrCPResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: aci_out_of_band_contract")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The out_of_band_contract resource for the 'vzOOBBrCP' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the Out Of Band Contract object.",
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
				MarkdownDescription: `The annotation of the Out Of Band Contract object.`,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `The description of the Out Of Band Contract object.`,
			},
			"intent": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("estimate_add", "estimate_delete", "install"),
				},
				MarkdownDescription: `Install Rules or Estimate Nummber of Rules.`,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The name of the Out Of Band Contract object.`,
			},
			"name_alias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `The name alias of the Out Of Band Contract object.`,
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
			"priority": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("level1", "level2", "level3", "level4", "level5", "level6", "unspecified"),
				},
				MarkdownDescription: `null.`,
			},
			"scope": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("application-profile", "context", "global", "tenant"),
				},
				MarkdownDescription: `Represents the scope of this contract. If the scope is set as application-profile, the epg can only communicate with epgs in the same application-profile.`,
			},
			"target_dscp": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("AF11", "AF12", "AF13", "AF21", "AF22", "AF23", "AF31", "AF32", "AF33", "AF41", "AF42", "AF43", "CS0", "CS1", "CS2", "CS3", "CS4", "CS5", "CS6", "CS7", "EF", "VA", "unspecified"),
				},
				MarkdownDescription: `contract level dscp value.`,
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
							MarkdownDescription: `The key or password used to uniquely identify this configuration object.`,
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
	tflog.Debug(ctx, "End schema of resource: aci_out_of_band_contract")
}

func (r *VzOOBBrCPResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: aci_out_of_band_contract")
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
	tflog.Debug(ctx, "End configure of resource: aci_out_of_band_contract")
}

func (r *VzOOBBrCPResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: aci_out_of_band_contract")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *VzOOBBrCPResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	setVzOOBBrCPId(ctx, stateData)
	setVzOOBBrCPAttributes(ctx, &resp.Diagnostics, r.client, stateData)

	var data *VzOOBBrCPResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setVzOOBBrCPId(ctx, data)

	tflog.Trace(ctx, fmt.Sprintf("Create of resource aci_out_of_band_contract with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationVzOOBBrCPResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	jsonPayload := getVzOOBBrCPCreateJsonPayload(ctx, &resp.Diagnostics, data, tagAnnotationPlan, tagAnnotationState)

	if resp.Diagnostics.HasError() {
		return
	}

	doVzOOBBrCPRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}

	setVzOOBBrCPAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, "End create of resource: aci_out_of_band_contract")
}

func (r *VzOOBBrCPResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: aci_out_of_band_contract")
	var data *VzOOBBrCPResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("Read of resource aci_out_of_band_contract with id '%s'", data.Id.ValueString()))

	setVzOOBBrCPAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, "End read of resource: aci_out_of_band_contract")
}

func (r *VzOOBBrCPResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: aci_out_of_band_contract")
	var data *VzOOBBrCPResourceModel
	var stateData *VzOOBBrCPResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("Update of resource aci_out_of_band_contract with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationVzOOBBrCPResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	jsonPayload := getVzOOBBrCPCreateJsonPayload(ctx, &resp.Diagnostics, data, tagAnnotationPlan, tagAnnotationState)

	if resp.Diagnostics.HasError() {
		return
	}

	doVzOOBBrCPRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	setVzOOBBrCPAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, "End update of resource: aci_out_of_band_contract")
}

func (r *VzOOBBrCPResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: aci_out_of_band_contract")
	var data *VzOOBBrCPResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("Delete of resource aci_out_of_band_contract with id '%s'", data.Id.ValueString()))
	jsonPayload := getVzOOBBrCPDeleteJsonPayload(ctx, &resp.Diagnostics, data)
	if resp.Diagnostics.HasError() {
		return
	}
	doVzOOBBrCPRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "End delete of resource: aci_out_of_band_contract")
}

func (r *VzOOBBrCPResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func setVzOOBBrCPAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *VzOOBBrCPResourceModel) {
	requestData := doVzOOBBrCPRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=children&rsp-subtree-class=%s", data.Id.ValueString(), "vzOOBBrCP,tagAnnotation"), "GET", nil)

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("vzOOBBrCP").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("vzOOBBrCP").Data().([]interface{})
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
				if attributeName == "intent" {
					data.Intent = basetypes.NewStringValue(attributeValue.(string))
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
				if attributeName == "prio" {
					data.Prio = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "scope" {
					data.Scope = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "targetDscp" {
					data.TargetDscp = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			TagAnnotationVzOOBBrCPList := make([]TagAnnotationVzOOBBrCPResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "tagAnnotation" {
							TagAnnotationVzOOBBrCP := TagAnnotationVzOOBBrCPResourceModel{}
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationVzOOBBrCP.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationVzOOBBrCP.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagAnnotationVzOOBBrCPList = append(TagAnnotationVzOOBBrCPList, TagAnnotationVzOOBBrCP)
						}
					}
				}
			}
			if len(TagAnnotationVzOOBBrCPList) > 0 {
				tagAnnotationSet, _ := types.SetValueFrom(ctx, data.TagAnnotation.ElementType(ctx), TagAnnotationVzOOBBrCPList)
				data.TagAnnotation = tagAnnotationSet
			}
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'vzOOBBrCP'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	}
}

func getVzOOBBrCPRn(ctx context.Context, data *VzOOBBrCPResourceModel) string {
	rn := "tn-mgmt//oobbrc-{name}"
	for _, identifier := range []string{"name"} {
		fieldName := fmt.Sprintf("%s%s", strings.ToUpper(identifier[:1]), identifier[1:])
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface().(basetypes.StringValue).ValueString()
		rn = strings.ReplaceAll(rn, fmt.Sprintf("{%s}", identifier), fieldValue)
	}
	return rn
}

func setVzOOBBrCPId(ctx context.Context, data *VzOOBBrCPResourceModel) {
	rn := getVzOOBBrCPRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", strings.Split([]string{"uni/tn-mgmt/oobbrc-{name}"}[0], "/")[0], rn))
}

func getVzOOBBrCPTagAnnotationChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *VzOOBBrCPResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationVzOOBBrCPResourceModel) []map[string]interface{} {

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

func getVzOOBBrCPCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, data *VzOOBBrCPResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationVzOOBBrCPResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}
	childPayloads := []map[string]interface{}{}

	TagAnnotationchildPayloads := getVzOOBBrCPTagAnnotationChildPayloads(ctx, diags, data, tagAnnotationPlan, tagAnnotationState)
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
	if !data.Intent.IsNull() && !data.Intent.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["intent"] = data.Intent.ValueString()
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
	if !data.Prio.IsNull() && !data.Prio.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["prio"] = data.Prio.ValueString()
	}
	if !data.Scope.IsNull() && !data.Scope.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["scope"] = data.Scope.ValueString()
	}
	if !data.TargetDscp.IsNull() && !data.TargetDscp.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["targetDscp"] = data.TargetDscp.ValueString()
	}

	payload, err := json.Marshal(map[string]interface{}{"vzOOBBrCP": payloadMap})
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

func getVzOOBBrCPDeleteJsonPayload(ctx context.Context, diags *diag.Diagnostics, data *VzOOBBrCPResourceModel) *container.Container {

	jsonString := fmt.Sprintf(`{"vzOOBBrCP":{"attributes":{"dn": "%s","status": "deleted"}}}`, data.Id.ValueString())
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

func doVzOOBBrCPRequest(ctx context.Context, diags *diag.Diagnostics, client *client.Client, path, method string, payload *container.Container) *container.Container {

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
