// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourcePkiKeyRing(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigPkiKeyRingDataSource + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "name", "test_name"),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "admin_state", "started"),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "certificate", ""),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "certificate_authority", ""),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "modulus", "mod2048"),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "regenerate", "no"),
					composeAggregateTestCheckFuncWithVersion(t, "6.0(2h)", ">",
						resource.TestCheckResourceAttr("data.aci_key_ring.test", "elliptic_curve", "none"),
						resource.TestCheckResourceAttr("data.aci_key_ring.test", "key_type", "RSA")),
				),
			},
			{
				Config:      testConfigPkiKeyRingNotExisting + testConfigDataSourceSystem,
				ExpectError: regexp.MustCompile("Failed to read aci_key_ring data source"),
			},
		},
	})
}
func TestAccDataSourcePkiKeyRingWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "cloud", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigPkiKeyRingDataSourceDependencyWithFvTenant + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "name", "test_name"),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "admin_state", "started"),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "certificate", ""),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "certificate_authority", ""),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "modulus", "mod2048"),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("data.aci_key_ring.test", "regenerate", "no"),
					composeAggregateTestCheckFuncWithVersion(t, "6.0(2h)", ">",
						resource.TestCheckResourceAttr("data.aci_key_ring.test", "elliptic_curve", "none"),
						resource.TestCheckResourceAttr("data.aci_key_ring.test", "key_type", "RSA")),
				),
			},
			{
				Config:      testConfigPkiKeyRingNotExistingFvTenant + testConfigDataSourceSystem,
				ExpectError: regexp.MustCompile("Failed to read aci_key_ring data source"),
			},
		},
	})
}

const testConfigPkiKeyRingDataSource = testConfigPkiKeyRingMin + `
data "aci_key_ring" "test" {
  name = "test_name"
  depends_on = [aci_key_ring.test]
}
`

const testConfigPkiKeyRingNotExisting = testConfigPkiKeyRingMin + `
data "aci_key_ring" "test_non_existing" {
  name = "non_existing_name"
}
`
const testConfigPkiKeyRingDataSourceDependencyWithFvTenant = testConfigPkiKeyRingMinDependencyWithFvTenant + `
data "aci_key_ring" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  depends_on = [aci_key_ring.test]
}
`

const testConfigPkiKeyRingNotExistingFvTenant = testConfigFvTenantMin + `
data "aci_key_ring" "test_non_existing" {
  parent_dn = aci_tenant.test.id
  name = "non_existing_name"
}
`
