package aci

import (
	"fmt"
	"testing"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAciEndpointSecurityGroupEPgSelector_Basic(t *testing.T) {
	var endpoint_security_group_epg_selector models.EndpointSecurityGroupEPgSelector
	fv_tenant_name := acctest.RandString(5)
	fv_ap_name := acctest.RandString(5)
	fv_esg_name := acctest.RandString(5)
	fv_epg_selector_name := acctest.RandString(5)
	description := "endpoint_security_group_epg_selector created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciEndpointSecurityGroupEPgSelectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciEndpointSecurityGroupEPgSelectorConfig_basic(fv_tenant_name, fv_ap_name, fv_esg_name, fv_epg_selector_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciEndpointSecurityGroupEPgSelectorExists("aci_endpoint_security_group_epg_selector.fooendpoint_security_group_epg_selector", &endpoint_security_group_epg_selector),
					testAccCheckAciEndpointSecurityGroupEPgSelectorAttributes(fv_tenant_name, fv_ap_name, fv_esg_name, fv_epg_selector_name, description, &endpoint_security_group_epg_selector),
				),
			},
		},
	})
}

func TestAccAciEndpointSecurityGroupEPgSelector_Update(t *testing.T) {
	var endpoint_security_group_epg_selector models.EndpointSecurityGroupEPgSelector
	fv_tenant_name := acctest.RandString(5)
	fv_ap_name := acctest.RandString(5)
	fv_esg_name := acctest.RandString(5)
	fv_epg_selector_name := acctest.RandString(5)
	description := "endpoint_security_group_epg_selector created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciEndpointSecurityGroupEPgSelectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciEndpointSecurityGroupEPgSelectorConfig_basic(fv_tenant_name, fv_ap_name, fv_esg_name, fv_epg_selector_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciEndpointSecurityGroupEPgSelectorExists("aci_endpoint_security_group_epg_selector.fooendpoint_security_group_epg_selector", &endpoint_security_group_epg_selector),
					testAccCheckAciEndpointSecurityGroupEPgSelectorAttributes(fv_tenant_name, fv_ap_name, fv_esg_name, fv_epg_selector_name, description, &endpoint_security_group_epg_selector),
				),
			},
			{
				Config: testAccCheckAciEndpointSecurityGroupEPgSelectorConfig_basic(fv_tenant_name, fv_ap_name, fv_esg_name, fv_epg_selector_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciEndpointSecurityGroupEPgSelectorExists("aci_endpoint_security_group_epg_selector.fooendpoint_security_group_epg_selector", &endpoint_security_group_epg_selector),
					testAccCheckAciEndpointSecurityGroupEPgSelectorAttributes(fv_tenant_name, fv_ap_name, fv_esg_name, fv_epg_selector_name, description, &endpoint_security_group_epg_selector),
				),
			},
		},
	})
}

func testAccCheckAciEndpointSecurityGroupEPgSelectorConfig_basic(fv_tenant_name, fv_ap_name, fv_esg_name, fv_epg_selector_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_application_profile" "fooapplication_profile" {
		name 		= "%s"
		description = "application_profile created while acceptance testing"
		tenant_dn = aci_tenant.footenant.id
	}

	resource "aci_endpoint_security_group" "fooendpoint_security_group" {
		name 		= "%s"
		description = "endpoint_security_group created while acceptance testing"
		application_profile_dn = aci_application_profile.fooapplication_profile.id
	}

	resource "aci_endpoint_security_group_epg_selector" "fooendpoint_security_group_epg_selector" {
		name 		= "%s"
		description = "endpoint_security_group_epg_selector created while acceptance testing"
		endpoint_security_group_dn = aci_endpoint_security_group.fooendpoint_security_group.id
	}

	`, fv_tenant_name, fv_ap_name, fv_esg_name, fv_epg_selector_name)
}

func testAccCheckAciEndpointSecurityGroupEPgSelectorExists(name string, endpoint_security_group_epg_selector *models.EndpointSecurityGroupEPgSelector) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Endpoint Security Group EPg Selector %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Endpoint Security Group EPg Selector dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		endpoint_security_group_epg_selectorFound := models.EndpointSecurityGroupEPgSelectorFromContainer(cont)
		if endpoint_security_group_epg_selectorFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Endpoint Security Group EPg Selector %s not found", rs.Primary.ID)
		}
		*endpoint_security_group_epg_selector = *endpoint_security_group_epg_selectorFound
		return nil
	}
}

func testAccCheckAciEndpointSecurityGroupEPgSelectorDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "aci_endpoint_security_group_epg_selector" {
			cont, err := client.Get(rs.Primary.ID)
			endpoint_security_group_epg_selector := models.EndpointSecurityGroupEPgSelectorFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Endpoint Security Group EPg Selector %s Still exists", endpoint_security_group_epg_selector.DistinguishedName)
			}
		} else {
			continue
		}
	}
	return nil
}

func testAccCheckAciEndpointSecurityGroupEPgSelectorAttributes(fv_tenant_name, fv_ap_name, fv_esg_name, fv_epg_selector_name, description string, endpoint_security_group_epg_selector *models.EndpointSecurityGroupEPgSelector) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if description != endpoint_security_group_epg_selector.Description {
			return fmt.Errorf("Bad endpoint_security_group_epg_selector Description %s", endpoint_security_group_epg_selector.Description)
		}
		return nil
	}
}
