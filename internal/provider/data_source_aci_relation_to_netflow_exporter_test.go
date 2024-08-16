// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceNetflowRsMonitorToExporterWithNetflowMonitorPol(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigNetflowRsMonitorToExporterDataSourceDependencyWithNetflowMonitorPol,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_relation_to_netflow_exporter.test", "netflow_exporter_policy_name", "test_tn_netflow_exporter_pol_name"),
					resource.TestCheckResourceAttr("data.aci_relation_to_netflow_exporter.test", "annotation", "orchestrator:terraform"),
				),
			},
			{
				Config:      testConfigNetflowRsMonitorToExporterNotExistingNetflowMonitorPol,
				ExpectError: regexp.MustCompile("Failed to read aci_relation_to_netflow_exporter data source"),
			},
		},
	})
}

const testConfigNetflowRsMonitorToExporterDataSourceDependencyWithNetflowMonitorPol = testConfigNetflowRsMonitorToExporterMinDependencyWithNetflowMonitorPol + `
data "aci_relation_to_netflow_exporter" "test" {
  parent_dn = aci_netflow_monitor_policy.test.id
  netflow_exporter_policy_name = "test_tn_netflow_exporter_pol_name"
  depends_on = [aci_relation_to_netflow_exporter.test]
}
`

const testConfigNetflowRsMonitorToExporterNotExistingNetflowMonitorPol = testConfigNetflowMonitorPolMinDependencyWithFvTenant + `
data "aci_relation_to_netflow_exporter" "test_non_existing" {
  parent_dn = aci_netflow_monitor_policy.test.id
  netflow_exporter_policy_name = "non_existing_tn_netflow_exporter_pol_name"
}
`
