// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceL3extProvLblWithL3extOut(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigL3extProvLblMinDependencyWithL3extOutAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test", "name", "prov_label"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test_2", "name", "prov_label"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test_2", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test", "tag", "yellow-green"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test_2", "tag", "yellow-green"),
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
				Config:      testConfigL3extProvLblMinDependencyWithL3extOutAllowExisting,
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
				Config:             testConfigL3extProvLblMinDependencyWithL3extOutAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test", "name", "prov_label"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test_2", "name", "prov_label"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test_2", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test", "tag", "yellow-green"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.allow_test_2", "tag", "yellow-green"),
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
				Config:             testConfigL3extProvLblMinDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "name", "prov_label"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tag", "yellow-green"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigL3extProvLblAllDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "name", "prov_label"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "owner_key", "owner_key_1"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "owner_tag", "owner_tag_1"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tag", "alice-blue"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigL3extProvLblMinDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "name", "prov_label"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigL3extProvLblResetDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "name", "prov_label"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tag", "yellow-green"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_l3out_provider_label.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigL3extProvLblChildrenDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "name", "prov_label"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "description", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tag", "yellow-green"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tags.1.value", "test_value"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_l3out_provider_label.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigL3extProvLblChildrenRemoveFromConfigDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigL3extProvLblChildrenRemoveOneDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigL3extProvLblChildrenRemoveAllDependencyWithL3extOut,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_l3out_provider_label.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigL3extProvLblMinDependencyWithL3extOutAllowExisting = testConfigL3extOutMinDependencyWithFvTenantInfra + `
resource "aci_l3out_provider_label" "allow_test" {
  parent_dn = aci_l3_outside.test.id
  name = "prov_label"
}
resource "aci_l3out_provider_label" "allow_test_2" {
  parent_dn = aci_l3_outside.test.id
  name = "prov_label"
  depends_on = [aci_l3out_provider_label.allow_test]
}
`

const testConfigL3extProvLblMinDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenantInfra + `
resource "aci_l3out_provider_label" "test" {
  parent_dn = aci_l3_outside.test.id
  name = "prov_label"
}
`

const testConfigL3extProvLblAllDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenantInfra + `
resource "aci_l3out_provider_label" "test" {
  parent_dn = aci_l3_outside.test.id
  name = "prov_label"
  annotation = "annotation"
  description = "description_1"
  name_alias = "name_alias_1"
  owner_key = "owner_key_1"
  owner_tag = "owner_tag_1"
  tag = "alice-blue"
}
`

const testConfigL3extProvLblResetDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenantInfra + `
resource "aci_l3out_provider_label" "test" {
  parent_dn = aci_l3_outside.test.id
  name = "prov_label"
  annotation = "orchestrator:terraform"
  description = ""
  name_alias = ""
  owner_key = ""
  owner_tag = ""
  tag = "yellow-green"
}
`
const testConfigL3extProvLblChildrenDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenantInfra + `
resource "aci_l3out_provider_label" "test" {
  parent_dn = aci_l3_outside.test.id
  name = "prov_label"
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

const testConfigL3extProvLblChildrenRemoveFromConfigDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenantInfra + `
resource "aci_l3out_provider_label" "test" {
  parent_dn = aci_l3_outside.test.id
  name = "prov_label"
}
`

const testConfigL3extProvLblChildrenRemoveOneDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenantInfra + `
resource "aci_l3out_provider_label" "test" {
  parent_dn = aci_l3_outside.test.id
  name = "prov_label"
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

const testConfigL3extProvLblChildrenRemoveAllDependencyWithL3extOut = testConfigL3extOutMinDependencyWithFvTenantInfra + `
resource "aci_l3out_provider_label" "test" {
  parent_dn = aci_l3_outside.test.id
  name = "prov_label"
  annotations = []
  tags = []
}
`
