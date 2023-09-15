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
var _ resource.Resource = &FvRsConsIfResource{}
var _ resource.ResourceWithImportState = &FvRsConsIfResource{}

func NewFvRsConsIfResource() resource.Resource {
	return &FvRsConsIfResource{}
}

// FvRsConsIfResource defines the resource implementation.
type FvRsConsIfResource struct {
	client *client.Client
}

// FvRsConsIfResourceModel describes the resource data model.
type FvRsConsIfResourceModel struct {
	Id            types.String `tfsdk:"id"`
	ParentDn      types.String `tfsdk:"parent_dn"`
	Annotation    types.String `tfsdk:"annotation"`
	Prio          types.String `tfsdk:"priority"`
	TnVzCPIfName  types.String `tfsdk:"contract_interface_name"`
	TagAnnotation types.Set    `tfsdk:"annotations"`
}

// TagAnnotationFvRsConsIfResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationFvRsConsIfResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type FvRsConsIfIdentifier struct {
	TnVzCPIfName types.String
}

func (r *FvRsConsIfResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Trace(ctx, "start schema of resource: aci_contract_interface")
	resp.TypeName = req.ProviderTypeName + "_contract_interface"
	tflog.Trace(ctx, "end schema of resource: aci_contract_interface")
}

func (r *FvRsConsIfResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The contract_interface resource for the 'fvRsConsIf' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinquised name (DN) of the Contract Interface object.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"parent_dn": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The distinquised name (DN) of the parent object.",
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
				MarkdownDescription: `The annotation of the Contract Interface object.`,
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
				MarkdownDescription: `The contract interface priority.`,
			},
			"contract_interface_name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The contract interface name.`,
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
							Optional: true,
							Computed: true,
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
}

func (r *FvRsConsIfResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Trace(ctx, "start configure of resource: aci_contract_interface")
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
	tflog.Trace(ctx, "end configure of resource: aci_contract_interface")
}

func (r *FvRsConsIfResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Trace(ctx, "start create of resource: aci_contract_interface")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *FvRsConsIfResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	setFvRsConsIfId(ctx, stateData)
	messageMap := setFvRsConsIfAttributes(ctx, r.client, stateData)
	if messageMap != nil {
		resp.Diagnostics.AddError(messageMap.(map[string]string)["message"], messageMap.(map[string]string)["messageDetail"])
	}

	var data *FvRsConsIfResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setFvRsConsIfId(ctx, data)

	tflog.Trace(ctx, fmt.Sprintf("create of resource aci_contract_interface with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationFvRsConsIfResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	jsonPayload, message, messageDetail := getFvRsConsIfCreateJsonPayload(ctx, data, tagAnnotationPlan, tagAnnotationState)

	if jsonPayload == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}

	requestData, message, messageDetail := doFvRsConsIfRequest(ctx, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if requestData == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}

	messageMap = setFvRsConsIfAttributes(ctx, r.client, data)
	if messageMap != nil {
		resp.Diagnostics.AddError(messageMap.(map[string]string)["message"], messageMap.(map[string]string)["messageDetail"])
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "end create of resource: aci_contract_interface")
}

func (r *FvRsConsIfResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Trace(ctx, "start read of resource: aci_contract_interface")
	var data *FvRsConsIfResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("read of resource aci_contract_interface with id '%s'", data.Id.ValueString()))

	messageMap := setFvRsConsIfAttributes(ctx, r.client, data)
	if messageMap != nil {
		resp.Diagnostics.AddError(messageMap.(map[string]string)["message"], messageMap.(map[string]string)["messageDetail"])
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "end read of resource: aci_contract_interface")
}

func (r *FvRsConsIfResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Trace(ctx, "start update of resource: aci_contract_interface")
	var data *FvRsConsIfResourceModel
	var stateData *FvRsConsIfResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("update of resource aci_contract_interface with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationFvRsConsIfResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	jsonPayload, message, messageDetail := getFvRsConsIfCreateJsonPayload(ctx, data, tagAnnotationPlan, tagAnnotationState)

	if jsonPayload == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}

	requestData, message, messageDetail := doFvRsConsIfRequest(ctx, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if requestData == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}

	messageMap := setFvRsConsIfAttributes(ctx, r.client, data)
	if messageMap != nil {
		resp.Diagnostics.AddError(messageMap.(map[string]string)["message"], messageMap.(map[string]string)["messageDetail"])
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "end update of resource: aci_contract_interface")
}

func (r *FvRsConsIfResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Trace(ctx, "start delete of resource: aci_contract_interface")
	var data *FvRsConsIfResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("delete of resource aci_contract_interface with id '%s'", data.Id.ValueString()))
	jsonPayload, message, messageDetail := getFvRsConsIfDeleteJsonPayload(ctx, data)
	if jsonPayload == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}
	requestData, message, messageDetail := doFvRsConsIfRequest(ctx, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if requestData == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}
	tflog.Trace(ctx, "end delete of resource: aci_contract_interface")
}

