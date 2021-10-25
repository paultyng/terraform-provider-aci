module github.com/terraform-providers/terraform-provider-aci

go 1.12

require (
	github.com/ciscoecosystem/aci-go-client v1.12.32
	github.com/ghodss/yaml v1.0.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.4.3
)

replace github.com/ciscoecosystem/aci-go-client => ../../ciscoecosystem/aci-go-client