// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceL3extRsLblToInstPWithL3extConsLbl(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigL3extRsLblToInstPDataSourceDependencyWithL3extConsLbl,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_relation_to_external_network_instance_profile.test", "target_dn", "uni/tn-test_tenant/out-test_l3_outside/instP-inst_profile"),
					resource.TestCheckResourceAttr("data.aci_relation_to_external_network_instance_profile.test", "annotation", "orchestrator:terraform"),
				),
			},
			{
				Config:      testConfigL3extRsLblToInstPNotExistingL3extConsLbl,
				ExpectError: regexp.MustCompile("Failed to read aci_relation_to_external_network_instance_profile data source"),
			},
		},
	})
}

const testConfigL3extRsLblToInstPDataSourceDependencyWithL3extConsLbl = testConfigL3extRsLblToInstPMinDependencyWithL3extConsLbl + `
data "aci_relation_to_external_network_instance_profile" "test" {
  parent_dn = aci_l3out_consumer_label.test.id
  target_dn = "uni/tn-test_tenant/out-test_l3_outside/instP-inst_profile"
  depends_on = [aci_relation_to_external_network_instance_profile.test]
}
`

const testConfigL3extRsLblToInstPNotExistingL3extConsLbl = testConfigL3extConsLblMinDependencyWithL3extOut + `
data "aci_relation_to_external_network_instance_profile" "test_non_existing" {
  parent_dn = aci_l3out_consumer_label.test.id
  target_dn = "uni/tn-test_tenant/out-test_l3_outside/instP-inst_profile_not_existing"
}
`
