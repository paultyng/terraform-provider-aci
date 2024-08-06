// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvFBRouteWithFvFBRGroup(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvFBRouteMinDependencyWithFvFBRGroupAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test", "prefix_address", "2.2.2.3/24"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test_2", "prefix_address", "2.2.2.3/24"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test", "name", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test_2", "name", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test_2", "name_alias", ""),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigFvFBRouteMinDependencyWithFvFBRGroupAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvFBRouteMinDependencyWithFvFBRGroupAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test", "prefix_address", "2.2.2.3/24"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test_2", "prefix_address", "2.2.2.3/24"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test", "name", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test_2", "name", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.allow_test_2", "name_alias", ""),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvFBRouteMinDependencyWithFvFBRGroup,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "prefix_address", "2.2.2.3/24"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "name", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "name_alias", ""),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvFBRouteAllDependencyWithFvFBRGroup,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "prefix_address", "2.2.2.3/24"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "name", "name_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "name_alias", "name_alias_1"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvFBRouteMinDependencyWithFvFBRGroup,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "prefix_address", "2.2.2.3/24"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvFBRouteResetDependencyWithFvFBRGroup,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "prefix_address", "2.2.2.3/24"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "name", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "name_alias", ""),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_vrf_fallback_route.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigFvFBRouteChildrenDependencyWithFvFBRGroup,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "prefix_address", "2.2.2.3/24"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "description", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "name", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "tags.1.value", "test_value"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_vrf_fallback_route.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigFvFBRouteChildrenRemoveFromConfigDependencyWithFvFBRGroup,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvFBRouteChildrenRemoveOneDependencyWithFvFBRGroup,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvFBRouteChildrenRemoveAllDependencyWithFvFBRGroup,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_vrf_fallback_route.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigFvFBRouteMinDependencyWithFvFBRGroupAllowExisting = testConfigFvFBRGroupMinDependencyWithFvCtx + `
resource "aci_vrf_fallback_route" "allow_test" {
  parent_dn = aci_vrf_fallback_route_group.test.id
  prefix_address = "2.2.2.3/24"
}
resource "aci_vrf_fallback_route" "allow_test_2" {
  parent_dn = aci_vrf_fallback_route_group.test.id
  prefix_address = "2.2.2.3/24"
  depends_on = [aci_vrf_fallback_route.allow_test]
}
`

const testConfigFvFBRouteMinDependencyWithFvFBRGroup = testConfigFvFBRGroupMinDependencyWithFvCtx + `
resource "aci_vrf_fallback_route" "test" {
  parent_dn = aci_vrf_fallback_route_group.test.id
  prefix_address = "2.2.2.3/24"
}
`

const testConfigFvFBRouteAllDependencyWithFvFBRGroup = testConfigFvFBRGroupMinDependencyWithFvCtx + `
resource "aci_vrf_fallback_route" "test" {
  parent_dn = aci_vrf_fallback_route_group.test.id
  prefix_address = "2.2.2.3/24"
  annotation = "annotation"
  description = "description_1"
  name = "name_1"
  name_alias = "name_alias_1"
}
`

const testConfigFvFBRouteResetDependencyWithFvFBRGroup = testConfigFvFBRGroupMinDependencyWithFvCtx + `
resource "aci_vrf_fallback_route" "test" {
  parent_dn = aci_vrf_fallback_route_group.test.id
  prefix_address = "2.2.2.3/24"
  annotation = "orchestrator:terraform"
  description = ""
  name = ""
  name_alias = ""
}
`
const testConfigFvFBRouteChildrenDependencyWithFvFBRGroup = testConfigFvFBRGroupMinDependencyWithFvCtx + `
resource "aci_vrf_fallback_route" "test" {
  parent_dn = aci_vrf_fallback_route_group.test.id
  prefix_address = "2.2.2.3/24"
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

const testConfigFvFBRouteChildrenRemoveFromConfigDependencyWithFvFBRGroup = testConfigFvFBRGroupMinDependencyWithFvCtx + `
resource "aci_vrf_fallback_route" "test" {
  parent_dn = aci_vrf_fallback_route_group.test.id
  prefix_address = "2.2.2.3/24"
}
`

const testConfigFvFBRouteChildrenRemoveOneDependencyWithFvFBRGroup = testConfigFvFBRGroupMinDependencyWithFvCtx + `
resource "aci_vrf_fallback_route" "test" {
  parent_dn = aci_vrf_fallback_route_group.test.id
  prefix_address = "2.2.2.3/24"
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

const testConfigFvFBRouteChildrenRemoveAllDependencyWithFvFBRGroup = testConfigFvFBRGroupMinDependencyWithFvCtx + `
resource "aci_vrf_fallback_route" "test" {
  parent_dn = aci_vrf_fallback_route_group.test.id
  prefix_address = "2.2.2.3/24"
  annotations = []
  tags = []
}
`
