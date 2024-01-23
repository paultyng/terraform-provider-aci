// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceNetflowExporterPolWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "2.2(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigNetflowExporterPolDataSourceDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_netflow_exporter_policy.test", "name", "netfow_exporter"),
					resource.TestCheckResourceAttr("data.aci_netflow_exporter_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_netflow_exporter_policy.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_netflow_exporter_policy.test", "destination_ip_address", "2.2.2.1"),
					resource.TestCheckResourceAttr("data.aci_netflow_exporter_policy.test", "destination_port", "https"),
					resource.TestCheckResourceAttr("data.aci_netflow_exporter_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_netflow_exporter_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_netflow_exporter_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("data.aci_netflow_exporter_policy.test", "qos_dscp_value", "CS2"),
					resource.TestCheckResourceAttr("data.aci_netflow_exporter_policy.test", "source_ip_address", "1.1.1.1/10"),
					resource.TestCheckResourceAttr("data.aci_netflow_exporter_policy.test", "source_ip_type", "custom-src-ip"),
					resource.TestCheckResourceAttr("data.aci_netflow_exporter_policy.test", "version", "v9"),
				),
			},
			{
				Config:      testConfigNetflowExporterPolNotExistingFvTenant,
				ExpectError: regexp.MustCompile("Failed to read aci_netflow_exporter_policy data source"),
			},
		},
	})
}

const testConfigNetflowExporterPolDataSourceDependencyWithFvTenant = testConfigNetflowExporterPolMinDependencyWithFvTenant + `
data "aci_netflow_exporter_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_exporter"
  depends_on = [aci_netflow_exporter_policy.test]
}
`

const testConfigNetflowExporterPolNotExistingFvTenant = testConfigFvTenantMin + `
data "aci_netflow_exporter_policy" "test_non_existing" {
  parent_dn = aci_tenant.test.id
  name = "netfow_exporter_non_existing"
}
`
