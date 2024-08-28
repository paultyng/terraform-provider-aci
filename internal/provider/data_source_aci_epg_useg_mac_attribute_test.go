// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvMacAttrWithFvCrtrn(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvMacAttrDataSourceDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_epg_useg_mac_attribute.test", "name", "mac_attr"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_mac_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_mac_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_mac_attribute.test", "mac", "AA:BB:CC:DD:EE:FF"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_mac_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_mac_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_mac_attribute.test", "owner_tag", ""),
				),
			},
			{
				Config:      testConfigFvMacAttrNotExistingFvCrtrn,
				ExpectError: regexp.MustCompile("Failed to read aci_epg_useg_mac_attribute data source"),
			},
		},
	})
}

const testConfigFvMacAttrDataSourceDependencyWithFvCrtrn = testConfigFvMacAttrMinDependencyWithFvCrtrn + `
data "aci_epg_useg_mac_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  name = "mac_attr"
  depends_on = [aci_epg_useg_mac_attribute.test]
}
`

const testConfigFvMacAttrNotExistingFvCrtrn = testConfigFvMacAttrMinDependencyWithFvCrtrn + `
data "aci_epg_useg_mac_attribute" "test_non_existing" {
  parent_dn = aci_epg_useg_block_statement.test.id
  name = "mac_attr_non_existing"
  depends_on = [aci_epg_useg_mac_attribute.test]
}
`
