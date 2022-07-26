package aci

import (
	"context"
	"fmt"
	"log"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/container"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var className = models.FvaepgClassName

func resourceAciBulkStaticPath() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAciBulkStaticPathCreate,
		UpdateContext: resourceAciBulkStaticPathUpdate,
		ReadContext:   resourceAciBulkStaticPathRead,
		DeleteContext: resourceAciBulkStaticPathDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciBulkStaticPathImport,
		},

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"application_epg_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"static_path": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"interface_dn": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"encap": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"deployment_immediacy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								"immediate",
								"lazy",
							}, false),
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								"regular",
								"native",
								"untagged",
							}, false),
						},
						"primary_encap": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceAciBulkStaticPathCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] BulkStaticPath: Beginning Creation")

	aciClient := m.(*client.Client)

	ApplicationEPGDn := d.Get("application_epg_dn").(string)
	staticPathSet := staticPathPayload(d.Get("static_path").(*schema.Set).List(), "create")

	contentMap := make(map[string]interface{})
	contentMap["dn"] = ApplicationEPGDn
	cont, err := preparePayload(className, toStrMap(contentMap), staticPathSet)
	if err != nil {
		return diag.FromErr(err)
	}

	_, diags := bulkStaticPathRequest(aciClient, "POST", ApplicationEPGDn, cont)
	if diags.HasError() {
		return diags
	}
	d.SetId(ApplicationEPGDn)
	log.Printf("[DEBUG] %s: Creation finished successfully", d.Id())

	return resourceAciBulkStaticPathRead(ctx, d, m)
}

func resourceAciBulkStaticPathRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Read", d.Id())

	aciClient := m.(*client.Client)

	dn := d.Id()
	bulkStaticPath, diags := getAciBulkStaticPath(aciClient, dn)
	if diags.HasError() {
		return diags
	}
	d = setAciBulkStaticPath(bulkStaticPath, d)
	d.Set("application_epg_dn", dn)

	log.Printf("[DEBUG] %s: Read finished successfully", d.Id())

	return nil
}

func resourceAciBulkStaticPathUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Update", d.Id())

	aciClient := m.(*client.Client)

	ApplicationEPGDn := d.Get("application_epg_dn").(string)

	old_static_path, new_static_path := d.GetChange("static_path")
	del := diffInStaticPath(old_static_path.(*schema.Set), new_static_path.(*schema.Set))

	createPayload := staticPathPayload(new_static_path.(*schema.Set).List(), "create")
	deletePayload := staticPathPayload(del, "delete")
	var payload []interface{}
	for _, createObj := range createPayload {
		payload = append(payload, createObj)
	}
	for _, deleteObj := range deletePayload {
		payload = append(payload, deleteObj)
	}

	contentMap := make(map[string]interface{})
	contentMap["dn"] = ApplicationEPGDn

	cont, err := preparePayload(className, toStrMap(contentMap), payload)
	if err != nil {
		return diag.FromErr(err)
	}

	_, diags := bulkStaticPathRequest(aciClient, "POST", ApplicationEPGDn, cont)
	if diags.HasError() {
		return diags
	}
	d.SetId(ApplicationEPGDn)
	log.Printf("[DEBUG] %s: Update finished successfully", d.Id())
	return resourceAciBulkStaticPathRead(ctx, d, m)
}

func getAciBulkStaticPath(aciClient *client.Client, dn string) ([]interface{}, diag.Diagnostics) {
	resp, diags := bulkStaticPathRequest(aciClient, "GET", dn, nil)
	if diags.HasError() {
		return nil, diags
	}
	staticPathSet := resp.S("imdata").S("fvRsPathAtt", "attributes")
	if staticPathSet == nil {
		return nil, nil
	}
	return staticPathSet.Data().([]interface{}), nil
}

func setAciBulkStaticPath(staticPathSet []interface{}, d *schema.ResourceData) *schema.ResourceData {
	stateStaticPathSet := make([]interface{}, 0, 1)
	for _, staticPath := range staticPathSet {
		staticPathMap := staticPath.(map[string]interface{})
		stateStaticPathMap := make(map[string]interface{})
		if descr, ok := staticPathMap["descr"]; ok {
			stateStaticPathMap["description"] = descr.(string)
		}
		if immediacy, ok := staticPathMap["instrImedcy"]; ok {
			stateStaticPathMap["deployment_immediacy"] = immediacy.(string)
		}
		if mode, ok := staticPathMap["mode"]; ok {
			stateStaticPathMap["mode"] = mode.(string)
		}
		if primaryEncap, ok := staticPathMap["primaryEncap"]; ok {
			stateStaticPathMap["primary_encap"] = primaryEncap.(string)
		}
		stateStaticPathMap["encap"] = staticPathMap["encap"].(string)
		stateStaticPathMap["interface_dn"] = staticPathMap["tDn"].(string)
		stateStaticPathSet = append(stateStaticPathSet, stateStaticPathMap)
	}
	d.Set("static_path", stateStaticPathSet)
	return d
}

