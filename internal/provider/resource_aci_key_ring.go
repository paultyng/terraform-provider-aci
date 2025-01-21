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
var _ resource.Resource = &PkiKeyRingResource{}
var _ resource.ResourceWithImportState = &PkiKeyRingResource{}

func NewPkiKeyRingResource() resource.Resource {
	return &PkiKeyRingResource{}
}

// PkiKeyRingResource defines the resource implementation.
type PkiKeyRingResource struct {
	client *client.Client
}

// PkiKeyRingResourceModel describes the resource data model.
type PkiKeyRingResourceModel struct {
	Id            types.String `tfsdk:"id"`
	ParentDn      types.String `tfsdk:"parent_dn"`
	AdminState    types.String `tfsdk:"admin_state"`
	Annotation    types.String `tfsdk:"annotation"`
	Cert          types.String `tfsdk:"certificate"`
	Descr         types.String `tfsdk:"description"`
	EccCurve      types.String `tfsdk:"elliptic_curve"`
	Key           types.String `tfsdk:"key"`
	KeyType       types.String `tfsdk:"key_type"`
	Modulus       types.String `tfsdk:"modulus"`
	Name          types.String `tfsdk:"name"`
	NameAlias     types.String `tfsdk:"name_alias"`
	OwnerKey      types.String `tfsdk:"owner_key"`
	OwnerTag      types.String `tfsdk:"owner_tag"`
	Regen         types.String `tfsdk:"regenerate"`
	Tp            types.String `tfsdk:"certificate_authority"`
	TagAnnotation types.Set    `tfsdk:"annotations"`
	TagTag        types.Set    `tfsdk:"tags"`
}

func getEmptyPkiKeyRingResourceModel() *PkiKeyRingResourceModel {
	return &PkiKeyRingResourceModel{
		Id:         basetypes.NewStringNull(),
		ParentDn:   basetypes.NewStringNull(),
		AdminState: basetypes.NewStringNull(),
		Annotation: basetypes.NewStringNull(),
		Cert:       basetypes.NewStringNull(),
		Descr:      basetypes.NewStringNull(),
		EccCurve:   basetypes.NewStringNull(),
		Key:        basetypes.NewStringNull(),
		KeyType:    basetypes.NewStringNull(),
		Modulus:    basetypes.NewStringNull(),
		Name:       basetypes.NewStringNull(),
		NameAlias:  basetypes.NewStringNull(),
		OwnerKey:   basetypes.NewStringNull(),
		OwnerTag:   basetypes.NewStringNull(),
		Regen:      basetypes.NewStringNull(),
		Tp:         basetypes.NewStringNull(),
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

// TagAnnotationPkiKeyRingResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationPkiKeyRingResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func getEmptyTagAnnotationPkiKeyRingResourceModel() TagAnnotationPkiKeyRingResourceModel {
	return TagAnnotationPkiKeyRingResourceModel{
		Key:   basetypes.NewStringNull(),
		Value: basetypes.NewStringNull(),
	}
}

var TagAnnotationPkiKeyRingType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"key":   types.StringType,
		"value": types.StringType,
	},
}

// TagTagPkiKeyRingResourceModel describes the resource data model for the children without relation ships.
type TagTagPkiKeyRingResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func getEmptyTagTagPkiKeyRingResourceModel() TagTagPkiKeyRingResourceModel {
	return TagTagPkiKeyRingResourceModel{
		Key:   basetypes.NewStringNull(),
		Value: basetypes.NewStringNull(),
	}
}

var TagTagPkiKeyRingType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"key":   types.StringType,
		"value": types.StringType,
	},
}

type PkiKeyRingIdentifier struct {
	Name types.String
}

