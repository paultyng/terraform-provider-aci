terraform {
  required_providers {
    aci = {
      source = "ciscodevnet/aci"
    }
  }
}

provider "aci" {
  username = ""
  password = ""
  url      = ""
  insecure = true
}

# Create context profile with hub network
resource "aci_tenant" "terraform_tenant" {
  name = "unmanaged-tenant1"
}

resource "aci_vrf" "vrf" {
  tenant_dn = aci_tenant.terraform_tenant.id
  name      = "unmanaged-VRF2"
}

resource "aci_cloud_context_profile" "ctx1" {
  name                     = "cloud_context_profile"
  description              = "cloud_context_profile created while acceptance testing"
  tenant_dn                = aci_tenant.terraform_tenant.id
  primary_cidr             = "10.0.0.0/16"
  region                   = "us-west-1"
  cloud_vendor             = "aws"
  relation_cloud_rs_to_ctx = aci_vrf.vrf.id
  hub_network              = "uni/tn-infra/gwrouterp-default"
}

# Import Brownfield Virtual Network in Azure cAPIC
resource "aci_tenant" "terraform_tenant" {
  name = "unmanaged-tenant1"
}

resource "aci_vrf" "vrf" {
  tenant_dn = aci_tenant.terraform_tenant.id
  name      = "unmanaged-VRF2"
}
# Azure cloud
resource "aci_cloud_context_profile" "ctx1" {
  name                     = "cloud_context_profile"
  description              = "update desc"
  tenant_dn                = aci_tenant.terraform_tenant.id
  primary_cidr             = "10.1.0.0/16"
  region                   = "eastus2"
  cloud_vendor             = "azure"
  relation_cloud_rs_to_ctx = aci_vrf.vrf.id
  cloud_brownfield         = "/subscriptions/aafaec5f-c828-4651-8504-3a1a01c5daeb/resourceGroups/Unmanaged-test/providers/Microsoft.Network/virtualNetworks/Unmanaged-VNet3"
  access_policy_type       = "read-only"
}

# Import Brownfield VPC in AWS cloud APIC
resource "aci_tenant" "terraform_tenant" {
  name = "tenant1"
}

resource "aci_vrf" "vrf" {
  tenant_dn = aci_tenant.terraform_tenant.id
  name      = "aws_vrf"
}

# AWS cloud
resource "aci_cloud_context_profile" "ctx1" {
  name                     = "cloud_context_profile"
  description              = "import brownfield vpc in aws"
  tenant_dn                = aci_tenant.terraform_tenant.id
  primary_cidr             = "10.2.0.0/24"
  region                   = "us-east-1"
  cloud_vendor             = "aws"
  relation_cloud_rs_to_ctx = aci_vrf.vrf.id
  cloud_brownfield         = "vpc-00a844d6354c53502"
  access_policy_type       = "read-only"
}