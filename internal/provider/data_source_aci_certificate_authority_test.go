// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourcePkiTP(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigPkiTPDataSource,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_certificate_authority.test", "name", "test_name"),
					resource.TestCheckResourceAttr("data.aci_certificate_authority.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_certificate_authority.test", "certificate_chain", "-----BEGIN CERTIFICATE-----\nMIICODCCAaGgAwIBAgIJAIt8XMntue0VMA0GCSqGSIb3DQEBCwUAMDQxDjAMBgNV\nBAMMBUFkbWluMRUwEwYDVQQKDAxZb3VyIENvbXBhbnkxCzAJBgNVBAYTAlVTMCAX\nDTE4MDEwOTAwNTk0NFoYDzIxMTcxMjE2MDA1OTQ0WjA0MQ4wDAYDVQQDDAVBZG1p\nbjEVMBMGA1UECgwMWW91ciBDb21wYW55MQswCQYDVQQGEwJVUzCBnzANBgkqhkiG\n9w0BAQEFAAOBjQAwgYkCgYEAohG/7axtt7CbSaMP7r+2mhTKbNgh0Ww36C7Ta14i\nv+VmLyKkQHnXinKGhp6uy3Nug+15a+eIu7CrgpBVMQeCiWfsnwRocKcQJWIYDrWl\nXHxGQn31yYKR6mylE7Dcj3rMFybnyhezr5D8GcP85YRPmwG9H2hO/0Y1FUnWu9Iw\nAQkCAwEAAaNQME4wHQYDVR0OBBYEFD0jLXfpkrU/ChzRvfruRs/fy1VXMB8GA1Ud\nIwQYMBaAFD0jLXfpkrU/ChzRvfruRs/fy1VXMAwGA1UdEwQFMAMBAf8wDQYJKoZI\nhvcNAQELBQADgYEAOmvre+5tgZ0+F3DgsfxNQqLTrGiBgGCIymPkP/cBXXkNuJyl\n3ac7tArHQc7WEA4U2R2rZbEq8FC3UJJm4nUVtCPvEh3G9OhN2xwYev79yt6pIn/l\nKU0Td2OpVyo0eLqjoX5u2G90IBWzhyjFbo+CcKMrSVKj1YOdG0E3OuiJf00=\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr("data.aci_certificate_authority.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_certificate_authority.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_certificate_authority.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_certificate_authority.test", "owner_tag", ""),
				),
			},
			{
				Config:      testConfigPkiTPNotExisting,
				ExpectError: regexp.MustCompile("Failed to read aci_certificate_authority data source"),
			},
		},
	})
}
func TestAccDataSourcePkiTPWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "cloud", "1.0(1e)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigPkiTPDataSourceDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_certificate_authority.test", "name", "test_name"),
					resource.TestCheckResourceAttr("data.aci_certificate_authority.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_certificate_authority.test", "certificate_chain", "-----BEGIN CERTIFICATE-----\nMIICODCCAaGgAwIBAgIJAIt8XMntue0VMA0GCSqGSIb3DQEBCwUAMDQxDjAMBgNV\nBAMMBUFkbWluMRUwEwYDVQQKDAxZb3VyIENvbXBhbnkxCzAJBgNVBAYTAlVTMCAX\nDTE4MDEwOTAwNTk0NFoYDzIxMTcxMjE2MDA1OTQ0WjA0MQ4wDAYDVQQDDAVBZG1p\nbjEVMBMGA1UECgwMWW91ciBDb21wYW55MQswCQYDVQQGEwJVUzCBnzANBgkqhkiG\n9w0BAQEFAAOBjQAwgYkCgYEAohG/7axtt7CbSaMP7r+2mhTKbNgh0Ww36C7Ta14i\nv+VmLyKkQHnXinKGhp6uy3Nug+15a+eIu7CrgpBVMQeCiWfsnwRocKcQJWIYDrWl\nXHxGQn31yYKR6mylE7Dcj3rMFybnyhezr5D8GcP85YRPmwG9H2hO/0Y1FUnWu9Iw\nAQkCAwEAAaNQME4wHQYDVR0OBBYEFD0jLXfpkrU/ChzRvfruRs/fy1VXMB8GA1Ud\nIwQYMBaAFD0jLXfpkrU/ChzRvfruRs/fy1VXMAwGA1UdEwQFMAMBAf8wDQYJKoZI\nhvcNAQELBQADgYEAOmvre+5tgZ0+F3DgsfxNQqLTrGiBgGCIymPkP/cBXXkNuJyl\n3ac7tArHQc7WEA4U2R2rZbEq8FC3UJJm4nUVtCPvEh3G9OhN2xwYev79yt6pIn/l\nKU0Td2OpVyo0eLqjoX5u2G90IBWzhyjFbo+CcKMrSVKj1YOdG0E3OuiJf00=\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr("data.aci_certificate_authority.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_certificate_authority.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_certificate_authority.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_certificate_authority.test", "owner_tag", ""),
				),
			},
			{
				Config:      testConfigPkiTPNotExistingFvTenant,
				ExpectError: regexp.MustCompile("Failed to read aci_certificate_authority data source"),
			},
		},
	})
}

const testConfigPkiTPDataSource = testConfigPkiTPMin + `
data "aci_certificate_authority" "test" {
  name = "test_name"
  depends_on = [aci_certificate_authority.test]
}
`

const testConfigPkiTPNotExisting = testConfigPkiTPMin + `
data "aci_certificate_authority" "test_non_existing" {
  name = "non_existing_name"
}
`
const testConfigPkiTPDataSourceDependencyWithFvTenant = testConfigPkiTPMinDependencyWithFvTenant + `
data "aci_certificate_authority" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  depends_on = [aci_certificate_authority.test]
}
`

const testConfigPkiTPNotExistingFvTenant = testConfigFvTenantMin + `
data "aci_certificate_authority" "test_non_existing" {
  parent_dn = aci_tenant.test.id
  name = "non_existing_name"
}
`
