// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvRsIntraEpgWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.0(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsIntraEpgMinDependencyWithFvAEPgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.allow_test", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.allow_test_2", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.allow_test_2", "annotation", "orchestrator:terraform"),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.0(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigFvRsIntraEpgMinDependencyWithFvAEPgAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.0(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsIntraEpgMinDependencyWithFvAEPgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.allow_test", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.allow_test_2", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.allow_test_2", "annotation", "orchestrator:terraform"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.0(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsIntraEpgMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvRsIntraEpgAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotation", "annotation"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvRsIntraEpgMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "contract_name", "test_tn_vz_br_cp_name"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvRsIntraEpgResetDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_relation_to_intra_epg_contract.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigFvRsIntraEpgChildrenDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.1.value", "test_value"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_relation_to_intra_epg_contract.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigFvRsIntraEpgChildrenRemoveFromConfigDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvRsIntraEpgChildrenRemoveOneDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvRsIntraEpgChildrenRemoveAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.#", "0"),
				),
			},
		},
		CheckDestroy: testCheckResourceDestroy,
	})
}
func TestAccResourceFvRsIntraEpgWithFvESg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.0(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsIntraEpgMinDependencyWithFvESgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.allow_test", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.allow_test_2", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.allow_test_2", "annotation", "orchestrator:terraform"),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.0(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigFvRsIntraEpgMinDependencyWithFvESgAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.0(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsIntraEpgMinDependencyWithFvESgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.allow_test", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.allow_test_2", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.allow_test_2", "annotation", "orchestrator:terraform"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.0(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsIntraEpgMinDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvRsIntraEpgAllDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotation", "annotation"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvRsIntraEpgMinDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "contract_name", "test_tn_vz_br_cp_name"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvRsIntraEpgResetDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_relation_to_intra_epg_contract.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigFvRsIntraEpgChildrenDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "contract_name", "test_tn_vz_br_cp_name"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.1.value", "test_value"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_relation_to_intra_epg_contract.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigFvRsIntraEpgChildrenRemoveFromConfigDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvRsIntraEpgChildrenRemoveOneDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvRsIntraEpgChildrenRemoveAllDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_relation_to_intra_epg_contract.test", "tags.#", "0"),
				),
			},
		},
		CheckDestroy: testCheckResourceDestroy,
	})
}

const testConfigFvRsIntraEpgMinDependencyWithFvAEPgAllowExisting = testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_intra_epg_contract" "allow_test" {
  parent_dn = aci_application_epg.test.id
  contract_name = "test_tn_vz_br_cp_name"
}
resource "aci_relation_to_intra_epg_contract" "allow_test_2" {
  parent_dn = aci_application_epg.test.id
  contract_name = "test_tn_vz_br_cp_name"
  depends_on = [aci_relation_to_intra_epg_contract.allow_test]
}
`

const testConfigFvRsIntraEpgMinDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_intra_epg_contract" "test" {
  parent_dn = aci_application_epg.test.id
  contract_name = "test_tn_vz_br_cp_name"
}
`

const testConfigFvRsIntraEpgAllDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_intra_epg_contract" "test" {
  parent_dn = aci_application_epg.test.id
  contract_name = "test_tn_vz_br_cp_name"
  annotation = "annotation"
}
`

const testConfigFvRsIntraEpgResetDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_intra_epg_contract" "test" {
  parent_dn = aci_application_epg.test.id
  contract_name = "test_tn_vz_br_cp_name"
  annotation = "orchestrator:terraform"
}
`
const testConfigFvRsIntraEpgChildrenDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_intra_epg_contract" "test" {
  parent_dn = aci_application_epg.test.id
  contract_name = "test_tn_vz_br_cp_name"
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

const testConfigFvRsIntraEpgChildrenRemoveFromConfigDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_intra_epg_contract" "test" {
  parent_dn = aci_application_epg.test.id
  contract_name = "test_tn_vz_br_cp_name"
}
`

const testConfigFvRsIntraEpgChildrenRemoveOneDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_intra_epg_contract" "test" {
  parent_dn = aci_application_epg.test.id
  contract_name = "test_tn_vz_br_cp_name"
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

const testConfigFvRsIntraEpgChildrenRemoveAllDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_intra_epg_contract" "test" {
  parent_dn = aci_application_epg.test.id
  contract_name = "test_tn_vz_br_cp_name"
  annotations = []
  tags = []
}
`

const testConfigFvRsIntraEpgMinDependencyWithFvESgAllowExisting = testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_intra_epg_contract" "allow_test" {
  parent_dn = aci_endpoint_security_group.test.id
  contract_name = "test_tn_vz_br_cp_name"
}
resource "aci_relation_to_intra_epg_contract" "allow_test_2" {
  parent_dn = aci_endpoint_security_group.test.id
  contract_name = "test_tn_vz_br_cp_name"
  depends_on = [aci_relation_to_intra_epg_contract.allow_test]
}
`

const testConfigFvRsIntraEpgMinDependencyWithFvESg = testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_intra_epg_contract" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  contract_name = "test_tn_vz_br_cp_name"
}
`

const testConfigFvRsIntraEpgAllDependencyWithFvESg = testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_intra_epg_contract" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  contract_name = "test_tn_vz_br_cp_name"
  annotation = "annotation"
}
`

const testConfigFvRsIntraEpgResetDependencyWithFvESg = testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_intra_epg_contract" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  contract_name = "test_tn_vz_br_cp_name"
  annotation = "orchestrator:terraform"
}
`
const testConfigFvRsIntraEpgChildrenDependencyWithFvESg = testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_intra_epg_contract" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  contract_name = "test_tn_vz_br_cp_name"
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

const testConfigFvRsIntraEpgChildrenRemoveFromConfigDependencyWithFvESg = testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_intra_epg_contract" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  contract_name = "test_tn_vz_br_cp_name"
}
`

const testConfigFvRsIntraEpgChildrenRemoveOneDependencyWithFvESg = testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_intra_epg_contract" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  contract_name = "test_tn_vz_br_cp_name"
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

const testConfigFvRsIntraEpgChildrenRemoveAllDependencyWithFvESg = testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_intra_epg_contract" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  contract_name = "test_tn_vz_br_cp_name"
  annotations = []
  tags = []
}
`
