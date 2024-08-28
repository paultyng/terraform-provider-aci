// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvRsPathAttWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvRsPathAttDataSourceDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_relation_to_static_path.test", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("data.aci_relation_to_static_path.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_relation_to_static_path.test", "deployment_immediacy", "lazy"),
					resource.TestCheckResourceAttr("data.aci_relation_to_static_path.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_relation_to_static_path.test", "encapsulation", "vlan-201"),
					resource.TestCheckResourceAttr("data.aci_relation_to_static_path.test", "mode", "regular"),
					resource.TestCheckResourceAttr("data.aci_relation_to_static_path.test", "primary_encapsulation", "unknown"),
				),
			},
			{
				Config:      testConfigFvRsPathAttNotExistingFvAEPg,
				ExpectError: regexp.MustCompile("Failed to read aci_relation_to_static_path data source"),
			},
		},
	})
}

const testConfigFvRsPathAttDataSourceDependencyWithFvAEPg = testConfigFvRsPathAttMinDependencyWithFvAEPg + `
data "aci_relation_to_static_path" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
  depends_on = [aci_relation_to_static_path.test]
}
`

const testConfigFvRsPathAttNotExistingFvAEPg = testConfigFvRsPathAttMinDependencyWithFvAEPg + `
data "aci_relation_to_static_path" "test_non_existing" {
  parent_dn = aci_application_epg.test.id
  target_dn = "topology/pod-2/paths-201/pathep-[eth1/1]"
  depends_on = [aci_relation_to_static_path.test]
}
`
