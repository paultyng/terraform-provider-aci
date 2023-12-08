resource "aci_external_management_network_instance_profile" "full_example" {
  annotation  = "annotation"
  description = "description"
  name        = "test_name"
  name_alias  = "name_alias"
  priority    = "level1"
  external_management_network_oob_contracts = [
    {
      annotation                = "orchestrator:terraform"
      priority                  = "level1"
      out_of_band_contract_name = "external_management_network_oob_contracts_1"
    }
  ]
  annotations = [
    {
      key   = "annotations_1"
      value = "value_1"
    }
  ]
}