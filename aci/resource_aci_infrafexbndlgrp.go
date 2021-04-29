package aci

import (
	"fmt"
	"log"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAciFexBundleGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciFexBundleGroupCreate,
		Update: resourceAciFexBundleGroupUpdate,
		Read:   resourceAciFexBundleGroupRead,
		Delete: resourceAciFexBundleGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciFexBundleGroupImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"fex_profile_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"name_alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"relation_infra_rs_mon_fex_infra_pol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"relation_infra_rs_fex_bndl_grp_to_aggr_if": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set:      schema.HashString,
			},
		}),
	}
}
func getRemoteFexBundleGroup(client *client.Client, dn string) (*models.FexBundleGroup, error) {
	infraFexBndlGrpCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	infraFexBndlGrp := models.FexBundleGroupFromContainer(infraFexBndlGrpCont)

	if infraFexBndlGrp.DistinguishedName == "" {
		return nil, fmt.Errorf("FexBundleGroup %s not found", infraFexBndlGrp.DistinguishedName)
	}

	return infraFexBndlGrp, nil
}

func setFexBundleGroupAttributes(infraFexBndlGrp *models.FexBundleGroup, d *schema.ResourceData) *schema.ResourceData {
	dn := d.Id()
	d.SetId(infraFexBndlGrp.DistinguishedName)
	d.Set("description", infraFexBndlGrp.Description)
	if dn != infraFexBndlGrp.DistinguishedName {
		d.Set("fex_profile_dn", "")
	}
	infraFexBndlGrpMap, _ := infraFexBndlGrp.ToMap()

	d.Set("name", infraFexBndlGrpMap["name"])

	d.Set("annotation", infraFexBndlGrpMap["annotation"])
	d.Set("name_alias", infraFexBndlGrpMap["nameAlias"])
	return d
}

func resourceAciFexBundleGroupImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())
	aciClient := m.(*client.Client)

	dn := d.Id()

	infraFexBndlGrp, err := getRemoteFexBundleGroup(aciClient, dn)

	if err != nil {
		return nil, err
	}
	infraFexBndlGrpMap, _ := infraFexBndlGrp.ToMap()

	name := infraFexBndlGrpMap["name"]
	pDN := GetParentDn(dn, fmt.Sprintf("/fexbundle-%s", name))
	d.Set("fex_profile_dn", pDN)
	schemaFilled := setFexBundleGroupAttributes(infraFexBndlGrp, d)

	log.Printf("[DEBUG] %s: Import finished successfully", d.Id())

	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciFexBundleGroupCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] FexBundleGroup: Beginning Creation")
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	FEXProfileDn := d.Get("fex_profile_dn").(string)

	infraFexBndlGrpAttr := models.FexBundleGroupAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		infraFexBndlGrpAttr.Annotation = Annotation.(string)
	} else {
		infraFexBndlGrpAttr.Annotation = "{}"
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		infraFexBndlGrpAttr.NameAlias = NameAlias.(string)
	}
	infraFexBndlGrp := models.NewFexBundleGroup(fmt.Sprintf("fexbundle-%s", name), FEXProfileDn, desc, infraFexBndlGrpAttr)

	err := aciClient.Save(infraFexBndlGrp)
	if err != nil {
		return err
	}
	d.Partial(true)

	d.Partial(false)

	checkDns := make([]string, 0, 1)

	if relationToinfraRsMonFexInfraPol, ok := d.GetOk("relation_infra_rs_mon_fex_infra_pol"); ok {
		relationParam := relationToinfraRsMonFexInfraPol.(string)
		checkDns = append(checkDns, relationParam)
	}

	if relationToinfraRsFexBndlGrpToAggrIf, ok := d.GetOk("relation_infra_rs_fex_bndl_grp_to_aggr_if"); ok {
		relationParamList := toStringList(relationToinfraRsFexBndlGrpToAggrIf.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			checkDns = append(checkDns, relationParam)
		}
	}

	d.Partial(true)
	err = checkTDn(aciClient, checkDns)
	if err != nil {
		return err
	}
	d.Partial(false)

	if relationToinfraRsMonFexInfraPol, ok := d.GetOk("relation_infra_rs_mon_fex_infra_pol"); ok {
		relationParam := relationToinfraRsMonFexInfraPol.(string)
		relationParamName := GetMOName(relationParam)
		err = aciClient.CreateRelationinfraRsMonFexInfraPolFromFexBundleGroup(infraFexBndlGrp.DistinguishedName, relationParamName)
		if err != nil {
			return err
		}
		d.Partial(true)
		d.Partial(false)

	}
	if relationToinfraRsFexBndlGrpToAggrIf, ok := d.GetOk("relation_infra_rs_fex_bndl_grp_to_aggr_if"); ok {
		relationParamList := toStringList(relationToinfraRsFexBndlGrpToAggrIf.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationinfraRsFexBndlGrpToAggrIfFromFexBundleGroup(infraFexBndlGrp.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
			d.Partial(true)
			d.Partial(false)
		}
	}

	d.SetId(infraFexBndlGrp.DistinguishedName)
	log.Printf("[DEBUG] %s: Creation finished successfully", d.Id())

	return resourceAciFexBundleGroupRead(d, m)
}

func resourceAciFexBundleGroupUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] FexBundleGroup: Beginning Update")

	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	FEXProfileDn := d.Get("fex_profile_dn").(string)

	infraFexBndlGrpAttr := models.FexBundleGroupAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		infraFexBndlGrpAttr.Annotation = Annotation.(string)
	} else {
		infraFexBndlGrpAttr.Annotation = "{}"
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		infraFexBndlGrpAttr.NameAlias = NameAlias.(string)
	}
	infraFexBndlGrp := models.NewFexBundleGroup(fmt.Sprintf("fexbundle-%s", name), FEXProfileDn, desc, infraFexBndlGrpAttr)

	infraFexBndlGrp.Status = "modified"

	err := aciClient.Save(infraFexBndlGrp)

	if err != nil {
		return err
	}
	d.Partial(true)

	d.Partial(false)

	checkDns := make([]string, 0, 1)

	if d.HasChange("relation_infra_rs_mon_fex_infra_pol") {
		_, newRelParam := d.GetChange("relation_infra_rs_mon_fex_infra_pol")
		checkDns = append(checkDns, newRelParam.(string))
	}

	if d.HasChange("relation_infra_rs_fex_bndl_grp_to_aggr_if") {
		oldRel, newRel := d.GetChange("relation_infra_rs_fex_bndl_grp_to_aggr_if")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToCreate {
			checkDns = append(checkDns, relDn)
		}
	}

	d.Partial(true)
	err = checkTDn(aciClient, checkDns)
	if err != nil {
		return err
	}
	d.Partial(false)

	if d.HasChange("relation_infra_rs_mon_fex_infra_pol") {
		_, newRelParam := d.GetChange("relation_infra_rs_mon_fex_infra_pol")
		newRelParamName := GetMOName(newRelParam.(string))
		err = aciClient.CreateRelationinfraRsMonFexInfraPolFromFexBundleGroup(infraFexBndlGrp.DistinguishedName, newRelParamName)
		if err != nil {
			return err
		}
		d.Partial(true)
		d.Partial(false)

	}
	if d.HasChange("relation_infra_rs_fex_bndl_grp_to_aggr_if") {
		oldRel, newRel := d.GetChange("relation_infra_rs_fex_bndl_grp_to_aggr_if")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationinfraRsFexBndlGrpToAggrIfFromFexBundleGroup(infraFexBndlGrp.DistinguishedName, relDn)
			if err != nil {
				return err
			}
			d.Partial(true)
			d.Partial(false)

		}

	}

	d.SetId(infraFexBndlGrp.DistinguishedName)
	log.Printf("[DEBUG] %s: Update finished successfully", d.Id())

	return resourceAciFexBundleGroupRead(d, m)

}

func resourceAciFexBundleGroupRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Read", d.Id())

	aciClient := m.(*client.Client)

	dn := d.Id()
	infraFexBndlGrp, err := getRemoteFexBundleGroup(aciClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setFexBundleGroupAttributes(infraFexBndlGrp, d)

	infraRsMonFexInfraPolData, err := aciClient.ReadRelationinfraRsMonFexInfraPolFromFexBundleGroup(dn)
	if err != nil {
		log.Printf("[DEBUG] Error while reading relation infraRsMonFexInfraPol %v", err)
		d.Set("relation_infra_rs_mon_fex_infra_pol", "")

	} else {
		paramName := GetMOName(d.Get("relation_infra_rs_mon_fex_infra_pol").(string))
		if paramName != infraRsMonFexInfraPolData {
			d.Set("relation_infra_rs_mon_fex_infra_pol", "")
		}
	}

	infraRsFexBndlGrpToAggrIfData, err := aciClient.ReadRelationinfraRsFexBndlGrpToAggrIfFromFexBundleGroup(dn)
	if err != nil {
		log.Printf("[DEBUG] Error while reading relation infraRsFexBndlGrpToAggrIf %v", err)
		d.Set("relation_infra_rs_fex_bndl_grp_to_aggr_if", make([]string, 0, 1))

	} else {
		d.Set("relation_infra_rs_fex_bndl_grp_to_aggr_if", infraRsFexBndlGrpToAggrIfData)
	}

	log.Printf("[DEBUG] %s: Read finished successfully", d.Id())

	return nil
}

func resourceAciFexBundleGroupDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Destroy", d.Id())

	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "infraFexBndlGrp")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] %s: Destroy finished successfully", d.Id())

	d.SetId("")
	return err
}
