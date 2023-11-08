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
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MgmtSubnetResource{}
var _ resource.ResourceWithImportState = &MgmtSubnetResource{}

func NewMgmtSubnetResource() resource.Resource {
	return &MgmtSubnetResource{}
}

// MgmtSubnetResource defines the resource implementation.
type MgmtSubnetResource struct {
	client *client.Client
}

// MgmtSubnetResourceModel describes the resource data model.
type MgmtSubnetResourceModel struct {
	Id            types.String `tfsdk:"id"`
	ParentDn      types.String `tfsdk:"parent_dn"`
	Annotation    types.String `tfsdk:"annotation"`
	Descr         types.String `tfsdk:"description"`
	Ip            types.String `tfsdk:"ip"`
	Name          types.String `tfsdk:"name"`
	NameAlias     types.String `tfsdk:"name_alias"`
	TagAnnotation types.Set    `tfsdk:"annotations"`
}

// TagAnnotationMgmtSubnetResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationMgmtSubnetResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type MgmtSubnetIdentifier struct {
	Ip types.String
}

func (r *MgmtSubnetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Trace(ctx, "Start metadata of resource: aci_l3out_management_network_subnet")
	resp.TypeName = req.ProviderTypeName + "_l3out_management_network_subnet"
	tflog.Trace(ctx, "End metadata of resource: aci_l3out_management_network_subnet")
}

