// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvESgWithFvAp(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvESgDataSourceDependencyWithFvAp,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_endpoint_security_group.test", "name", "test_name"),
					resource.TestCheckResourceAttr("data.aci_endpoint_security_group.test", "admin_state", "no"),
					resource.TestCheckResourceAttr("data.aci_endpoint_security_group.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_endpoint_security_group.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_endpoint_security_group.test", "exception_tag", ""),
					resource.TestCheckResourceAttr("data.aci_endpoint_security_group.test", "intra_esg_isolation", "unenforced"),
					resource.TestCheckResourceAttr("data.aci_endpoint_security_group.test", "match_criteria", "AtleastOne"),
					resource.TestCheckResourceAttr("data.aci_endpoint_security_group.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_endpoint_security_group.test", "preferred_group_member", "exclude"),
				),
			},
			{
				Config:      testConfigFvESgNotExistingFvAp,
				ExpectError: regexp.MustCompile("Failed to read aci_endpoint_security_group data source"),
			},
		},
	})
}

const testConfigFvESgDataSourceDependencyWithFvAp = testConfigFvESgMinDependencyWithFvAp + `
data "aci_endpoint_security_group" "test" {
  parent_dn = aci_application_profile.test.id
  name = "test_name"
  depends_on = [aci_endpoint_security_group.test]
}
`

const testConfigFvESgNotExistingFvAp = testConfigFvESgMinDependencyWithFvAp + `
data "aci_endpoint_security_group" "test_non_existing" {
  parent_dn = aci_application_profile.test.id
  name = "non_existing_name"
  depends_on = [aci_endpoint_security_group.test]
}
`
