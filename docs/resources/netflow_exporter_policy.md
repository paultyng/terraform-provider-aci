---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Tenant Policies"
layout: "aci"
page_title: "ACI: aci_netflow_exporter_policy"
sidebar_current: "docs-aci-resource-aci_netflow_exporter_policy"
description: |-
  Manages ACI NetFlow Exporter Policy
---

# aci_netflow_exporter_policy #

Manages ACI NetFlow Exporter Policy



## API Information ##

* Class: [netflowExporterPol](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/netflowExporterPol/overview)

* Supported in ACI versions: 2.2(1k) and later.

* Distinguished Name Formats:
  - `uni/infra/exporterpol-{name}`
  - `uni/tn-{name}/exporterpol-{name}`

## GUI Information ##

* Locations:
  - `Tenants -> Policies -> NetFlow -> NetFlow Exporters`
  - `Fabric -> Access Policies -> Policies -> Interface -> NetFlow -> NetFlow Exporters`

## Example Usage ##

The configuration snippet below creates a NetFlow Exporter Policy with only required attributes.

```hcl

resource "aci_netflow_exporter_policy" "example_tenant" {
  parent_dn              = aci_tenant.example.id
  destination_ip_address = "2.2.2.1"
  destination_port       = "https"
  name                   = "netfow_exporter"
}

```
The configuration snippet below shows all possible attributes of the NetFlow Exporter Policy.

!> This example might not be valid configuration and is only used to show all possible attributes.

```hcl

resource "aci_netflow_exporter_policy" "full_example_tenant" {
  parent_dn              = aci_tenant.example.id
  annotation             = "annotation"
  description            = "description_1"
  qos_dscp_value         = "AF11"
  destination_ip_address = "2.2.2.1"
  destination_port       = "https"
  name                   = "netfow_exporter"
  name_alias             = "name_alias_1"
  owner_key              = "owner_key_1"
  owner_tag              = "owner_tag_1"
  source_ip_type         = "custom-src-ip"
  source_ip_address      = "1.1.1.1/10"
  version                = "v9"
  relation_to_vrf = [
    {
      annotation = "annotation_1"
      target_dn  = aci_vrf.example.id
    }
  ]
  relation_to_epg = [
    {
      annotation = "annotation_1"
      target_dn  = aci_application_epg.example.id
    }
  ]
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

```

All examples for the NetFlow Exporter Policy resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-aci/tree/master/examples/resources/aci_netflow_exporter_policy) folder.

## Schema ##

### Required ###

* `destination_ip_address` (dstAddr) - (string) The destination IP address of the remote node.
* `destination_port` (dstPort) - (string) The destination port of the remote node.
  - Valid Values:
    * `dns`, `ftpData`, `http`, `https`, `pop3`, `rtsp`, `smtp`, `ssh`, `unspecified`.
    * Or a value in the range of `0` to `65535`.
* `name` (name) - (string) The name of the NetFlow Exporter Policy object.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the NetFlow Exporter Policy object.

### Optional ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_tenant](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tenant) ([fvTenant](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvTenant/overview))
  - The distinguished name (DN) of classes below can be used but currently there is no available resource for it:
    - [infraInfra](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/infraInfra/overview)

  - Default: `uni/infra`
* `annotation` (annotation) - (string) The annotation of the NetFlow Exporter Policy object.
  - Default: `orchestrator:terraform`
* `description` (descr) - (string) The description of the NetFlow Exporter Policy object.
* `qos_dscp_value` (dscp) - (string) The DSCP value of the NetFlow Exporter Policy object.
  - Default: `CS2`
  - Valid Values:
    * `AF11`, `AF12`, `AF13`, `AF21`, `AF22`, `AF23`, `AF31`, `AF32`, `AF33`, `AF41`, `AF42`, `AF43`, `CS0`, `CS1`, `CS2`, `CS3`, `CS4`, `CS5`, `CS6`, `CS7`, `EF`, `VA`.
    * Or a value in the range of `0` to `63`.
* `name_alias` (nameAlias) - (string) The name alias of the NetFlow Exporter Policy object.
* `owner_key` (ownerKey) - (string) The key for enabling clients to own their data for entity correlation.
* `owner_tag` (ownerTag) - (string) A tag for enabling clients to add their own data. For example, to indicate who created this object.
* `source_ip_type` (sourceIpType) - (string) The type of the source IP address for the NetFlow Exporter Policy object.
  - Default: `custom-src-ip`
  - Valid Values: `custom-src-ip`, `inband-mgmt-ip`, `oob-mgmt-ip`, `ptep`.
* `source_ip_address` (srcAddr) - (string) The source IP address.
* `version` (ver) - (string) The NetFlow Exporter Version of the NetFlow Exporter Policy object.
  - Default: `v9`
  - Valid Values: `cisco-v1`, `v5`, `v9`.

* `relation_to_vrf` - (list) A list of Relation From NetFlow Exporter To VRF (ACI object [netflowRsExporterToCtx](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/netflowRsExporterToCtx/overview)) pointing to VRF (ACI Object [fvCtx](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvCtx/overview)) which can be configured using the [aci_vrf](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/vrf) resource.
    - Max Items: 1
  

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation From NetFlow Exporter To VRF object.
      - Default: `orchestrator:terraform`
  * `target_dn` (tDn) - (string) The distinguished name of the target.

      
* `relation_to_epg` - (list) A list of Relation From NetFlow Exporter To EPG (ACI object [netflowRsExporterToEPg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/netflowRsExporterToEPg/overview)). This relation can point to multiple ACI objects:
    - [fvAEPg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvAEPg/overview) which can be configured using the [aci_application_epg](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/application_epg) resource.
    - [l3extInstP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/l3extInstP/overview) which can be configured using the [aci_external_network_instance_profile](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/external_network_instance_profile) resource.
    - [l2extInstP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/l2extInstP/overview) which can be configured using the [aci_l2out_extepg](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/l2out_extepg) resource.
  

  #### Optional ####
    
  * `annotation` (annotation) - (string) The annotation of the Relation From NetFlow Exporter To EPG object.
      - Default: `orchestrator:terraform`
  * `target_dn` (tDn) - (string) The distinguished name of the target.

* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). Annotations can also be configured using a separate [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  
  #### Required ####
  
  * `key` (key) - (string) The key used to uniquely identify this configuration object.
  * `value` (value) - (string) The value of the property.

* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). Tags can also be configured using a separate [aci_tag](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tag) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  
  #### Required ####
  
  * `key` (key) - (string) The key used to uniquely identify this configuration object.
  * `value` (value) - (string) The value of the property.

## Importing

An existing NetFlow Exporter Policy can be [imported](https://www.terraform.io/docs/import/index.html) into this resource with its distinguished name (DN), via the following command:

```
terraform import aci_netflow_exporter_policy.example_tenant uni/infra/exporterpol-{name}
```

Starting in Terraform version 1.5, an existing NetFlow Exporter Policy can be imported
using [import blocks](https://developer.hashicorp.com/terraform/language/import) via the following configuration:

```
import {
  id = "uni/infra/exporterpol-{name}"
  to = aci_netflow_exporter_policy.example_tenant
}
```