func (r *FvRsConsIfResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func setFvRsConsIfAttributes(ctx context.Context, client *client.Client, data *FvRsConsIfResourceModel) interface{} {
	requestData, message, messageDetail := doFvRsConsIfRequest(ctx, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=children&rsp-subtree-class=%s", data.Id.ValueString(), "fvRsConsIf,tagAnnotation"), "GET", nil)

	if requestData == nil {
		return map[string]string{"message": message, "messageDetail": messageDetail}
	}
	if requestData.Search("imdata").Search("fvRsConsIf").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("fvRsConsIf").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					data.Id = basetypes.NewStringValue(attributeValue.(string))
					setFvRsConsIfParentDn(ctx, attributeValue.(string), data)
				}
				if attributeName == "annotation" {
					data.Annotation = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "prio" {
					data.Prio = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "tnVzCPIfName" {
					data.TnVzCPIfName = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			TagAnnotationFvRsConsIfList := make([]TagAnnotationFvRsConsIfResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "tagAnnotation" {
							TagAnnotationFvRsConsIf := TagAnnotationFvRsConsIfResourceModel{}
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationFvRsConsIf.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationFvRsConsIf.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagAnnotationFvRsConsIfList = append(TagAnnotationFvRsConsIfList, TagAnnotationFvRsConsIf)
						}
					}
				}
			}
			if len(TagAnnotationFvRsConsIfList) > 0 {
				tagAnnotationSet, _ := types.SetValueFrom(ctx, data.TagAnnotation.ElementType(ctx), TagAnnotationFvRsConsIfList)
				data.TagAnnotation = tagAnnotationSet
			}
		} else {
			return map[string]string{
				"message":       "too many results in response",
				"messageDetail": fmt.Sprintf("%v matches returned for class 'fvRsConsIf'. Please report this issue to the provider developers.", len(classReadInfo)),
			}
		}
	}
	return nil
}

func getFvRsConsIfRn(ctx context.Context, data *FvRsConsIfResourceModel) string {
	rn := "rsconsIf-{tnVzCPIfName}"
	for _, identifier := range []string{"tnVzCPIfName"} {
		fieldName := fmt.Sprintf("%s%s", strings.ToUpper(identifier[:1]), identifier[1:])
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface().(basetypes.StringValue).ValueString()
		rn = strings.ReplaceAll(rn, fmt.Sprintf("{%s}", identifier), fieldValue)
	}
	return rn
}

func setFvRsConsIfParentDn(ctx context.Context, dn string, data *FvRsConsIfResourceModel) {
	rn := getFvRsConsIfRn(ctx, data)
	data.ParentDn = basetypes.NewStringValue(strings.ReplaceAll(dn, fmt.Sprintf("/%s", rn), ""))
}

func setFvRsConsIfId(ctx context.Context, data *FvRsConsIfResourceModel) {
	rn := getFvRsConsIfRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
}

func getFvRsConsIfTagAnnotationChildPayloads(ctx context.Context, data *FvRsConsIfResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationFvRsConsIfResourceModel) ([]map[string]interface{}, string, string) {

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

	return childPayloads, "", ""
}

func getFvRsConsIfCreateJsonPayload(ctx context.Context, data *FvRsConsIfResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationFvRsConsIfResourceModel) (*container.Container, string, string) {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}
	childPayloads := []map[string]interface{}{}

	TagAnnotationchildPayloads, errMessage, errMessageDetail := getFvRsConsIfTagAnnotationChildPayloads(ctx, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil, errMessage, errMessageDetail
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)

	payloadMap["children"] = childPayloads
	if !data.Annotation.IsNull() && !data.Annotation.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["annotation"] = data.Annotation.ValueString()
	}
	if !data.Prio.IsNull() && !data.Prio.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["prio"] = data.Prio.ValueString()
	}
	if !data.TnVzCPIfName.IsNull() && !data.TnVzCPIfName.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["tnVzCPIfName"] = data.TnVzCPIfName.ValueString()
	}

	payload, err := json.Marshal(map[string]interface{}{"fvRsConsIf": payloadMap})
	if err != nil {
		errMessage := "marshalling of json payload failed"
		errMessageDetail := fmt.Sprintf("err: %s. Please report this issue to the provider developers.", err)
		return nil, errMessage, errMessageDetail
	}

	jsonPayload, err := container.ParseJSON(payload)

	if err != nil {
		errMessage := "construction of json payload failed"
		errMessageDetail := fmt.Sprintf("err: %s. Please report this issue to the provider developers.", err)
		return nil, errMessage, errMessageDetail
	}
	return jsonPayload, "", ""
}

func getFvRsConsIfDeleteJsonPayload(ctx context.Context, data *FvRsConsIfResourceModel) (*container.Container, string, string) {

	jsonString := fmt.Sprintf(`{"fvRsConsIf":{"attributes":{"dn": "%s","status": "deleted"}}}`, data.Id.ValueString())
	jsonPayload, err := container.ParseJSON([]byte(jsonString))
	if err != nil {
		errMessage := "construction of json payload failed"
		errMessageDetail := fmt.Sprintf("err: %s. Please report this issue to the provider developers.", err)
		return nil, errMessage, errMessageDetail
	}
	return jsonPayload, "", ""
}

func doFvRsConsIfRequest(ctx context.Context, client *client.Client, path, method string, payload *container.Container) (*container.Container, string, string) {

	restRequest, err := client.MakeRestRequest(method, path, payload, true)
	if err != nil {
		message := fmt.Sprintf("creation of %s rest request failed", strings.ToLower(method))
		messageDetail := fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err)
		return nil, message, messageDetail
	}

	cont, restResponse, err := client.Do(restRequest)

	if restResponse != nil && restResponse.StatusCode != 200 {
		message := fmt.Sprintf("%s rest request failed", strings.ToLower(method))
		messageDetail := fmt.Sprintf("Response: %s, err: %s. Please report this issue to the provider developers.", cont.Data().(map[string]interface{})["imdata"], err)
		return nil, message, messageDetail
	} else if err != nil {
		message := fmt.Sprintf("%s rest request failed", strings.ToLower(method))
		messageDetail := fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err)
		return nil, message, messageDetail
	}

	return cont, "", ""
}
