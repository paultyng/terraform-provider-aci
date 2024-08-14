---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "L3Out"
layout: "aci"
page_title: "ACI: aci_relation_from_l3out_consumer_label_to_route_control_profile"
sidebar_current: "docs-aci-data-source-aci_relation_from_l3out_consumer_label_to_route_control_profile"
description: |-
  Data source for ACI Relation From L3Out Consumer Label To Route Control Profile
---

# aci_relation_from_l3out_consumer_label_to_route_control_profile #

Data source for ACI Relation From L3Out Consumer Label To Route Control Profile

## API Information ##

* Class: [l3extRsLblToProfile](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/l3extRsLblToProfile/overview)

* Supported in ACI versions: 5.0(1k) and later.

* Distinguished Name Format: `uni/tn-{name}/out-{name}/conslbl-{name}/rslblToProfile-[{tDn}]-{direction}`

## GUI Information ##

* Location: `Tenants -> Networking -> L3Outs -> Policy -> Main -> Consumer Label`

## Example Usage ##

```hcl

data "aci_relation_from_l3out_consumer_label_to_route_control_profile" "example_l3out_consumer_label" {
  parent_dn = aci_l3out_consumer_label.example.id
  direction = "import"
  target_dn = aci_route_control_profile.example.id
}

```

## Schema ##

### Required ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_l3out_consumer_label](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/l3out_consumer_label) ([l3extConsLbl](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/l3extConsLbl/overview))
  - The distinguished name (DN) of classes below can be used but currently there is no available resource for it:
    - [l3extConsLblDef](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/l3extConsLblDef/overview)

* `direction` (direction) - (string) The connector direction.
  - Valid Values: `export`, `import`.
* `target_dn` (tDn) - (string) The distinguished name (DN) of the Route Control Profile object.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Relation From L3Out Consumer Label To Route Control Profile object.
* `annotation` (annotation) - (string) The annotation of the Relation From L3Out Consumer Label To Route Control Profile object.

* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
  * `key` (key) - (string) The key used to uniquely identify this configuration object.
  * `value` (value) - (string) The value of the property.

* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
  * `key` (key) - (string) The key used to uniquely identify this configuration object.
  * `value` (value) - (string) The value of the property.