func (r *MgmtSubnetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Trace(ctx, "Start schema of resource: aci_l3out_management_network_subnet")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The l3out_management_network_subnet resource for the 'mgmtSubnet' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the L3out Management Network Subnet object.",
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
				},
				Default:             stringdefault.StaticString(globalAnnotation),
				MarkdownDescription: `The annotation of the L3out Management Network Subnet object.`,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `The description of the L3out Management Network Subnet object.`,
			},
			"ip": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The external subnet IP address and subnet mask. This IP address is used for creating an external management entity. The subnet mask for the IP address to be imported from the outside into the fabric. The contracts associated with its parent instance profile (l3ext:InstP) are applied to the subnet.`,
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `The name of the L3out Management Network Subnet object.`,
			},
			"name_alias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `The name alias of the L3out Management Network Subnet object.`,
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
	tflog.Trace(ctx, "End schema of resource: aci_l3out_management_network_subnet")
}

func (r *MgmtSubnetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Trace(ctx, "Start configure of resource: aci_l3out_management_network_subnet")
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		resp.Diagnostics.AddError(
			"Provider has not been configured",
			"The req.ProviderData is nil. Please report this issue to the provider developers.",
		)
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
	tflog.Trace(ctx, "End configure of resource: aci_l3out_management_network_subnet")
}

func (r *MgmtSubnetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Trace(ctx, "Start create of resource: aci_l3out_management_network_subnet")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *MgmtSubnetResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	setMgmtSubnetId(ctx, stateData)
	setMgmtSubnetAttributes(ctx, &resp.Diagnostics, r.client, stateData)

	var data *MgmtSubnetResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setMgmtSubnetId(ctx, data)

	tflog.Trace(ctx, fmt.Sprintf("Create of resource aci_l3out_management_network_subnet with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationMgmtSubnetResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	jsonPayload := getMgmtSubnetCreateJsonPayload(ctx, &resp.Diagnostics, data, tagAnnotationPlan, tagAnnotationState)

	if resp.Diagnostics.HasError() {
		return
	}

	doMgmtSubnetRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}

	setMgmtSubnetAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "End create of resource: aci_l3out_management_network_subnet")
}

func (r *MgmtSubnetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Trace(ctx, "Start read of resource: aci_l3out_management_network_subnet")
	var data *MgmtSubnetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("Read of resource aci_l3out_management_network_subnet with id '%s'", data.Id.ValueString()))

	setMgmtSubnetAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "End read of resource: aci_l3out_management_network_subnet")
}

func (r *MgmtSubnetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Trace(ctx, "Start update of resource: aci_l3out_management_network_subnet")
	var data *MgmtSubnetResourceModel
	var stateData *MgmtSubnetResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("Update of resource aci_l3out_management_network_subnet with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationMgmtSubnetResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	jsonPayload := getMgmtSubnetCreateJsonPayload(ctx, &resp.Diagnostics, data, tagAnnotationPlan, tagAnnotationState)

	if resp.Diagnostics.HasError() {
		return
	}

	doMgmtSubnetRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	setMgmtSubnetAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "End update of resource: aci_l3out_management_network_subnet")
}

func (r *MgmtSubnetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Trace(ctx, "Start delete of resource: aci_l3out_management_network_subnet")
	var data *MgmtSubnetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("Delete of resource aci_l3out_management_network_subnet with id '%s'", data.Id.ValueString()))
	jsonPayload := getMgmtSubnetDeleteJsonPayload(ctx, &resp.Diagnostics, data)
	if resp.Diagnostics.HasError() {
		return
	}
	doMgmtSubnetRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Trace(ctx, "End delete of resource: aci_l3out_management_network_subnet")
}

func (r *MgmtSubnetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func setMgmtSubnetAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *MgmtSubnetResourceModel) {
	requestData := doMgmtSubnetRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=children&rsp-subtree-class=%s", data.Id.ValueString(), "mgmtSubnet,tagAnnotation"), "GET", nil)

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("mgmtSubnet").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("mgmtSubnet").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					data.Id = basetypes.NewStringValue(attributeValue.(string))
					setMgmtSubnetParentDn(ctx, attributeValue.(string), data)
				}
				if attributeName == "annotation" {
					data.Annotation = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "descr" {
					data.Descr = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ip" {
					data.Ip = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "name" {
					data.Name = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "nameAlias" {
					data.NameAlias = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			TagAnnotationMgmtSubnetList := make([]TagAnnotationMgmtSubnetResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "tagAnnotation" {
							TagAnnotationMgmtSubnet := TagAnnotationMgmtSubnetResourceModel{}
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationMgmtSubnet.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationMgmtSubnet.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagAnnotationMgmtSubnetList = append(TagAnnotationMgmtSubnetList, TagAnnotationMgmtSubnet)
						}
					}
				}
			}
			if len(TagAnnotationMgmtSubnetList) > 0 {
				tagAnnotationSet, _ := types.SetValueFrom(ctx, data.TagAnnotation.ElementType(ctx), TagAnnotationMgmtSubnetList)
				data.TagAnnotation = tagAnnotationSet
			}
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'mgmtSubnet'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	}
}

func getMgmtSubnetRn(ctx context.Context, data *MgmtSubnetResourceModel) string {
	rn := "subnet-[{ip}]"
	for _, identifier := range []string{"ip"} {
		fieldName := fmt.Sprintf("%s%s", strings.ToUpper(identifier[:1]), identifier[1:])
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface().(basetypes.StringValue).ValueString()
		rn = strings.ReplaceAll(rn, fmt.Sprintf("{%s}", identifier), fieldValue)
	}
	return rn
}

func setMgmtSubnetParentDn(ctx context.Context, dn string, data *MgmtSubnetResourceModel) {
	rn := getMgmtSubnetRn(ctx, data)
	data.ParentDn = basetypes.NewStringValue(strings.ReplaceAll(dn, fmt.Sprintf("/%s", rn), ""))
}

func setMgmtSubnetId(ctx context.Context, data *MgmtSubnetResourceModel) {
	rn := getMgmtSubnetRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
}

func getMgmtSubnetTagAnnotationChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *MgmtSubnetResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationMgmtSubnetResourceModel) []map[string]interface{} {

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

func getMgmtSubnetCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, data *MgmtSubnetResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationMgmtSubnetResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}
	childPayloads := []map[string]interface{}{}

	TagAnnotationchildPayloads := getMgmtSubnetTagAnnotationChildPayloads(ctx, diags, data, tagAnnotationPlan, tagAnnotationState)
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
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["ip"] = data.Ip.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["name"] = data.Name.ValueString()
	}
	if !data.NameAlias.IsNull() && !data.NameAlias.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["nameAlias"] = data.NameAlias.ValueString()
	}

	payload, err := json.Marshal(map[string]interface{}{"mgmtSubnet": payloadMap})
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

func getMgmtSubnetDeleteJsonPayload(ctx context.Context, diags *diag.Diagnostics, data *MgmtSubnetResourceModel) *container.Container {

	jsonString := fmt.Sprintf(`{"mgmtSubnet":{"attributes":{"dn": "%s","status": "deleted"}}}`, data.Id.ValueString())
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

func doMgmtSubnetRequest(ctx context.Context, diags *diag.Diagnostics, client *client.Client, path, method string, payload *container.Container) *container.Container {

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