func bulkStaticPathRequest(aciClient *client.Client, method string, epgDn string, body *container.Container) (*container.Container, diag.Diagnostics) {
	url := "/api/mo/" + epgDn + ".json"
	if method == "GET" {
		url += "?query-target=children&target-subtree-class=fvRsPathAtt"
	}
	req, err := aciClient.MakeRestRequest(method, url, body, true)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	respCont, _, err := aciClient.Do(req)
	if err != nil {
		return respCont, diag.FromErr(err)
	}
	err = client.CheckForErrors(respCont, method, false)
	if err != nil {
		return respCont, diag.FromErr(err)
	}
	if method == "POST" {
		return body, nil
	} else {
		return respCont, nil
	}
}

func diffInStaticPath(oldStaticPaths, newStaticPaths *schema.Set) []interface{} {
	var toDelete []interface{}

	for _, old := range oldStaticPaths.List() {
		found := false

		for _, new := range newStaticPaths.List() {
			if old.(map[string]interface{})["interface_dn"] == new.(map[string]interface{})["interface_dn"] {
				found = true
				break
			}
		}
		if !found {
			toDelete = append(toDelete, old)
		}
	}

	return toDelete
}

func staticPathPayload(staticPathList []interface{}, status string) []interface{} {
	staticPathSet := make([]interface{}, 0, 1)
	for _, staticPath := range staticPathList {
		staticPathMap := make(map[string]interface{})
		staticPathMap["class_name"] = "fvRsPathAtt"
		staticPathContent := make(map[string]interface{})
		staticPath := staticPath.(map[string]interface{})
		staticPathContent["encap"] = staticPath["encap"]
		staticPathContent["tDn"] = staticPath["interface_dn"]
		log.Printf("[DEBUG]: interface_dn in staticPathPayload %s", staticPath["interface_dn"])
		if descr, ok := staticPath["description"]; ok {
			staticPathContent["descr"] = descr
		}
		if immediacy, ok := staticPath["deployment_immediacy"]; ok {
			staticPathContent["instrImedcy"] = immediacy
		}
		if mode, ok := staticPath["mode"]; ok {
			staticPathContent["mode"] = mode
		}
		if primaryEncap, ok := staticPath["primary_encap"]; ok {
			staticPathContent["primaryEncap"] = primaryEncap
		}
		if status == "delete" {
			staticPathContent["status"] = "deleted"
		}
		log.Printf("[DEBUG]: primaryEncap in staticPathPayload %s", staticPath["primary_encap"])
		log.Printf("[DEBUG]: staticPathContent in staticPathPayload %s", staticPathContent)
		staticPathMap["content"] = toStrMap(staticPathContent)
		staticPathSet = append(staticPathSet, staticPathMap)
	}
	return staticPathSet
}

func resourceAciBulkStaticPathDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Destroy", d.Id())

	aciClient := m.(*client.Client)
	ApplicationEPGDn := d.Get("application_epg_dn").(string)

	dn := d.Id()
	bulkStaticPath, diags := getAciBulkStaticPath(aciClient, dn)
	if diags.HasError() {
		return diags
	}
	d = setAciBulkStaticPath(bulkStaticPath, d)

	log.Printf("[DEBUG]: BulkStaticPath in Delete: %s", bulkStaticPath)
	deletePayload := staticPathPayload(d.Get("static_path").(*schema.Set).List(), "delete")

	contentMap := make(map[string]interface{})
	contentMap["dn"] = ApplicationEPGDn
	cont, err := preparePayload(className, toStrMap(contentMap), deletePayload)
	if err != nil {
		return diag.FromErr(err)
	}

	_, diags = bulkStaticPathRequest(aciClient, "POST", ApplicationEPGDn, cont)
	if diags.HasError() {
		return diags
	}

	d.SetId("")
	log.Printf("[DEBUG] %s: Destroy finished successfully", d.Id())
	return nil
}

func resourceAciBulkStaticPathImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())

	aciClient := m.(*client.Client)
	dn := d.Id()
	d.Set("application_epg_dn", dn)

	bulkStaticPath, diags := getAciBulkStaticPath(aciClient, dn)
	if diags.HasError() {
		return nil, fmt.Errorf("could not read static path object when importing: %s", diags[0].Summary)
	}
	d = setAciBulkStaticPath(bulkStaticPath, d)

	log.Printf("[DEBUG] %s: Import finished successfully", d.Id())
	return []*schema.ResourceData{d}, nil
}
