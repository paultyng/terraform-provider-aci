
resource "aci_epg_useg_criterion_mac_attribute" "full_example_epg_useg_criterion" {
  parent_dn   = aci_epg_useg_criterion.example.id
  annotation  = "annotation"
  description = "description"
  mac         = "AA:BB:CC:DD:EE:FF"
  name        = "mac_attr"
  name_alias  = "name_alias"
  owner_key   = "owner_key"
  owner_tag   = "owner_tag"
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