func (r *PkiKeyRingResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if !req.Plan.Raw.IsNull() {
		var planData, stateData *PkiKeyRingResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

		if resp.Diagnostics.HasError() {
			return
		}

		if (planData.Id.IsUnknown() || planData.Id.IsNull()) && !planData.ParentDn.IsUnknown() && !planData.Name.IsUnknown() {
			setPkiKeyRingId(ctx, planData)
		}

		if stateData == nil && !globalAllowExistingOnCreate && !planData.Id.IsUnknown() && !planData.Id.IsNull() {
			CheckDn(ctx, &resp.Diagnostics, r.client, "pkiKeyRing", planData.Id.ValueString())
			if resp.Diagnostics.HasError() {
				return
			}
		}

		resp.Diagnostics.Append(resp.Plan.Set(ctx, &planData)...)
	}
}

func (r *PkiKeyRingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: aci_key_ring")
	resp.TypeName = req.ProviderTypeName + "_key_ring"
	tflog.Debug(ctx, "End metadata of resource: aci_key_ring")
}

func (r *PkiKeyRingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: aci_key_ring")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The key_ring resource for the 'pkiKeyRing' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the Key Ring object.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"parent_dn": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("uni/userext/pkiext"),
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
					stringvalidator.OneOf("completed", "created", "reqCreated", "started", "tpSet"),
				},
				MarkdownDescription: `The current administrative state of the certificate request process.`,
			},
			"annotation": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Default:             stringdefault.StaticString(globalAnnotation),
				MarkdownDescription: `The annotation of the Key Ring object.`,
			},
			"certificate": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `A certificate contains a device's public key along with signed information verifying the identity of the device.`,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The description of the Key Ring object.`,
			},
			"elliptic_curve": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("none", "prime256v1", "secp384r1", "secp521r1"),
				},
				MarkdownDescription: `The elliptic curve used by the provided key.`,
			},
			"key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The private key of the certificate. This sensitive value is excluded from the resource's lifecycle configuration and is not tracked by Terraform.`,
			},
			"key_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("ECC", "RSA", "indeterminate"),
				},
				MarkdownDescription: `The type used by the provided key.`,
			},
			"modulus": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("mod1024", "mod1536", "mod2048", "mod3072", "mod4096", "mod512", "none"),
				},
				MarkdownDescription: `The length of the encryption keys. A longer key length increases the difficulty of breaking the key.`,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The name of the Key Ring object.`,
			},
			"name_alias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The name alias of the Key Ring object.`,
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
			"regenerate": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("no", "yes"),
				},
				MarkdownDescription: `Forces regeneration of the keypair. Each PKI device holds a pair of asymmetric Rivest-Shamir-Adleman (RSA) or Elliptic Curve Cryptography (ECC) encryption keys, one kept private and one made public, stored in an internal key ring.`,
			},
			"certificate_authority": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The certificate of the Certificate Authority (CA) that issued the certificate provided in the 'certificate' attribute. The CA can be a root CA, an intermediate CA, or a trust anchor in a chain of trust leading to a root CA.`,
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
	tflog.Debug(ctx, "End schema of resource: aci_key_ring")
}

func (r *PkiKeyRingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: aci_key_ring")
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
	tflog.Debug(ctx, "End configure of resource: aci_key_ring")
}

func (r *PkiKeyRingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: aci_key_ring")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *PkiKeyRingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	if stateData.Id.IsUnknown() || stateData.Id.IsNull() {
		setPkiKeyRingId(ctx, stateData)
	}
	getAndSetPkiKeyRingAttributes(ctx, &resp.Diagnostics, r.client, stateData)
	if !globalAllowExistingOnCreate && !stateData.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Object Already Exists",
			fmt.Sprintf("The pkiKeyRing object with DN '%s' already exists.", stateData.Id.ValueString()),
		)
		return
	}

	var data *PkiKeyRingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsUnknown() || data.Id.IsNull() {
		setPkiKeyRingId(ctx, data)
	}

	tflog.Debug(ctx, fmt.Sprintf("Create of resource aci_key_ring with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationPkiKeyRingResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagPkiKeyRingResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getPkiKeyRingCreateJsonPayload(ctx, &resp.Diagnostics, true, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	wrapperClassMap := map[string]string{"uni/userext/pkiext": "", "certstore": "cloudCertStore"}
	for rnPrepend, wrapperClass := range wrapperClassMap {
		if strings.Contains(data.Id.ValueString(), rnPrepend) && wrapperClass != "" {
			DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s%s.json", strings.Split(data.Id.ValueString(), rnPrepend)[0], rnPrepend), "POST", jsonPayload)
			break
		} else if strings.Contains(data.Id.ValueString(), rnPrepend) && wrapperClass == "" {
			DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetPkiKeyRingAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End create of resource aci_key_ring with id '%s'", data.Id.ValueString()))
}

func (r *PkiKeyRingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: aci_key_ring")
	var data *PkiKeyRingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource aci_key_ring with id '%s'", data.Id.ValueString()))

	getAndSetPkiKeyRingAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *PkiKeyRingResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, fmt.Sprintf("End read of resource aci_key_ring with id '%s'", data.Id.ValueString()))
}

func (r *PkiKeyRingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: aci_key_ring")
	var data *PkiKeyRingResourceModel
	var stateData *PkiKeyRingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource aci_key_ring with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationPkiKeyRingResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagPkiKeyRingResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getPkiKeyRingCreateJsonPayload(ctx, &resp.Diagnostics, false, data, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	wrapperClassMap := map[string]string{
		"uni/userext/pkiext": "",
		"certstore":          "cloudCertStore",
	}
	for rnPrepend, wrapperClass := range wrapperClassMap {
		if strings.Contains(data.Id.ValueString(), rnPrepend) && wrapperClass != "" {
			DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s%s.json", strings.Split(data.Id.ValueString(), rnPrepend)[0], rnPrepend), "POST", jsonPayload)
			break
		} else if strings.Contains(data.Id.ValueString(), rnPrepend) && wrapperClass == "" {
			DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetPkiKeyRingAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End update of resource aci_key_ring with id '%s'", data.Id.ValueString()))
}

func (r *PkiKeyRingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: aci_key_ring")
	var data *PkiKeyRingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource aci_key_ring with id '%s'", data.Id.ValueString()))
	jsonPayload := GetDeleteJsonPayload(ctx, &resp.Diagnostics, "pkiKeyRing", data.Id.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}
	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("End delete of resource aci_key_ring with id '%s'", data.Id.ValueString()))
}

func (r *PkiKeyRingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: aci_key_ring")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *PkiKeyRingResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)
	tflog.Debug(ctx, fmt.Sprintf("Import state of resource aci_key_ring with id '%s'", stateData.Id.ValueString()))

	tflog.Debug(ctx, "End import of state resource: aci_key_ring")
}

func getAndSetPkiKeyRingAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *PkiKeyRingResourceModel) {
	requestData := DoRestRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=full&rsp-subtree-class=%s", data.Id.ValueString(), "pkiKeyRing,tagAnnotation,tagTag"), "GET", nil)

	readData := getEmptyPkiKeyRingResourceModel()

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("pkiKeyRing").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("pkiKeyRing").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					readData.Id = basetypes.NewStringValue(attributeValue.(string))
					setPkiKeyRingParentDn(ctx, attributeValue.(string), readData)
				}
				if attributeName == "adminState" {
					readData.AdminState = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "annotation" {
					readData.Annotation = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "cert" {
					readData.Cert = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "descr" {
					readData.Descr = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "eccCurve" && attributeValue.(string) == "" {
					readData.EccCurve = basetypes.NewStringValue("none")
				} else if attributeName == "eccCurve" {
					readData.EccCurve = basetypes.NewStringValue(attributeValue.(string))
				}
				// Sensitive attributes are not returned by the APIC, so they are explicitly set to their current state values.
				readData.Key = data.Key
				if attributeName == "keyType" {
					readData.KeyType = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "modulus" && attributeValue.(string) == "" {
					readData.Modulus = basetypes.NewStringValue("none")
				} else if attributeName == "modulus" {
					readData.Modulus = basetypes.NewStringValue(attributeValue.(string))
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
				if attributeName == "regen" {
					readData.Regen = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "tp" {
					readData.Tp = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			TagAnnotationPkiKeyRingList := make([]TagAnnotationPkiKeyRingResourceModel, 0)
			TagTagPkiKeyRingList := make([]TagTagPkiKeyRingResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "tagAnnotation" {
							TagAnnotationPkiKeyRing := getEmptyTagAnnotationPkiKeyRingResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationPkiKeyRing.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationPkiKeyRing.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}

							}
							TagAnnotationPkiKeyRingList = append(TagAnnotationPkiKeyRingList, TagAnnotationPkiKeyRing)
						}
						if childClassName == "tagTag" {
							TagTagPkiKeyRing := getEmptyTagTagPkiKeyRingResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagTagPkiKeyRing.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagTagPkiKeyRing.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}

							}
							TagTagPkiKeyRingList = append(TagTagPkiKeyRingList, TagTagPkiKeyRing)
						}
					}
				}
			}
			tagAnnotationSet, _ := types.SetValueFrom(ctx, readData.TagAnnotation.ElementType(ctx), TagAnnotationPkiKeyRingList)
			readData.TagAnnotation = tagAnnotationSet
			tagTagSet, _ := types.SetValueFrom(ctx, readData.TagTag.ElementType(ctx), TagTagPkiKeyRingList)
			readData.TagTag = tagTagSet
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'pkiKeyRing'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	} else {
		readData.Id = basetypes.NewStringNull()
	}
	*data = *readData
}

func getPkiKeyRingRn(ctx context.Context, data *PkiKeyRingResourceModel) string {
	return fmt.Sprintf("keyring-%s", data.Name.ValueString())
}

func setPkiKeyRingParentDn(ctx context.Context, dn string, data *PkiKeyRingResourceModel) {
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
	parentDn := dn[:rnIndex]
	rnMap := map[string]string{
		"tn": "certstore",
	}
	for parentIdentifier, rnPrepend := range rnMap {
		if strings.Contains(parentDn, parentIdentifier) {
			parentDn = parentDn[:strings.Index(parentDn, fmt.Sprintf("/%s", rnPrepend))]
			break
		}
	}
	data.ParentDn = basetypes.NewStringValue(parentDn)
}

func setPkiKeyRingId(ctx context.Context, data *PkiKeyRingResourceModel) {
	rn := getPkiKeyRingRn(ctx, data)
	parentDn := data.ParentDn.ValueString()
	rnMap := map[string]string{
		"tn": "certstore",
	}
	id := fmt.Sprintf("%s/%s", parentDn, rn)
	for parentIdentifier, rnPrepend := range rnMap {
		if strings.Contains(parentDn, parentIdentifier) {
			id = fmt.Sprintf("%s/%s/%s", parentDn, rnPrepend, rn)
			break
		}
	}
	data.Id = types.StringValue(id)
}

func getPkiKeyRingTagAnnotationChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *PkiKeyRingResourceModel, tagAnnotationPkiKeyRingPlan, tagAnnotationPkiKeyRingState []TagAnnotationPkiKeyRingResourceModel) []map[string]interface{} {
	childPayloads := []map[string]interface{}{}
	if !data.TagAnnotation.IsNull() && !data.TagAnnotation.IsUnknown() {
		tagAnnotationIdentifiers := []TagAnnotationIdentifier{}
		for _, tagAnnotationPkiKeyRing := range tagAnnotationPkiKeyRingPlan {
			childMap := NewAciObject()
			if !tagAnnotationPkiKeyRing.Key.IsNull() && !tagAnnotationPkiKeyRing.Key.IsUnknown() {
				childMap.Attributes["key"] = tagAnnotationPkiKeyRing.Key.ValueString()
			}
			if !tagAnnotationPkiKeyRing.Value.IsNull() && !tagAnnotationPkiKeyRing.Value.IsUnknown() {
				childMap.Attributes["value"] = tagAnnotationPkiKeyRing.Value.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"tagAnnotation": childMap})
			tagAnnotationIdentifier := TagAnnotationIdentifier{}
			tagAnnotationIdentifier.Key = tagAnnotationPkiKeyRing.Key
			tagAnnotationIdentifiers = append(tagAnnotationIdentifiers, tagAnnotationIdentifier)
		}
		for _, tagAnnotation := range tagAnnotationPkiKeyRingState {
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

func getPkiKeyRingTagTagChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *PkiKeyRingResourceModel, tagTagPkiKeyRingPlan, tagTagPkiKeyRingState []TagTagPkiKeyRingResourceModel) []map[string]interface{} {
	childPayloads := []map[string]interface{}{}
	if !data.TagTag.IsNull() && !data.TagTag.IsUnknown() {
		tagTagIdentifiers := []TagTagIdentifier{}
		for _, tagTagPkiKeyRing := range tagTagPkiKeyRingPlan {
			childMap := NewAciObject()
			if !tagTagPkiKeyRing.Key.IsNull() && !tagTagPkiKeyRing.Key.IsUnknown() {
				childMap.Attributes["key"] = tagTagPkiKeyRing.Key.ValueString()
			}
			if !tagTagPkiKeyRing.Value.IsNull() && !tagTagPkiKeyRing.Value.IsUnknown() {
				childMap.Attributes["value"] = tagTagPkiKeyRing.Value.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"tagTag": childMap})
			tagTagIdentifier := TagTagIdentifier{}
			tagTagIdentifier.Key = tagTagPkiKeyRing.Key
			tagTagIdentifiers = append(tagTagIdentifiers, tagTagIdentifier)
		}
		for _, tagTag := range tagTagPkiKeyRingState {
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

func getPkiKeyRingCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, createType bool, data *PkiKeyRingResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationPkiKeyRingResourceModel, tagTagPlan, tagTagState []TagTagPkiKeyRingResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}

	if createType && !globalAllowExistingOnCreate {
		payloadMap["attributes"].(map[string]string)["status"] = "created"
	}
	childPayloads := []map[string]interface{}{}

	TagAnnotationchildPayloads := getPkiKeyRingTagAnnotationChildPayloads(ctx, diags, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)

	TagTagchildPayloads := getPkiKeyRingTagTagChildPayloads(ctx, diags, data, tagTagPlan, tagTagState)
	if TagTagchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagTagchildPayloads...)

	payloadMap["children"] = childPayloads
	if !data.AdminState.IsNull() && !data.AdminState.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["adminState"] = data.AdminState.ValueString()
	}
	if !data.Annotation.IsNull() && !data.Annotation.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["annotation"] = data.Annotation.ValueString()
	}
	if !data.Cert.IsNull() && !data.Cert.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["cert"] = data.Cert.ValueString()
	}
	if !data.Descr.IsNull() && !data.Descr.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["descr"] = data.Descr.ValueString()
	}
	if !data.EccCurve.IsNull() && !data.EccCurve.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["eccCurve"] = data.EccCurve.ValueString()
	}
	if !data.Key.IsNull() && !data.Key.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["key"] = data.Key.ValueString()
	}
	if !data.KeyType.IsNull() && !data.KeyType.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["keyType"] = data.KeyType.ValueString()
	}
	if !data.Modulus.IsNull() && !data.Modulus.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["modulus"] = data.Modulus.ValueString()
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
	if !data.Regen.IsNull() && !data.Regen.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["regen"] = data.Regen.ValueString()
	}
	if !data.Tp.IsNull() && !data.Tp.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["tp"] = data.Tp.ValueString()
	}
	wrapperClassMap := map[string]string{
		"uni/userext/pkiext": "",
		"certstore":          "cloudCertStore",
	}

	var payload []byte
	var err error
	for rnPrepend, wrapperClass := range wrapperClassMap {
		if strings.Contains(data.Id.ValueString(), rnPrepend) && wrapperClass != "" && createType {
			wrapperPayloadMap := map[string]interface{}{
				wrapperClass: map[string]interface{}{
					"attributes": map[string]interface{}{},
					"children":   []interface{}{map[string]interface{}{"pkiKeyRing": payloadMap}},
				},
			}
			payload, err = json.Marshal(wrapperPayloadMap)
			break
		} else if (strings.Contains(data.Id.ValueString(), rnPrepend) && wrapperClass == "") || !createType {
			payload, err = json.Marshal(map[string]interface{}{"pkiKeyRing": payloadMap})
			break
		}
	}
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
