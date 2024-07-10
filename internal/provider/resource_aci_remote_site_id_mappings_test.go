// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvRemoteIdWithFvSiteAssociated(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRemoteIdMinDependencyWithFvSiteAssociatedAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "description", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "name", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "name", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_pc_tag", "16386"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "remote_pc_tag", "16386"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_vrf_pc_tag", "2818057"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "remote_vrf_pc_tag", "2818057"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "site_id", "100"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "site_id", "100"),
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
				Config:      testConfigFvRemoteIdMinDependencyWithFvSiteAssociatedAllowExisting,
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
				Config:             testConfigFvRemoteIdMinDependencyWithFvSiteAssociatedAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "description", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "name", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "name", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_pc_tag", "16386"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "remote_pc_tag", "16386"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_vrf_pc_tag", "2818057"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "remote_vrf_pc_tag", "2818057"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "site_id", "100"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test_2", "site_id", "100"),
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
				Config:             testConfigFvRemoteIdMinDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "description", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "name", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_pc_tag", "16386"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_vrf_pc_tag", "2818057"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "site_id", "100"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvRemoteIdAllDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "description", "description"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "name", "name"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "name_alias", "name_alias"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "owner_key", "owner_key"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "owner_tag", "owner_tag"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_pc_tag", "16386"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_vrf_pc_tag", "2818057"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "site_id", "100"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvRemoteIdMinDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_pc_tag", "16386"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_vrf_pc_tag", "2818057"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "site_id", "100"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvRemoteIdResetDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_pc_tag", "16386"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_vrf_pc_tag", "2818057"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "site_id", "100"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "description", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "name", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "owner_tag", ""),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_remote_site_id_mappings.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_pc_tag", "16386"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_vrf_pc_tag", "2818057"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "site_id", "100"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "description", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "name", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "owner_tag", ""),
				),
			},
			// Update with children
			{
				Config:             testConfigFvRemoteIdChildrenDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_pc_tag", "16386"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_vrf_pc_tag", "2818057"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "site_id", "100"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "description", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "name", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.1.value", "value_2"),
				),
			},
			// Import testing with children
			{
				ResourceName:      "aci_remote_site_id_mappings.test",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_pc_tag", "16386"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "remote_vrf_pc_tag", "2818057"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "site_id", "100"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "description", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "name", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.1.value", "value_2"),
				),
			},
			// Update with children removed from config
			{
				Config:             testConfigFvRemoteIdChildrenRemoveFromConfigDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.1.value", "value_2"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvRemoteIdChildrenRemoveOneDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.0.value", "value_2"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvRemoteIdChildrenRemoveAllDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_remote_site_id_mappings.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigFvRemoteIdMinDependencyWithFvSiteAssociatedAllowExisting = testConfigFvSiteAssociatedMinDependencyWithFvCtx + `
resource "aci_remote_site_id_mappings" "test" {
  parent_dn = aci_associated_site.test.id
  remote_pc_tag = "16386"
  remote_vrf_pc_tag = "2818057"
  site_id = "100"
}
resource "aci_remote_site_id_mappings" "test_2" {
  parent_dn = aci_associated_site.test.id
  remote_pc_tag = "16386"
  remote_vrf_pc_tag = "2818057"
  site_id = "100"
  depends_on = [aci_remote_site_id_mappings.test]
}
`

const testConfigFvRemoteIdMinDependencyWithFvSiteAssociated = testConfigFvSiteAssociatedMinDependencyWithFvCtx + `
resource "aci_remote_site_id_mappings" "test" {
  parent_dn = aci_associated_site.test.id
  remote_pc_tag = "16386"
  remote_vrf_pc_tag = "2818057"
  site_id = "100"
}
`

const testConfigFvRemoteIdAllDependencyWithFvSiteAssociated = testConfigFvSiteAssociatedMinDependencyWithFvCtx + `
resource "aci_remote_site_id_mappings" "test" {
  parent_dn = aci_associated_site.test.id
  annotation = "annotation"
  description = "description"
  name = "name"
  name_alias = "name_alias"
  owner_key = "owner_key"
  owner_tag = "owner_tag"
  remote_pc_tag = "16386"
  remote_vrf_pc_tag = "2818057"
  site_id = "100"
}
`

const testConfigFvRemoteIdResetDependencyWithFvSiteAssociated = testConfigFvSiteAssociatedMinDependencyWithFvCtx + `
resource "aci_remote_site_id_mappings" "test" {
  parent_dn = aci_associated_site.test.id
  annotation = "orchestrator:terraform"
  description = ""
  name = ""
  name_alias = ""
  owner_key = ""
  owner_tag = ""
  remote_pc_tag = "16386"
  remote_vrf_pc_tag = "2818057"
  site_id = "100"
}
`
const testConfigFvRemoteIdChildrenDependencyWithFvSiteAssociated = testConfigFvSiteAssociatedMinDependencyWithFvCtx + `
resource "aci_remote_site_id_mappings" "test" {
  parent_dn = aci_associated_site.test.id
  remote_pc_tag = "16386"
  remote_vrf_pc_tag = "2818057"
  site_id = "100"
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

const testConfigFvRemoteIdChildrenRemoveFromConfigDependencyWithFvSiteAssociated = testConfigFvSiteAssociatedMinDependencyWithFvCtx + `
resource "aci_remote_site_id_mappings" "test" {
  parent_dn = aci_associated_site.test.id
  remote_pc_tag = "16386"
  remote_vrf_pc_tag = "2818057"
  site_id = "100"
}
`

const testConfigFvRemoteIdChildrenRemoveOneDependencyWithFvSiteAssociated = testConfigFvSiteAssociatedMinDependencyWithFvCtx + `
resource "aci_remote_site_id_mappings" "test" {
  parent_dn = aci_associated_site.test.id
  remote_pc_tag = "16386"
  remote_vrf_pc_tag = "2818057"
  site_id = "100"
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

const testConfigFvRemoteIdChildrenRemoveAllDependencyWithFvSiteAssociated = testConfigFvSiteAssociatedMinDependencyWithFvCtx + `
resource "aci_remote_site_id_mappings" "test" {
  parent_dn = aci_associated_site.test.id
  remote_pc_tag = "16386"
  remote_vrf_pc_tag = "2818057"
  site_id = "100"
  annotations = []
  tags = []
}
`
