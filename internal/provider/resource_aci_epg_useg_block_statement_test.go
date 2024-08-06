// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvCrtrnWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvCrtrnMinDependencyWithFvAEPgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "match", "any"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "match", "any"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "name", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "name", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "precedence", "0"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "precedence", "0"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "scope", "scope-bd"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "scope", "scope-bd"),
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
				Config:      testConfigFvCrtrnMinDependencyWithFvAEPgAllowExisting,
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
				Config:             testConfigFvCrtrnMinDependencyWithFvAEPgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "match", "any"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "match", "any"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "name", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "name", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "precedence", "0"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "precedence", "0"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test", "scope", "scope-bd"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.allow_test_2", "scope", "scope-bd"),
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
				Config:             testConfigFvCrtrnMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "match", "any"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "name", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "precedence", "0"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "scope", "scope-bd"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvCrtrnAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "match", "all"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "name", "criterion"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "owner_key", "owner_key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "owner_tag", "owner_tag_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "precedence", "1"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "scope", "scope-bd"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvCrtrnMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check:              resource.ComposeAggregateTestCheckFunc(),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvCrtrnResetDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "match", "any"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "name", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "precedence", "0"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "scope", "scope-bd"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_epg_useg_block_statement.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigFvCrtrnChildrenDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "match", "any"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "name", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "precedence", "0"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "scope", "scope-bd"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "tags.1.value", "test_value"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_epg_useg_block_statement.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigFvCrtrnChildrenRemoveFromConfigDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvCrtrnChildrenRemoveOneDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvCrtrnChildrenRemoveAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_epg_useg_block_statement.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigFvCrtrnMinDependencyWithFvAEPgAllowExisting = testConfigFvAEPgMin + `
resource "aci_epg_useg_block_statement" "allow_test" {
  parent_dn = aci_application_epg.test.id
}
resource "aci_epg_useg_block_statement" "allow_test_2" {
  parent_dn = aci_application_epg.test.id
  depends_on = [aci_epg_useg_block_statement.allow_test]
}
`

const testConfigFvCrtrnMinDependencyWithFvAEPg = testConfigFvAEPgMin + `
resource "aci_epg_useg_block_statement" "test" {
  parent_dn = aci_application_epg.test.id
}
`

const testConfigFvCrtrnAllDependencyWithFvAEPg = testConfigFvAEPgMin + `
resource "aci_epg_useg_block_statement" "test" {
  parent_dn = aci_application_epg.test.id
  annotation = "annotation"
  description = "description_1"
  match = "all"
  name = "criterion"
  name_alias = "name_alias_1"
  owner_key = "owner_key_1"
  owner_tag = "owner_tag_1"
  precedence = "1"
  scope = "scope-bd"
}
`

const testConfigFvCrtrnResetDependencyWithFvAEPg = testConfigFvAEPgMin + `
resource "aci_epg_useg_block_statement" "test" {
  parent_dn = aci_application_epg.test.id
  annotation = "orchestrator:terraform"
  description = ""
  match = "any"
  name = ""
  name_alias = ""
  owner_key = ""
  owner_tag = ""
  precedence = "0"
  scope = "scope-bd"
}
`
const testConfigFvCrtrnChildrenDependencyWithFvAEPg = testConfigFvAEPgMin + `
resource "aci_epg_useg_block_statement" "test" {
  parent_dn = aci_application_epg.test.id
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

const testConfigFvCrtrnChildrenRemoveFromConfigDependencyWithFvAEPg = testConfigFvAEPgMin + `
resource "aci_epg_useg_block_statement" "test" {
  parent_dn = aci_application_epg.test.id
}
`

const testConfigFvCrtrnChildrenRemoveOneDependencyWithFvAEPg = testConfigFvAEPgMin + `
resource "aci_epg_useg_block_statement" "test" {
  parent_dn = aci_application_epg.test.id
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

const testConfigFvCrtrnChildrenRemoveAllDependencyWithFvAEPg = testConfigFvAEPgMin + `
resource "aci_epg_useg_block_statement" "test" {
  parent_dn = aci_application_epg.test.id
  annotations = []
  tags = []
}
`
