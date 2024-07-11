// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvRemoteIdWithFvSiteAssociated(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvRemoteIdDataSourceDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_remote_site_id_mappings.test", "site_id", "100"),
					resource.TestCheckResourceAttr("data.aci_remote_site_id_mappings.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_remote_site_id_mappings.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_remote_site_id_mappings.test", "name", ""),
					resource.TestCheckResourceAttr("data.aci_remote_site_id_mappings.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_remote_site_id_mappings.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_remote_site_id_mappings.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("data.aci_remote_site_id_mappings.test", "remote_pc_tag", "16386"),
					resource.TestCheckResourceAttr("data.aci_remote_site_id_mappings.test", "remote_vrf_pc_tag", "2818057"),
					resource.TestCheckResourceAttr("data.aci_remote_site_id_mappings.test", "site_id", "100"),
				),
			},
			{
				Config:      testConfigFvRemoteIdNotExistingFvSiteAssociated,
				ExpectError: regexp.MustCompile("Failed to read aci_remote_site_id_mappings data source"),
			},
		},
	})
}

const testConfigFvRemoteIdDataSourceDependencyWithFvSiteAssociated = testConfigFvRemoteIdMinDependencyWithFvSiteAssociated + `
data "aci_remote_site_id_mappings" "test" {
  parent_dn = aci_associated_site.test.id
  site_id = "100"
  depends_on = [aci_remote_site_id_mappings.test]
}
`

const testConfigFvRemoteIdNotExistingFvSiteAssociated = testConfigFvRemoteIdMinDependencyWithFvSiteAssociated + `
data "aci_remote_site_id_mappings" "test_non_existing" {
  parent_dn = aci_associated_site.test.id
  site_id = "102"
  depends_on = [aci_remote_site_id_mappings.test]
}
`
