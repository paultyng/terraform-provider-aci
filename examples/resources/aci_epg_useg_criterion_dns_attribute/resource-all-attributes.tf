
resource "aci_epg_useg_criterion_dns_attribute" "full_example_epg_useg_criterion" {
  parent_dn   = aci_epg_useg_criterion.example.id
  annotation  = "annotation"
  description = "description"
  filter      = "filter"
  name        = "dns_attribute"
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