// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceL3extRsOutToFBRGroupWithL3extOut(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigL3extRsOutToFBRGroupDataSourceDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_l3out_fallback_route_group.test", "target_dn", "uni/tn-test_tenant/ctx-test_vrf/fbrg-fallback_route_group"),
					resource.TestCheckResourceAttr("data.aci_l3out_fallback_route_group.test", "annotation", "orchestrator:terraform"),
				),
			},
			{
				Config:      testConfigL3extRsOutToFBRGroupNotExistingL3extOut,
				ExpectError: regexp.MustCompile("Failed to read aci_l3out_fallback_route_group data source"),
			},
		},
	})
}

const testConfigL3extRsOutToFBRGroupDataSourceDependencyWithL3extOut = testConfigL3extRsOutToFBRGroupMinDependencyWithL3extOut + `
data "aci_l3out_fallback_route_group" "test" {
  parent_dn = aci_l3_outside.test.id
  target_dn = aci_vrf_fallback_route_group.test.id
  depends_on = [aci_l3out_fallback_route_group.test]
}
`

const testConfigL3extRsOutToFBRGroupNotExistingL3extOut = testConfigL3extRsOutToFBRGroupMinDependencyWithL3extOut + `
data "aci_l3out_fallback_route_group" "test_non_existing" {
  parent_dn = aci_l3_outside.test.id
  target_dn = "uni/tn-test_tenant/ctx-test_vrf/fbrg-fallback_route_group_not_existing"
  depends_on = [aci_l3out_fallback_route_group.test]
}
`
