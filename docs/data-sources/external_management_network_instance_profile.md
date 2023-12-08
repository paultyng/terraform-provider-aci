---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
subcategory: "Node Management"
layout: "aci"
page_title: "ACI: aci_external_management_network_instance_profile"
sidebar_current: "docs-aci-data-source-aci_external_management_network_instance_profile"
description: |-
  Data source for External Management Network Instance Profile
---

# aci_external_management_network_instance_profile #

Data source for External Management Network Instance Profile

## API Information ##

* `Class` - [mgmtInstP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/mgmtInstP/overview)

* `Distinguished Name Formats`
  - `uni/tn-mgmt/extmgmt-default/instp-{name}`

## GUI Information ##

* `Location` - `Tenants (mgmt) -> External Management Network Instance Profiles`

## Example Usage ##

```hcl
data "aci_external_management_network_instance_profile" "example" {
  name = "test_name"
}
```

## Schema

### Required

* `name` (name) - (string) The name of the External Management Network Instance Profile object.

### Read-Only

* `id` - (string) The distinguished name (DN) of the External Management Network Instance Profile object.
* `annotation` (annotation) - (string) The annotation of the External Management Network Instance Profile object.
* `description` (descr) - (string) The description of the External Management Network Instance Profile object.
* `name_alias` (nameAlias) - (string) The name alias of the External Management Network Instance Profile object.
* `priority` (prio) - (string) The QoS priority class identifier.

* `external_management_network_oob_contracts` - (list) A list of External Management Network Oob Contracts relationship objects ([mgmtRsOoBCons](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/mgmtRsOoBCons/overview)) pointing to the Out Of Band Contract ([vzOOBBrCP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzOOBBrCP/overview)) object.
  * `annotation` (annotation) - (string) The annotation of the External Management Network Oob Contract object.
  * `priority` (prio) - (string) The Quality of service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.
  * `out_of_band_contract_name` (tnVzOOBBrCPName) - (string) An out-of-band management endpoint group contract consists of switches (leaves/spines) and APICs that are part of the associated out-of-band management zone. Each node in the group is assigned an IP address that is dynamically allocated from the address pool associated with the corresponding out-of-band management zone.
* `annotations` - (list) A list of Annotations objects ([tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)).
  * `key` (key) - (string) The key or password used to uniquely identify this configuration object.
  * `value` (value) - (string) The value of the property.