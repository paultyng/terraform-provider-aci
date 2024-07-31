---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Generic"
layout: "aci"
page_title: "ACI: aci_host_path_selector"
sidebar_current: "docs-aci-data-source-aci_host_path_selector"
description: |-
  Data source for Host Path Selector
---

# aci_host_path_selector #

Data source for Host Path Selector

## API Information ##

* Class: [infraHPathS](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/infraHPathS/overview)

* Supported in ACI versions: 1.1(1j) and later.

* Distinguished Name Format: `uni/infra/hpaths-{name}`

## GUI Information ##

* Location: `Generic`

## Example Usage ##

```hcl

data "aci_host_path_selector" "example" {
  name = "host_path_selector"
}

```

## Schema ##

### Required ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - The distinguished name (DN) of classes below can be used but currently there is no available resource for it:
    - [infraInfra](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/infraInfra/overview)

  - Default: `uni/infra`
  
* `name` (name) - (string) The name of the Host Path Selector object.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Host Path Selector object.
* `annotation` (annotation) - (string) The annotation of the Host Path Selector object.
* `description` (descr) - (string) The description of the Host Path Selector object.
* `name_alias` (nameAlias) - (string) The name alias of the Host Path Selector object.
* `owner_key` (ownerKey) - (string) The key for enabling clients to own their data for entity correlation.
* `owner_tag` (ownerTag) - (string) A tag for enabling clients to add their own data. For example, to indicate who created this object.

* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
  * `key` (key) - (string) The key used to uniquely identify this configuration object.
  * `value` (value) - (string) The value of the property.

* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
  * `key` (key) - (string) The key used to uniquely identify this configuration object.
  * `value` (value) - (string) The value of the property.