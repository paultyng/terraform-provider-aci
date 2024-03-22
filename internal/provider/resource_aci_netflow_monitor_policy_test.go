// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceNetflowMonitorPolWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "2.2(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigNetflowMonitorPolMinDependencyWithFvTenantAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test", "name", "netfow_monitor"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test_2", "name", "netfow_monitor"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test_2", "owner_tag", ""),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "2.2(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigNetflowMonitorPolMinDependencyWithFvTenantAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "2.2(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigNetflowMonitorPolMinDependencyWithFvTenantAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test", "name", "netfow_monitor"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test_2", "name", "netfow_monitor"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.allow_test_2", "owner_tag", ""),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "2.2(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigNetflowMonitorPolMinDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "name", "netfow_monitor"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "owner_tag", ""),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigNetflowMonitorPolAllDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "name", "netfow_monitor"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "owner_key", "owner_key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "owner_tag", "owner_tag_1"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigNetflowMonitorPolMinDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "name", "netfow_monitor"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigNetflowMonitorPolResetDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "name", "netfow_monitor"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "owner_tag", ""),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_netflow_monitor_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigNetflowMonitorPolChildrenDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "name", "netfow_monitor"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.netflow_exporter_policy_name", "netflow_exporter_policy_name_0"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.1.netflow_exporter_policy_name", "netflow_exporter_policy_name_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.netflow_record_policy_name", "netflow_record_policy_name_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "tags.1.value", "test_value"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_netflow_monitor_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigNetflowMonitorPolChildrenRemoveFromConfigDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "tags.#", "2"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.netflow_exporter_policy_name", "netflow_exporter_policy_name_0"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.1.annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.1.annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.1.annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.1.annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.1.tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.1.tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.1.tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.1.tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "tags.#", "2"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.1.netflow_exporter_policy_name", "netflow_exporter_policy_name_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.#", "2"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "tags.#", "2"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.netflow_record_policy_name", "netflow_record_policy_name_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigNetflowMonitorPolChildrenRemoveOneDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.tags.#", "1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.0.netflow_exporter_policy_name", "netflow_exporter_policy_name_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.#", "1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.tags.#", "1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.netflow_record_policy_name", "netflow_record_policy_name_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigNetflowMonitorPolChildrenRemoveAllDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_exporters.#", "0"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.tags.#", "0"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "relation_to_netflow_record.netflow_record_policy_name", "netflow_record_policy_name_1"),
					resource.TestCheckResourceAttr("aci_netflow_monitor_policy.test", "tags.#", "0"),
				),
			},
		},
		CheckDestroy: testCheckResourceDestroy,
	})
}

const testConfigNetflowMonitorPolMinDependencyWithFvTenantAllowExisting = testConfigFvTenantMin + `
resource "aci_netflow_monitor_policy" "allow_test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_monitor"
}
resource "aci_netflow_monitor_policy" "allow_test_2" {
  parent_dn = aci_tenant.test.id
  name = "netfow_monitor"
  depends_on = [aci_netflow_monitor_policy.allow_test]
}
`

const testConfigNetflowMonitorPolMinDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_netflow_monitor_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_monitor"
}
`

const testConfigNetflowMonitorPolAllDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_netflow_monitor_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_monitor"
  annotation = "annotation"
  description = "description_1"
  name_alias = "name_alias_1"
  owner_key = "owner_key_1"
  owner_tag = "owner_tag_1"
}
`

const testConfigNetflowMonitorPolResetDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_netflow_monitor_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_monitor"
  annotation = "orchestrator:terraform"
  description = ""
  name_alias = ""
  owner_key = ""
  owner_tag = ""
}
`
const testConfigNetflowMonitorPolChildrenDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_netflow_monitor_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_monitor"
  annotations = [
    {
      key = "key_0"
      value = "value_1"
    },
    {
      key = "key_1"
      value = "test_value"
    },
  ]
  relation_to_netflow_exporters = [
    {
      annotation = "annotation_1"
      annotations = [
	    {
          key = "key_0"
          value = "value_1"
	    },
	    {
          key = "key_1"
          value = "test_value"
	    },
      ]
      tags = [
	    {
          key = "key_0"
          value = "value_1"
	    },
	    {
          key = "key_1"
          value = "test_value"
	    },
      ]
      netflow_exporter_policy_name = "netflow_exporter_policy_name_0"
    },
    {
      annotation = "annotation_2"
      annotations = [
	    {
          key = "key_0"
          value = "value_1"
	    },
	    {
          key = "key_1"
          value = "test_value"
	    },
      ]
      tags = [
	    {
          key = "key_0"
          value = "value_1"
	    },
	    {
          key = "key_1"
          value = "test_value"
	    },
      ]
      netflow_exporter_policy_name = "netflow_exporter_policy_name_1"
    },
  ]
  relation_to_netflow_record = {
    annotation = "annotation_1"
    annotations = [
	  {
        key = "key_0"
        value = "value_1"
	  },
	  {
        key = "key_1"
        value = "test_value"
	  },
    ]
    tags = [
	  {
        key = "key_0"
        value = "value_1"
	  },
	  {
        key = "key_1"
        value = "test_value"
	  },
    ]
    netflow_record_policy_name = "netflow_record_policy_name_1"
  }
  tags = [
    {
      key = "key_0"
      value = "value_1"
    },
    {
      key = "key_1"
      value = "test_value"
    },
  ]
}
`

const testConfigNetflowMonitorPolChildrenRemoveFromConfigDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_netflow_monitor_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_monitor"
}
`

const testConfigNetflowMonitorPolChildrenRemoveOneDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_netflow_monitor_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_monitor"
  annotations = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
  relation_to_netflow_exporters = [ 
	{
	  annotation = "annotation_2"
      annotations = [ 
	    {
          key = "key_1"
          value = "test_value"
	    },
      ]
      tags = [ 
	    {
          key = "key_1"
          value = "test_value"
	    },
      ]
	  netflow_exporter_policy_name = "netflow_exporter_policy_name_1"
	},
  ]
  relation_to_netflow_record = {
    annotation = "annotation_1"
    annotations = [ 
	  {
        key = "key_1"
        value = "test_value"
	  },
    ]
    tags = [ 
	  {
        key = "key_1"
        value = "test_value"
	  },
    ]
    netflow_record_policy_name = "netflow_record_policy_name_1"
  }
  tags = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
}
`

const testConfigNetflowMonitorPolChildrenRemoveAllDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_netflow_monitor_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "netfow_monitor"
  annotations = []
  relation_to_netflow_exporters = []
  relation_to_netflow_record = {
    annotation = "annotation_1"
    annotations = []
    tags = []
    netflow_record_policy_name = "netflow_record_policy_name_1"
  }
  tags = []
}
`
