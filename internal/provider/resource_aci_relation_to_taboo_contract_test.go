// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvRsProtByWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsProtByMinDependencyWithFvAEPgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "taboo_contract_name", "test_tn_vz_taboo_name"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test_2", "taboo_contract_name", "test_tn_vz_taboo_name"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test_2", "annotation", "orchestrator:terraform"),
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
				Config:      testConfigFvRsProtByMinDependencyWithFvAEPgAllowExisting,
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
				Config:             testConfigFvRsProtByMinDependencyWithFvAEPgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "taboo_contract_name", "test_tn_vz_taboo_name"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test_2", "taboo_contract_name", "test_tn_vz_taboo_name"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test_2", "annotation", "orchestrator:terraform"),
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
				Config:             testConfigFvRsProtByMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "taboo_contract_name", "test_tn_vz_taboo_name"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvRsProtByAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "taboo_contract_name", "test_tn_vz_taboo_name"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotation", "annotation"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvRsProtByMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "taboo_contract_name", "test_tn_vz_taboo_name"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvRsProtByResetDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "taboo_contract_name", "test_tn_vz_taboo_name"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_relation_to_taboo_contract.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigFvRsProtByChildrenDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "taboo_contract_name", "test_tn_vz_taboo_name"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "tags.1.value", "value_2"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_relation_to_taboo_contract.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigFvRsProtByChildrenRemoveFromConfigDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvRsProtByChildrenRemoveOneDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "tags.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvRsProtByChildrenRemoveAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_relation_to_taboo_contract.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigFvRsProtByMinDependencyWithFvAEPgAllowExisting = testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_taboo_contract" "test" {
  parent_dn = aci_application_epg.test.id
  taboo_contract_name = "test_tn_vz_taboo_name"
}
resource "aci_relation_to_taboo_contract" "test_2" {
  parent_dn = aci_application_epg.test.id
  taboo_contract_name = "test_tn_vz_taboo_name"
  depends_on = [aci_relation_to_taboo_contract.test]
}
`

const testConfigFvRsProtByMinDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_taboo_contract" "test" {
  parent_dn = aci_application_epg.test.id
  taboo_contract_name = "test_tn_vz_taboo_name"
}
`

const testConfigFvRsProtByAllDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_taboo_contract" "test" {
  parent_dn = aci_application_epg.test.id
  taboo_contract_name = "test_tn_vz_taboo_name"
  annotation = "annotation"
}
`

const testConfigFvRsProtByResetDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_taboo_contract" "test" {
  parent_dn = aci_application_epg.test.id
  taboo_contract_name = "test_tn_vz_taboo_name"
  annotation = "orchestrator:terraform"
}
`
const testConfigFvRsProtByChildrenDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_taboo_contract" "test" {
  parent_dn = aci_application_epg.test.id
  taboo_contract_name = "test_tn_vz_taboo_name"
  annotations = [
	{
	  key = "key_0"
	  value = "value_1"
	},
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
  tags = [
	{
	  key = "key_0"
	  value = "value_1"
	},
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
}
`

const testConfigFvRsProtByChildrenRemoveFromConfigDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_taboo_contract" "test" {
  parent_dn = aci_application_epg.test.id
  taboo_contract_name = "test_tn_vz_taboo_name"
}
`

const testConfigFvRsProtByChildrenRemoveOneDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_taboo_contract" "test" {
  parent_dn = aci_application_epg.test.id
  taboo_contract_name = "test_tn_vz_taboo_name"
  annotations = [ 
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
  tags = [ 
	{
	  key = "key_1"
	  value = "value_2"
	},
  ]
}
`

const testConfigFvRsProtByChildrenRemoveAllDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_taboo_contract" "test" {
  parent_dn = aci_application_epg.test.id
  taboo_contract_name = "test_tn_vz_taboo_name"
  annotations = []
  tags = []
}
`
