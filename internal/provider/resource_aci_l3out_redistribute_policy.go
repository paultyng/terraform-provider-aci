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
var _ resource.Resource = &L3extRsRedistributePolResource{}
var _ resource.ResourceWithImportState = &L3extRsRedistributePolResource{}

func NewL3extRsRedistributePolResource() resource.Resource {
	return &L3extRsRedistributePolResource{}
}

// L3extRsRedistributePolResource defines the resource implementation.
type L3extRsRedistributePolResource struct {
	client *client.Client
}

// L3extRsRedistributePolResourceModel describes the resource data model.
type L3extRsRedistributePolResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	ParentDn            types.String `tfsdk:"parent_dn"`
	Annotation          types.String `tfsdk:"annotation"`
	Src                 types.String `tfsdk:"src"`
	TnRtctrlProfileName types.String `tfsdk:"route_control_profile_name"`
	TagAnnotation       types.Set    `tfsdk:"annotations"`
}

// TagAnnotationL3extRsRedistributePolResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationL3extRsRedistributePolResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type L3extRsRedistributePolIdentifier struct {
	Src                 types.String
	TnRtctrlProfileName types.String
}

func (r *L3extRsRedistributePolResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Trace(ctx, "start schema of resource: aci_l3out_redistribute_policy")
	resp.TypeName = req.ProviderTypeName + "_l3out_redistribute_policy"
	tflog.Trace(ctx, "end schema of resource: aci_l3out_redistribute_policy")
}

func (r *L3extRsRedistributePolResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The l3out_redistribute_policy resource for the 'l3extRsRedistributePol' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinquised name (DN) of the L3out Redistribute Policy object.",
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
				MarkdownDescription: `The annotation of the L3out Redistribute Policy object.`,
			},
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("attached-host", "direct", "static"),
				},
				MarkdownDescription: `The source IP address.`,
			},
			"route_control_profile_name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The name of the route profile associated with this object.`,
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
}

func (r *L3extRsRedistributePolResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Trace(ctx, "start configure of resource: aci_l3out_redistribute_policy")
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
	tflog.Trace(ctx, "end configure of resource: aci_l3out_redistribute_policy")
}

func (r *L3extRsRedistributePolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Trace(ctx, "start create of resource: aci_l3out_redistribute_policy")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *L3extRsRedistributePolResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	setL3extRsRedistributePolId(ctx, stateData)
	messageMap := setL3extRsRedistributePolAttributes(ctx, r.client, stateData)
	if messageMap != nil {
		resp.Diagnostics.AddError(messageMap.(map[string]string)["message"], messageMap.(map[string]string)["messageDetail"])
	}

	var data *L3extRsRedistributePolResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setL3extRsRedistributePolId(ctx, data)

	tflog.Trace(ctx, fmt.Sprintf("create of resource aci_l3out_redistribute_policy with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationL3extRsRedistributePolResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	jsonPayload, message, messageDetail := getL3extRsRedistributePolCreateJsonPayload(ctx, data, tagAnnotationPlan, tagAnnotationState)

	if jsonPayload == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}

	requestData, message, messageDetail := doL3extRsRedistributePolRequest(ctx, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if requestData == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}

	messageMap = setL3extRsRedistributePolAttributes(ctx, r.client, data)
	if messageMap != nil {
		resp.Diagnostics.AddError(messageMap.(map[string]string)["message"], messageMap.(map[string]string)["messageDetail"])
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "end create of resource: aci_l3out_redistribute_policy")
}

func (r *L3extRsRedistributePolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Trace(ctx, "start read of resource: aci_l3out_redistribute_policy")
	var data *L3extRsRedistributePolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("read of resource aci_l3out_redistribute_policy with id '%s'", data.Id.ValueString()))

	messageMap := setL3extRsRedistributePolAttributes(ctx, r.client, data)
	if messageMap != nil {
		resp.Diagnostics.AddError(messageMap.(map[string]string)["message"], messageMap.(map[string]string)["messageDetail"])
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "end read of resource: aci_l3out_redistribute_policy")
}

func (r *L3extRsRedistributePolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Trace(ctx, "start update of resource: aci_l3out_redistribute_policy")
	var data *L3extRsRedistributePolResourceModel
	var stateData *L3extRsRedistributePolResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("update of resource aci_l3out_redistribute_policy with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationL3extRsRedistributePolResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	jsonPayload, message, messageDetail := getL3extRsRedistributePolCreateJsonPayload(ctx, data, tagAnnotationPlan, tagAnnotationState)

	if jsonPayload == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}

	requestData, message, messageDetail := doL3extRsRedistributePolRequest(ctx, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if requestData == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}

	messageMap := setL3extRsRedistributePolAttributes(ctx, r.client, data)
	if messageMap != nil {
		resp.Diagnostics.AddError(messageMap.(map[string]string)["message"], messageMap.(map[string]string)["messageDetail"])
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "end update of resource: aci_l3out_redistribute_policy")
}

func (r *L3extRsRedistributePolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Trace(ctx, "start delete of resource: aci_l3out_redistribute_policy")
	var data *L3extRsRedistributePolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("delete of resource aci_l3out_redistribute_policy with id '%s'", data.Id.ValueString()))
	jsonPayload, message, messageDetail := getL3extRsRedistributePolDeleteJsonPayload(ctx, data)
	if jsonPayload == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}
	requestData, message, messageDetail := doL3extRsRedistributePolRequest(ctx, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if requestData == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}
	tflog.Trace(ctx, "end delete of resource: aci_l3out_redistribute_policy")
}

