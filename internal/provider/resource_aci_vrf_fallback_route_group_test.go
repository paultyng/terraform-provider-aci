// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccResourceFvFBRGroupWithFvCtx(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "5.2(4d)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvFBRGroupMinDependencyWithFvCtxAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.allow_test", "name", "fallback_route_group"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.allow_test_2", "name", "fallback_route_group"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.allow_test_2", "name_alias", ""),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "5.2(4d)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigFvFBRGroupMinDependencyWithFvCtxAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "5.2(4d)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvFBRGroupMinDependencyWithFvCtxAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.allow_test", "name", "fallback_route_group"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.allow_test_2", "name", "fallback_route_group"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.allow_test_2", "name_alias", ""),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "5.2(4d)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvFBRGroupMinDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name", "fallback_route_group"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name_alias", ""),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvFBRGroupAllDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name", "fallback_route_group"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name_alias", "name_alias_1"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvFBRGroupMinDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name", "fallback_route_group"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvFBRGroupResetDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name", "fallback_route_group"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name_alias", ""),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_vrf_fallback_route_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigFvFBRGroupChildrenDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name", "fallback_route_group"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.description", "description_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.name", "name_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.prefix_address", "2.2.2.2/24"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.description", "description_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.fallback_member", "2.2.2.2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.name", "name_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.description", "description_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.fallback_member", "2.2.2.3"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.name", "name_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.name_alias", "name_alias_2"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_vrf_fallback_route_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigFvFBRGroupChildrenRemoveFromConfigDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.#", "2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.#", "2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.description", "description_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.name", "name_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route.prefix_address", "2.2.2.2/24"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.#", "2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.description", "description_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.fallback_member", "2.2.2.2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.name", "name_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.#", "2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.description", "description_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.fallback_member", "2.2.2.3"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.name", "name_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.1.name_alias", "name_alias_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvFBRGroupChildrenRemoveOneDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.tags.#", "1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.description", "description_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.fallback_member", "2.2.2.3"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.name", "name_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.0.name_alias", "name_alias_2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.#", "1"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("aci_vrf_fallback_route_group.test",
						tfjsonpath.New("vrf_fallback_route"),
						knownvalue.MapExact(
							map[string]knownvalue.Check{
								"annotation":     knownvalue.Null(),
								"annotations":    knownvalue.Null(),
								"tags":           knownvalue.Null(),
								"description":    knownvalue.Null(),
								"name":           knownvalue.Null(),
								"name_alias":     knownvalue.Null(),
								"prefix_address": knownvalue.Null(),
							},
						),
					),
				},
			},
			// Update with all children removed
			{
				Config:             testConfigFvFBRGroupChildrenRemoveAllDependencyWithFvCtx,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "tags.#", "0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route_group.test", "vrf_fallback_route_group_members.#", "0"),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("aci_vrf_fallback_route_group.test",
						tfjsonpath.New("vrf_fallback_route"),
						knownvalue.MapExact(
							map[string]knownvalue.Check{
								"annotation":     knownvalue.Null(),
								"annotations":    knownvalue.Null(),
								"tags":           knownvalue.Null(),
								"description":    knownvalue.Null(),
								"name":           knownvalue.Null(),
								"name_alias":     knownvalue.Null(),
								"prefix_address": knownvalue.Null(),
							},
						),
					),
				},
			},
		},
		CheckDestroy: testCheckResourceDestroy,
	})
}

const testConfigFvFBRGroupMinDependencyWithFvCtxAllowExisting = testConfigFvCtxMinDependencyWithFvTenant + `
resource "aci_vrf_fallback_route_group" "allow_test" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
}
resource "aci_vrf_fallback_route_group" "allow_test_2" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
  depends_on = [aci_vrf_fallback_route_group.allow_test]
}
`

const testConfigFvFBRGroupMinDependencyWithFvCtx = testConfigFvCtxMinDependencyWithFvTenant + `
resource "aci_vrf_fallback_route_group" "test" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
}
`

const testConfigFvFBRGroupAllDependencyWithFvCtx = testConfigFvCtxMinDependencyWithFvTenant + `
resource "aci_vrf_fallback_route_group" "test" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
  annotation = "annotation"
  description = "description_1"
  name_alias = "name_alias_1"
}
`

const testConfigFvFBRGroupResetDependencyWithFvCtx = testConfigFvCtxMinDependencyWithFvTenant + `
resource "aci_vrf_fallback_route_group" "test" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
  annotation = "orchestrator:terraform"
  description = ""
  name_alias = ""
}
`
const testConfigFvFBRGroupChildrenDependencyWithFvCtx = testConfigFvCtxMinDependencyWithFvTenant + `
resource "aci_vrf_fallback_route_group" "test" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
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
  vrf_fallback_route = {
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
    description = "description_1"
    name = "name_1"
    name_alias = "name_alias_1"
    prefix_address = "2.2.2.2/24"
  }
  vrf_fallback_route_group_members = [
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
      description = "description_1"
      fallback_member = "2.2.2.2"
      name = "name_1"
      name_alias = "name_alias_1"
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
      description = "description_2"
      fallback_member = "2.2.2.3"
      name = "name_2"
      name_alias = "name_alias_2"
    },
  ]
}
`

const testConfigFvFBRGroupChildrenRemoveFromConfigDependencyWithFvCtx = testConfigFvCtxMinDependencyWithFvTenant + `
resource "aci_vrf_fallback_route_group" "test" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
}
`

const testConfigFvFBRGroupChildrenRemoveOneDependencyWithFvCtx = testConfigFvCtxMinDependencyWithFvTenant + `
resource "aci_vrf_fallback_route_group" "test" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
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
  vrf_fallback_route = {}
  vrf_fallback_route_group_members = [ 
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
	  description = "description_2"
	  fallback_member = "2.2.2.3"
	  name = "name_2"
	  name_alias = "name_alias_2"
	},
  ]
}
`

const testConfigFvFBRGroupChildrenRemoveAllDependencyWithFvCtx = testConfigFvCtxMinDependencyWithFvTenant + `
resource "aci_vrf_fallback_route_group" "test" {
  parent_dn = aci_vrf.test.id
  name = "fallback_route_group"
  annotations = []
  tags = []
  vrf_fallback_route = {}
  vrf_fallback_route_group_members = []
}
`
