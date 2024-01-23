// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceQosDppPolWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.2(2g)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigQosDppPolDataSourceDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "admin_state", "disabled"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "burst", "unspecified"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "burst_unit", "unspecified"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "conform_action", "transmit"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "conform_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "conform_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "exceed_action", "drop"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "exceed_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "exceed_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "excessive_burst", "unspecified"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "excessive_burst_unit", "unspecified"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "mode", "bit"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "peak_rate", "0"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "peak_rate_unit", "unspecified"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "rate", "0"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "rate_unit", "unspecified"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "sharing_mode", "dedicated"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "type", "1R2C"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "violate_action", "drop"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "violate_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("data.aci_data_plane_policing_policy.test", "violate_mark_dscp", "unspecified"),
				),
			},
			{
				Config:      testConfigQosDppPolNotExistingFvTenant,
				ExpectError: regexp.MustCompile("Failed to read aci_data_plane_policing_policy data source"),
			},
		},
	})
}

const testConfigQosDppPolDataSourceDependencyWithFvTenant = testConfigQosDppPolMinDependencyWithFvTenant + `
data "aci_data_plane_policing_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  depends_on = [aci_data_plane_policing_policy.test]
}
`

const testConfigQosDppPolNotExistingFvTenant = testConfigFvTenantMin + `
data "aci_data_plane_policing_policy" "test_non_existing" {
  parent_dn = aci_tenant.test.id
  name = "non_existing_name"
}
`
