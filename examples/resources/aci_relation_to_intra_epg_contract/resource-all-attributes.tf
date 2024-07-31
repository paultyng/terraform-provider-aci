
resource "aci_relation_to_intra_epg_contract" "full_example_endpoint_security_group" {
  parent_dn     = aci_endpoint_security_group.example.id
  annotation    = "annotation"
  contract_name = aci_contract.example.name
  annotations = [
    {
      key   = "key_0"
      value = "value_1"
    }
  ]
  tags = [
    {
      key   = "key_0"
      value = "value_1"
    }
  ]
}
