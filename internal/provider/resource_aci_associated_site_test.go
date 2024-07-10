// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvSiteAssociatedWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvSiteAssociatedMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "description", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "site_id", "100"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvSiteAssociatedAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "description", "description"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name", "name"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name_alias", "name_alias"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_key", "owner_key"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_tag", "owner_tag"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "site_id", "100"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvSiteAssociatedMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check:              resource.ComposeAggregateTestCheckFunc(),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvSiteAssociatedResetDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "description", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "site_id", "100"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_associated_site.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "description", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "site_id", "100"),
				),
			},
			// Update with children
			{
				Config:             testConfigFvSiteAssociatedChildrenDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "description", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "site_id", "100"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.1.value", "value_2"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_associated_site.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "description", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "site_id", "100"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.1.value", "value_2"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigFvSiteAssociatedChildrenRemoveFromConfigDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvSiteAssociatedChildrenRemoveOneDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvSiteAssociatedChildrenRemoveAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.#", "0"),
				),
			},
		},
	})
}
func TestAccResourceFvSiteAssociatedWithFvBD(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvSiteAssociatedMinDependencyWithFvBD,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "description", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "site_id", "100"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvSiteAssociatedAllDependencyWithFvBD,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "description", "description"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name", "name"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name_alias", "name_alias"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_key", "owner_key"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_tag", "owner_tag"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "site_id", "100"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvSiteAssociatedMinDependencyWithFvBD,
				ExpectNonEmptyPlan: false,
				Check:              resource.ComposeAggregateTestCheckFunc(),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvSiteAssociatedResetDependencyWithFvBD,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "description", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "site_id", "100"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_associated_site.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "description", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "site_id", "100"),
				),
			},
			// Update with children
			{
				Config:             testConfigFvSiteAssociatedChildrenDependencyWithFvBD,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "description", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "site_id", "100"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.1.value", "value_2"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_associated_site.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "description", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_associated_site.test", "site_id", "100"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.1.value", "value_2"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigFvSiteAssociatedChildrenRemoveFromConfigDependencyWithFvBD,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvSiteAssociatedChildrenRemoveOneDependencyWithFvBD,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvSiteAssociatedChildrenRemoveAllDependencyWithFvBD,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_associated_site.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_associated_site.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigFvSiteAssociatedMinDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvTenant + `
resource "aci_associated_site" "test" {
  parent_dn = aci_application_epg.test.id
}
`

const testConfigFvSiteAssociatedAllDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvTenant + `
resource "aci_associated_site" "test" {
  parent_dn = aci_application_epg.test.id
  annotation = "annotation"
  description = "description"
  name = "name"
  name_alias = "name_alias"
  owner_key = "owner_key"
  owner_tag = "owner_tag"
  site_id = "100"
}
`

const testConfigFvSiteAssociatedResetDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvTenant + `
resource "aci_associated_site" "test" {
  parent_dn = aci_application_epg.test.id
  annotation = "orchestrator:terraform"
  description = ""
  name = ""
  name_alias = ""
  owner_key = ""
  owner_tag = ""
  site_id = "100"
}
`
const testConfigFvSiteAssociatedChildrenDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvTenant + `
resource "aci_associated_site" "test" {
  parent_dn = aci_application_epg.test.id
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

const testConfigFvSiteAssociatedChildrenRemoveFromConfigDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvTenant + `
resource "aci_associated_site" "test" {
  parent_dn = aci_application_epg.test.id
}
`

const testConfigFvSiteAssociatedChildrenRemoveOneDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvTenant + `
resource "aci_associated_site" "test" {
  parent_dn = aci_application_epg.test.id
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

const testConfigFvSiteAssociatedChildrenRemoveAllDependencyWithFvAEPg = testConfigFvAEPgMinDependencyWithFvTenant + `
resource "aci_associated_site" "test" {
  parent_dn = aci_application_epg.test.id
  annotations = []
  tags = []
}
`

const testConfigFvSiteAssociatedMinDependencyWithFvBD = testConfigFvBDMinDependencyWithFvAp + `
resource "aci_associated_site" "test" {
  parent_dn = aci_bridge_domain.test.id
}
`

const testConfigFvSiteAssociatedAllDependencyWithFvBD = testConfigFvBDMinDependencyWithFvAp + `
resource "aci_associated_site" "test" {
  parent_dn = aci_bridge_domain.test.id
  annotation = "annotation"
  description = "description"
  name = "name"
  name_alias = "name_alias"
  owner_key = "owner_key"
  owner_tag = "owner_tag"
  site_id = "100"
}
`

const testConfigFvSiteAssociatedResetDependencyWithFvBD = testConfigFvBDMinDependencyWithFvAp + `
resource "aci_associated_site" "test" {
  parent_dn = aci_bridge_domain.test.id
  annotation = "orchestrator:terraform"
  description = ""
  name = ""
  name_alias = ""
  owner_key = ""
  owner_tag = ""
  site_id = "100"
}
`
const testConfigFvSiteAssociatedChildrenDependencyWithFvBD = testConfigFvBDMinDependencyWithFvAp + `
resource "aci_associated_site" "test" {
  parent_dn = aci_bridge_domain.test.id
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

const testConfigFvSiteAssociatedChildrenRemoveFromConfigDependencyWithFvBD = testConfigFvBDMinDependencyWithFvAp + `
resource "aci_associated_site" "test" {
  parent_dn = aci_bridge_domain.test.id
}
`

const testConfigFvSiteAssociatedChildrenRemoveOneDependencyWithFvBD = testConfigFvBDMinDependencyWithFvAp + `
resource "aci_associated_site" "test" {
  parent_dn = aci_bridge_domain.test.id
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

const testConfigFvSiteAssociatedChildrenRemoveAllDependencyWithFvBD = testConfigFvBDMinDependencyWithFvAp + `
resource "aci_associated_site" "test" {
  parent_dn = aci_bridge_domain.test.id
  annotations = []
  tags = []
}
`
