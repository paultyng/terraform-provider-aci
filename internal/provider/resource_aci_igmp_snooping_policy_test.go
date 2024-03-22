// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceIgmpSnoopPolWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigIgmpSnoopPolMinDependencyWithFvTenantAllowExisting + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "control.#", "0"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "control.#", "0"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "last_member_interval", "1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "last_member_interval", "1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "query_interval", "125"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "query_interval", "125"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "response_interval", "10"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "response_interval", "10"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "start_query_count", "2"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "start_query_count", "2"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "start_query_interval", "31"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "start_query_interval", "31"),
					composeAggregateTestCheckFuncWithVersion(t, "5.1(1h)", ">=",
						resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "querier_version", "v3"),
						resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "querier_version", "v3")),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigIgmpSnoopPolMinDependencyWithFvTenantAllowExisting + testConfigDataSourceSystem,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigIgmpSnoopPolMinDependencyWithFvTenantAllowExisting + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "control.#", "0"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "control.#", "0"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "last_member_interval", "1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "last_member_interval", "1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "query_interval", "125"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "query_interval", "125"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "response_interval", "10"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "response_interval", "10"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "start_query_count", "2"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "start_query_count", "2"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "start_query_interval", "31"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "start_query_interval", "31"),
					composeAggregateTestCheckFuncWithVersion(t, "5.1(1h)", ">=",
						resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test", "querier_version", "v3"),
						resource.TestCheckResourceAttr("aci_igmp_snooping_policy.allow_test_2", "querier_version", "v3")),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigIgmpSnoopPolMinDependencyWithFvTenant + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "control.#", "0"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "last_member_interval", "1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "query_interval", "125"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "response_interval", "10"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "start_query_count", "2"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "start_query_interval", "31"),
					composeAggregateTestCheckFuncWithVersion(t, "5.1(1h)", ">=",
						resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "querier_version", "v3")),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigIgmpSnoopPolAllDependencyWithFvTenant + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "admin_state", "disabled"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "control.#", "2"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "control.0", "fast-leave"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "control.1", "querier"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "last_member_interval", "3"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "owner_key", "owner_key_1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "owner_tag", "owner_tag_1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "query_interval", "140"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "response_interval", "11"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "start_query_count", "9"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "start_query_interval", "2"),
					composeAggregateTestCheckFuncWithVersion(t, "5.1(1h)", ">=",
						resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "querier_version", "v2")),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigIgmpSnoopPolMinDependencyWithFvTenant + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "name", "test_name"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigIgmpSnoopPolResetDependencyWithFvTenant + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "control.#", "0"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "last_member_interval", "1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "query_interval", "125"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "response_interval", "10"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "start_query_count", "2"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "start_query_interval", "31"),
					composeAggregateTestCheckFuncWithVersion(t, "5.1(1h)", ">=",
						resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "querier_version", "v3")),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_igmp_snooping_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigIgmpSnoopPolChildrenDependencyWithFvTenant + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "control.#", "0"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "last_member_interval", "1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "query_interval", "125"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "response_interval", "10"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "start_query_count", "2"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "start_query_interval", "31"),
					composeAggregateTestCheckFuncWithVersion(t, "5.1(1h)", ">=",
						resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "querier_version", "v3")),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "tags.1.value", "test_value"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_igmp_snooping_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigIgmpSnoopPolChildrenRemoveFromConfigDependencyWithFvTenant + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigIgmpSnoopPolChildrenRemoveOneDependencyWithFvTenant + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigIgmpSnoopPolChildrenRemoveAllDependencyWithFvTenant + testConfigDataSourceSystem,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_igmp_snooping_policy.test", "tags.#", "0"),
				),
			},
		},
		CheckDestroy: testCheckResourceDestroy,
	})
}

const testConfigIgmpSnoopPolMinDependencyWithFvTenantAllowExisting = testConfigFvTenantMin + `
resource "aci_igmp_snooping_policy" "allow_test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
}
resource "aci_igmp_snooping_policy" "allow_test_2" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  depends_on = [aci_igmp_snooping_policy.allow_test]
}
`

const testConfigIgmpSnoopPolMinDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_igmp_snooping_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
}
`

const testConfigIgmpSnoopPolAllDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_igmp_snooping_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  admin_state = "disabled"
  annotation = "annotation"
  control = ["fast-leave", "querier"]
  description = "description_1"
  last_member_interval = "3"
  name_alias = "name_alias_1"
  owner_key = "owner_key_1"
  owner_tag = "owner_tag_1"
  query_interval = "140"
  response_interval = "11"
  start_query_count = "9"
  start_query_interval = "2"
  querier_version = provider::aci::compare_versions(data.aci_system.version.version,">=","5.1(1h)") ? "v2" : null
}
`

const testConfigIgmpSnoopPolResetDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_igmp_snooping_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  admin_state = "enabled"
  annotation = "orchestrator:terraform"
  control = []
  description = ""
  last_member_interval = "1"
  name_alias = ""
  owner_key = ""
  owner_tag = ""
  query_interval = "125"
  response_interval = "10"
  start_query_count = "2"
  start_query_interval = "31"
  querier_version = provider::aci::compare_versions(data.aci_system.version.version,">=","5.1(1h)") ? "v3" : null
}
`
const testConfigIgmpSnoopPolChildrenDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_igmp_snooping_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
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
}
`

const testConfigIgmpSnoopPolChildrenRemoveFromConfigDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_igmp_snooping_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
}
`

const testConfigIgmpSnoopPolChildrenRemoveOneDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_igmp_snooping_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
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
}
`

const testConfigIgmpSnoopPolChildrenRemoveAllDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_igmp_snooping_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  annotations = []
  tags = []
}
`
