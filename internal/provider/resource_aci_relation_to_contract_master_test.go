// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvRsSecInheritedWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsSecInheritedMinDependencyWithFvAEPgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.allow_test", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.allow_test_2", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.allow_test_2", "annotation", "orchestrator:terraform"),
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
				Config:      testConfigFvRsSecInheritedMinDependencyWithFvAEPgAllowExisting,
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
				Config:             testConfigFvRsSecInheritedMinDependencyWithFvAEPgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.allow_test", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.allow_test_2", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.allow_test_2", "annotation", "orchestrator:terraform"),
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
				Config:             testConfigFvRsSecInheritedMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvRsSecInheritedAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "annotation"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvRsSecInheritedMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvRsSecInheritedResetDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_relation_to_contract_master.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigFvRsSecInheritedChildrenDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/epg-epg_2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.value", "test_value"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_relation_to_contract_master.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigFvRsSecInheritedChildrenRemoveFromConfigDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvRsSecInheritedChildrenRemoveOneDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvRsSecInheritedChildrenRemoveAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.#", "0"),
				),
			},
		},
	})
}
func TestAccResourceFvRsSecInheritedWithFvESg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsSecInheritedMinDependencyWithFvESgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.allow_test", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.allow_test_2", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.allow_test_2", "annotation", "orchestrator:terraform"),
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
				Config:      testConfigFvRsSecInheritedMinDependencyWithFvESgAllowExisting,
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
				Config:             testConfigFvRsSecInheritedMinDependencyWithFvESgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.allow_test", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.allow_test_2", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.allow_test_2", "annotation", "orchestrator:terraform"),
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
				Config:             testConfigFvRsSecInheritedMinDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvRsSecInheritedAllDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "annotation"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvRsSecInheritedMinDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvRsSecInheritedResetDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_relation_to_contract_master.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigFvRsSecInheritedChildrenDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "target_dn", "uni/tn-test_tenant/ap-test_ap/esg-esg_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.value", "test_value"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_relation_to_contract_master.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigFvRsSecInheritedChildrenRemoveFromConfigDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvRsSecInheritedChildrenRemoveOneDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvRsSecInheritedChildrenRemoveAllDependencyWithFvESg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_relation_to_contract_master.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testDependencyConfigFvRsSecInherited = `
resource "aci_endpoint_security_group" "test_endpoint_security_group_0" {
  application_profile_dn = aci_application_profile.test.id
  name = "esg_0"
}
resource "aci_endpoint_security_group" "test_endpoint_security_group_1" {
  application_profile_dn = aci_application_profile.test.id
  name = "esg_1"
}
resource "aci_application_epg" "test_application_epg_0" {
  application_profile_dn = aci_application_profile.test.id
  name = "epg_2"
}
resource "aci_application_epg" "test_application_epg_1" {
  application_profile_dn = aci_application_profile.test.id
  name = "epg_3"
}
`

const testConfigFvRsSecInheritedMinDependencyWithFvAEPgAllowExisting = testDependencyConfigFvRsSecInherited + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "allow_test" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_application_epg_0.id
}
resource "aci_relation_to_contract_master" "allow_test_2" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_application_epg_0.id
  depends_on = [aci_relation_to_contract_master.allow_test]
}
`

const testConfigFvRsSecInheritedMinDependencyWithFvAEPg = testDependencyConfigFvRsSecInherited + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_application_epg_0.id
}
`

const testConfigFvRsSecInheritedAllDependencyWithFvAEPg = testDependencyConfigFvRsSecInherited + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_application_epg_0.id
  annotation = "annotation"
}
`

const testConfigFvRsSecInheritedResetDependencyWithFvAEPg = testDependencyConfigFvRsSecInherited + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_application_epg_0.id
  annotation = "orchestrator:terraform"
}
`
const testConfigFvRsSecInheritedChildrenDependencyWithFvAEPg = testDependencyConfigFvRsSecInherited + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_application_epg_0.id
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

const testConfigFvRsSecInheritedChildrenRemoveFromConfigDependencyWithFvAEPg = testDependencyConfigFvRsSecInherited + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_application_epg_0.id
}
`

const testConfigFvRsSecInheritedChildrenRemoveOneDependencyWithFvAEPg = testDependencyConfigFvRsSecInherited + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_application_epg_0.id
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

const testConfigFvRsSecInheritedChildrenRemoveAllDependencyWithFvAEPg = testDependencyConfigFvRsSecInherited + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = aci_application_epg.test_application_epg_0.id
  annotations = []
  tags = []
}
`

const testConfigFvRsSecInheritedMinDependencyWithFvESgAllowExisting = testDependencyConfigFvRsSecInherited + testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "allow_test" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_endpoint_security_group_0.id
}
resource "aci_relation_to_contract_master" "allow_test_2" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_endpoint_security_group_0.id
  depends_on = [aci_relation_to_contract_master.allow_test]
}
`

const testConfigFvRsSecInheritedMinDependencyWithFvESg = testDependencyConfigFvRsSecInherited + testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_endpoint_security_group_0.id
}
`

const testConfigFvRsSecInheritedAllDependencyWithFvESg = testDependencyConfigFvRsSecInherited + testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_endpoint_security_group_0.id
  annotation = "annotation"
}
`

const testConfigFvRsSecInheritedResetDependencyWithFvESg = testDependencyConfigFvRsSecInherited + testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_endpoint_security_group_0.id
  annotation = "orchestrator:terraform"
}
`
const testConfigFvRsSecInheritedChildrenDependencyWithFvESg = testDependencyConfigFvRsSecInherited + testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_endpoint_security_group_0.id
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

const testConfigFvRsSecInheritedChildrenRemoveFromConfigDependencyWithFvESg = testDependencyConfigFvRsSecInherited + testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_endpoint_security_group_0.id
}
`

const testConfigFvRsSecInheritedChildrenRemoveOneDependencyWithFvESg = testDependencyConfigFvRsSecInherited + testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_endpoint_security_group_0.id
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

const testConfigFvRsSecInheritedChildrenRemoveAllDependencyWithFvESg = testDependencyConfigFvRsSecInherited + testConfigFvESgMinDependencyWithFvAp + `
resource "aci_relation_to_contract_master" "test" {
  parent_dn = aci_endpoint_security_group.test.id
  target_dn = aci_endpoint_security_group.test_endpoint_security_group_0.id
  annotations = []
  tags = []
}
`
