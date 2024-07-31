// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceInfraHPathS(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config: testConfigInfraHPathSMinAllowExisting,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test", "name", "host_path_selector"),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test_2", "name", "host_path_selector"),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test_2", "owner_tag", ""),
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
				Config:      testConfigInfraHPathSMinAllowExisting,
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
				Config: testConfigInfraHPathSMinAllowExisting,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test", "name", "host_path_selector"),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test_2", "name", "host_path_selector"),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.allow_test_2", "owner_tag", ""),
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
				Config: testConfigInfraHPathSMin,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "name", "host_path_selector"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "description", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "owner_tag", ""),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config: testConfigInfraHPathSAll,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "name", "host_path_selector"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "owner_key", "owner_key_1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "owner_tag", "owner_tag_1"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config: testConfigInfraHPathSMin,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "name", "host_path_selector"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "owner_key", "owner_key_1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "owner_tag", "owner_tag_1"),
				),
			},
			// Update with empty strings config or default value
			{
				Config: testConfigInfraHPathSReset,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "name", "host_path_selector"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "description", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "owner_tag", ""),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_host_path_selector.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config: testConfigInfraHPathSChildren,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "description", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "tags.#", "2"),
				),
			},
			// Update with children removed from config
			{
				Config: testConfigInfraHPathSChildrenRemoveFromConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "description", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config: testConfigInfraHPathSChildrenRemoveOne,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "description", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config: testConfigInfraHPathSChildrenRemoveAll,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "description", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_host_path_selector.test", "tags.#", "0"),
				),
			},
		},
	})
}

const testConfigInfraHPathSMinAllowExisting = `
resource "aci_host_path_selector" "allow_test" {
  name = "host_path_selector"
}
resource "aci_host_path_selector" "allow_test_2" {
  name = "host_path_selector"
  depends_on = [aci_host_path_selector.allow_test]
}
`

const testConfigInfraHPathSMin = `
resource "aci_host_path_selector" "test" {
  name = "host_path_selector"
}
`

const testConfigInfraHPathSAll = `
resource "aci_host_path_selector" "test" {
  name = "host_path_selector"
  annotation = "annotation"
  description = "description_1"
  name_alias = "name_alias_1"
  owner_key = "owner_key_1"
  owner_tag = "owner_tag_1"
}
`

const testConfigInfraHPathSReset = `
resource "aci_host_path_selector" "test" {
  name = "host_path_selector"
  annotation = "orchestrator:terraform"
  description = ""
  name_alias = ""
  owner_key = ""
  owner_tag = ""
}
`
const testConfigInfraHPathSChildren = `
resource "aci_host_path_selector" "test" {
  name = "host_path_selector"
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

const testConfigInfraHPathSChildrenRemoveFromConfig = `
resource "aci_host_path_selector" "test" {
  name = "host_path_selector"
}
`

const testConfigInfraHPathSChildrenRemoveOne = `
resource "aci_host_path_selector" "test" {
  name = "host_path_selector"
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

const testConfigInfraHPathSChildrenRemoveAll = `
resource "aci_host_path_selector" "test" {
  name = "host_path_selector"
  annotations = []
  tags = []
}
`