// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceTagAnnotationWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigTagAnnotationDataSourceDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_annotation.test", "key", "test_key"),
					resource.TestCheckResourceAttr("data.aci_annotation.test", "value", "test_value"),
				),
			},
			{
				Config:      testConfigTagAnnotationNotExistingFvTenant,
				ExpectError: regexp.MustCompile("Failed to read aci_annotation data source"),
			},
		},
	})
}
func TestAccDataSourceTagAnnotationWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigTagAnnotationDataSourceDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_annotation.test", "key", "test_key"),
					resource.TestCheckResourceAttr("data.aci_annotation.test", "value", "test_value"),
				),
			},
			{
				Config:      testConfigTagAnnotationNotExistingFvAEPg,
				ExpectError: regexp.MustCompile("Failed to read aci_annotation data source"),
			},
		},
	})
}

const testConfigTagAnnotationDataSourceDependencyWithFvTenant = testConfigTagAnnotationMinDependencyWithFvTenant + `
data "aci_annotation" "test" {
  parent_dn = aci_tenant.test.id
  key = "test_key"
  depends_on = [aci_annotation.test]
}
`

const testConfigTagAnnotationNotExistingFvTenant = testConfigTagAnnotationMinDependencyWithFvTenant + `
data "aci_annotation" "test_non_existing" {
  parent_dn = aci_tenant.test.id
  key = "non_existing_key"
  depends_on = [aci_annotation.test]
}
`
const testConfigTagAnnotationDataSourceDependencyWithFvAEPg = testConfigTagAnnotationMinDependencyWithFvAEPg + `
data "aci_annotation" "test" {
  parent_dn = aci_application_epg.test.id
  key = "test_key"
  depends_on = [aci_annotation.test]
}
`

const testConfigTagAnnotationNotExistingFvAEPg = testConfigTagAnnotationMinDependencyWithFvAEPg + `
data "aci_annotation" "test_non_existing" {
  parent_dn = aci_application_epg.test.id
  key = "non_existing_key"
  depends_on = [aci_annotation.test]
}
`
