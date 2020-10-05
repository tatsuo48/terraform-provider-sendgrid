package sendgrid

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceWhitelistIP_Basic(t *testing.T) {
	resourceName := "sendgrid_whitelist_ip.first"
	datasourceName := "data.sendgrid_whitelist_ip.first"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceWhitelistIP_NonExistent(),
				ExpectError: regexp.MustCompile(`Request is failed`),
			},
			{
				Config: testAccDataSourceWhitelistIP_Existent(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, "ip", resourceName, "ip"),
					resource.TestCheckResourceAttrPair(datasourceName, "id", resourceName, "id"),
					resource.TestCheckResourceAttrPair(datasourceName, "created_at", resourceName, "created_at"),
					resource.TestCheckResourceAttrPair(datasourceName, "updated_at", resourceName, "updated_at"),
				),
			},
		},
	})
}

func testAccDataSourceWhitelistIP_Existent() string {
	return fmt.Sprintf(`
resource "sendgrid_whitelist_ip" "first" {
	ip = "192.168.0.1/32"
}
data "sendgrid_whitelist_ip" "first" {
  id = sendgrid_whitelist_ip.first.id
}
`)
}

func testAccDataSourceWhitelistIP_NonExistent() string {
	return fmt.Sprintf(`
data "sendgrid_whitelist_ip" "first" {
  id = 1234567
}
resource "sendgrid_whitelist_ip" "first" {
  ip = "192.168.0.1/32"
}`)
}
