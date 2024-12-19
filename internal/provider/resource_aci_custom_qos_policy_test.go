// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceQosCustomPolWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigQosCustomPolMinDependencyWithFvTenantAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test_2", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test_2", "owner_tag", ""),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigQosCustomPolMinDependencyWithFvTenantAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigQosCustomPolMinDependencyWithFvTenantAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test_2", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.allow_test_2", "owner_tag", ""),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigQosCustomPolMinDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "owner_tag", ""),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigQosCustomPolAllDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "owner_key", "owner_key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "owner_tag", "owner_tag_1"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigQosCustomPolMinDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "name", "test_name"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigQosCustomPolResetDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "owner_tag", ""),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_custom_qos_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigQosCustomPolChildrenDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.description", "description_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.from", "0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.name", "name_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.priority", "level1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.target", "AF11"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.target_cos", "0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.to", "0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.description", "description_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.from", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.name", "name_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.name_alias", "name_alias_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.priority", "level2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.target", "AF12"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.target_cos", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.to", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.description", "description_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.from", "AF11"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.name", "name_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.priority", "level1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.target", "AF11"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.target_cos", "0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.to", "AF11"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.description", "description_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.from", "AF12"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.name", "name_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.name_alias", "name_alias_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.priority", "level2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.target", "AF12"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.target_cos", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.to", "AF12"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.1.value", "test_value"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_custom_qos_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigQosCustomPolChildrenRemoveFromConfigDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.#", "2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.description", "description_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.from", "0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.name", "name_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.priority", "level1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.target", "AF11"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.target_cos", "0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.to", "0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.#", "2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.description", "description_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.from", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.name", "name_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.name_alias", "name_alias_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.priority", "level2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.target", "AF12"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.target_cos", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.1.to", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.#", "2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.annotation", "annotation_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.#", "2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.description", "description_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.from", "AF11"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.name", "name_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.priority", "level1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.target", "AF11"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.target_cos", "0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.to", "AF11"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.#", "2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.description", "description_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.from", "AF12"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.name", "name_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.name_alias", "name_alias_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.priority", "level2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.target", "AF12"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.target_cos", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.1.to", "AF12"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.#", "2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigQosCustomPolChildrenRemoveOneDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.tags.#", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.description", "description_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.from", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.name", "name_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.name_alias", "name_alias_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.priority", "level2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.target", "AF12"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.target_cos", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.0.to", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.#", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.annotation", "annotation_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.tags.#", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.description", "description_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.from", "AF12"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.name", "name_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.name_alias", "name_alias_2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.priority", "level2"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.target", "AF12"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.target_cos", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.0.to", "AF12"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.#", "1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigQosCustomPolChildrenRemoveAllDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dot1p_classifiers.#", "0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "dscp_to_priority_maps.#", "0"),
					resource.TestCheckResourceAttr("aci_custom_qos_policy.test", "tags.#", "0"),
				),
			},
		},
		CheckDestroy: testCheckResourceDestroy,
	})
}

const testConfigQosCustomPolMinDependencyWithFvTenantAllowExisting = testConfigFvTenantMin + `
resource "aci_custom_qos_policy" "allow_test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
}
resource "aci_custom_qos_policy" "allow_test_2" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  depends_on = [aci_custom_qos_policy.allow_test]
}
`

const testConfigQosCustomPolMinDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_custom_qos_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
}
`

const testConfigQosCustomPolAllDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_custom_qos_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  annotation = "annotation"
  description = "description_1"
  name_alias = "name_alias_1"
  owner_key = "owner_key_1"
  owner_tag = "owner_tag_1"
}
`

const testConfigQosCustomPolResetDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_custom_qos_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  annotation = "orchestrator:terraform"
  description = ""
  name_alias = ""
  owner_key = ""
  owner_tag = ""
}
`
const testConfigQosCustomPolChildrenDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_custom_qos_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
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
  dot1p_classifiers = [
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
      from = "0"
      name = "name_1"
      name_alias = "name_alias_1"
      priority = "level1"
      target = "AF11"
      target_cos = "0"
      to = "0"
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
      from = "1"
      name = "name_2"
      name_alias = "name_alias_2"
      priority = "level2"
      target = "AF12"
      target_cos = "1"
      to = "1"
    },
  ]
  dscp_to_priority_maps = [
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
      from = "AF11"
      name = "name_1"
      name_alias = "name_alias_1"
      priority = "level1"
      target = "AF11"
      target_cos = "0"
      to = "AF11"
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
      from = "AF12"
      name = "name_2"
      name_alias = "name_alias_2"
      priority = "level2"
      target = "AF12"
      target_cos = "1"
      to = "AF12"
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

const testConfigQosCustomPolChildrenRemoveFromConfigDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_custom_qos_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
}
`

const testConfigQosCustomPolChildrenRemoveOneDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_custom_qos_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  annotations = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
  dot1p_classifiers = [ 
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
	  from = "1"
	  name = "name_2"
	  name_alias = "name_alias_2"
	  priority = "level2"
	  target = "AF12"
	  target_cos = "1"
	  to = "1"
	},
  ]
  dscp_to_priority_maps = [ 
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
	  from = "AF12"
	  name = "name_2"
	  name_alias = "name_alias_2"
	  priority = "level2"
	  target = "AF12"
	  target_cos = "1"
	  to = "AF12"
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

const testConfigQosCustomPolChildrenRemoveAllDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_custom_qos_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  annotations = []
  dot1p_classifiers = []
  dscp_to_priority_maps = []
  tags = []
}
`
