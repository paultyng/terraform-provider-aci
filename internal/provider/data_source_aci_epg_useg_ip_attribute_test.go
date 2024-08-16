// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvIpAttrWithFvCrtrn(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvIpAttrDataSourceDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_epg_useg_ip_attribute.test", "name", "131"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_ip_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_ip_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_ip_attribute.test", "ip", "131.107.1.200"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_ip_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_ip_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_ip_attribute.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_ip_attribute.test", "use_epg_subnet", "no"),
				),
			},
			{
				Config:      testConfigFvIpAttrNotExistingFvCrtrn,
				ExpectError: regexp.MustCompile("Failed to read aci_epg_useg_ip_attribute data source"),
			},
		},
	})
}

const testConfigFvIpAttrDataSourceDependencyWithFvCrtrn = testConfigFvIpAttrMinDependencyWithFvCrtrn + `
data "aci_epg_useg_ip_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  name = "131"
  depends_on = [aci_epg_useg_ip_attribute.test]
}
`

const testConfigFvIpAttrNotExistingFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
data "aci_epg_useg_ip_attribute" "test_non_existing" {
  parent_dn = aci_epg_useg_block_statement.test.id
  name = "131_non_existing"
}
`
