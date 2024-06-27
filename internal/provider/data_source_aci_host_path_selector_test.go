// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceInfraHPathS(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testConfigInfraHPathSDataSource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_host_path_selector.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("data.aci_host_path_selector.test", "description", "description"),
					resource.TestCheckResourceAttr("data.aci_host_path_selector.test", "name_alias", "name_alias"),
					resource.TestCheckResourceAttr("data.aci_host_path_selector.test", "owner_key", "owner_key"),
					resource.TestCheckResourceAttr("data.aci_host_path_selector.test", "owner_tag", "owner_tag"),
				),
			},
			{
				Config:      testConfigInfraHPathSNotExisting,
				ExpectError: regexp.MustCompile("Failed to read aci_host_path_selector data source"),
			},
		},
	})
}

const testConfigInfraHPathSDataSource = testConfigInfraHPathSAll + `
data "aci_host_path_selector" "test" {
  name = "host_path_selector"
  depends_on = [aci_host_path_selector.test]
}
`

const testConfigInfraHPathSNotExisting = testConfigInfraHPathSAll + `
data "aci_host_path_selector" "test" {
  name = "host_path_selector_non_existing"
  depends_on = [aci_host_path_selector.test]
}
`