func (r *L3extRsRedistributePolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func setL3extRsRedistributePolAttributes(ctx context.Context, client *client.Client, data *L3extRsRedistributePolResourceModel) interface{} {
	requestData, message, messageDetail := doL3extRsRedistributePolRequest(ctx, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=children&rsp-subtree-class=%s", data.Id.ValueString(), "l3extRsRedistributePol,tagAnnotation"), "GET", nil)

	if requestData == nil {
		return map[string]string{"message": message, "messageDetail": messageDetail}
	}
	if requestData.Search("imdata").Search("l3extRsRedistributePol").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("l3extRsRedistributePol").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					data.Id = basetypes.NewStringValue(attributeValue.(string))
					setL3extRsRedistributePolParentDn(ctx, attributeValue.(string), data)
				}
				if attributeName == "annotation" {
					data.Annotation = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "src" {
					data.Src = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "tnRtctrlProfileName" {
					data.TnRtctrlProfileName = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			TagAnnotationL3extRsRedistributePolList := make([]TagAnnotationL3extRsRedistributePolResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "tagAnnotation" {
							TagAnnotationL3extRsRedistributePol := TagAnnotationL3extRsRedistributePolResourceModel{}
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationL3extRsRedistributePol.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationL3extRsRedistributePol.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagAnnotationL3extRsRedistributePolList = append(TagAnnotationL3extRsRedistributePolList, TagAnnotationL3extRsRedistributePol)
						}
					}
				}
			}
			if len(TagAnnotationL3extRsRedistributePolList) > 0 {
				tagAnnotationSet, _ := types.SetValueFrom(ctx, data.TagAnnotation.ElementType(ctx), TagAnnotationL3extRsRedistributePolList)
				data.TagAnnotation = tagAnnotationSet
			}
		} else {
			return map[string]string{
				"message":       "too many results in response",
				"messageDetail": fmt.Sprintf("%v matches returned for class 'l3extRsRedistributePol'. Please report this issue to the provider developers.", len(classReadInfo)),
			}
		}
	}
	return nil
}

func getL3extRsRedistributePolRn(ctx context.Context, data *L3extRsRedistributePolResourceModel) string {
	rn := "rsredistributePol-[{tnRtctrlProfileName}]-{src}"
	for _, identifier := range []string{"tnRtctrlProfileName", "src"} {
		fieldName := fmt.Sprintf("%s%s", strings.ToUpper(identifier[:1]), identifier[1:])
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface().(basetypes.StringValue).ValueString()
		rn = strings.ReplaceAll(rn, fmt.Sprintf("{%s}", identifier), fieldValue)
	}
	return rn
}

func setL3extRsRedistributePolParentDn(ctx context.Context, dn string, data *L3extRsRedistributePolResourceModel) {
	rn := getL3extRsRedistributePolRn(ctx, data)
	data.ParentDn = basetypes.NewStringValue(strings.ReplaceAll(dn, fmt.Sprintf("/%s", rn), ""))
}

func setL3extRsRedistributePolId(ctx context.Context, data *L3extRsRedistributePolResourceModel) {
	rn := getL3extRsRedistributePolRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
}

func getL3extRsRedistributePolTagAnnotationChildPayloads(ctx context.Context, data *L3extRsRedistributePolResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationL3extRsRedistributePolResourceModel) ([]map[string]interface{}, string, string) {

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

func getL3extRsRedistributePolCreateJsonPayload(ctx context.Context, data *L3extRsRedistributePolResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationL3extRsRedistributePolResourceModel) (*container.Container, string, string) {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}
	childPayloads := []map[string]interface{}{}

	TagAnnotationchildPayloads, errMessage, errMessageDetail := getL3extRsRedistributePolTagAnnotationChildPayloads(ctx, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil, errMessage, errMessageDetail
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)

	payloadMap["children"] = childPayloads
	if !data.Annotation.IsNull() && !data.Annotation.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["annotation"] = data.Annotation.ValueString()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["src"] = data.Src.ValueString()
	}
	if !data.TnRtctrlProfileName.IsNull() && !data.TnRtctrlProfileName.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["tnRtctrlProfileName"] = data.TnRtctrlProfileName.ValueString()
	}

	payload, err := json.Marshal(map[string]interface{}{"l3extRsRedistributePol": payloadMap})
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

func getL3extRsRedistributePolDeleteJsonPayload(ctx context.Context, data *L3extRsRedistributePolResourceModel) (*container.Container, string, string) {

	jsonString := fmt.Sprintf(`{"l3extRsRedistributePol":{"attributes":{"dn": "%s","status": "deleted"}}}`, data.Id.ValueString())
	jsonPayload, err := container.ParseJSON([]byte(jsonString))
	if err != nil {
		errMessage := "construction of json payload failed"
		errMessageDetail := fmt.Sprintf("err: %s. Please report this issue to the provider developers.", err)
		return nil, errMessage, errMessageDetail
	}
	return jsonPayload, "", ""
}

func doL3extRsRedistributePolRequest(ctx context.Context, client *client.Client, path, method string, payload *container.Container) (*container.Container, string, string) {

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
