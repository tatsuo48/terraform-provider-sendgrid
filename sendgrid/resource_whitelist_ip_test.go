package sendgrid

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/sendgrid/sendgrid-go"
)

func TestAccResourceWhitelistIP_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWhitelistIP(),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceWhitelistIPExists("sendgrid_whitelist_ip.first"),
				),
			},
		},
	})
}

func testAccResourceWhitelistIP() string {
	return fmt.Sprintf(`
resource "sendgrid_whitelist_ip" "first" {
	ip = "192.168.0.1/32"
}
`)
}

func testAccResourceWhitelistIPExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Whitelist IP id is set")
		}
		apiKey := testAccProvider.Meta().(string)
		path := fmt.Sprintf("/v3/access_settings/whitelist/%s", rs.Primary.ID)
		request := sendgrid.GetRequest(apiKey, path, "https://api.sendgrid.com")
		request.Method = "GET"

		r, err := sendgrid.API(request)
		if err != nil {
			return fmt.Errorf("Request is failed")
		}
		if r.StatusCode != 200 {
			return fmt.Errorf("Request is failed\n path: %s\n body: %s\n status code: %d\n ", path, r.Body, r.StatusCode)
		}
		return nil
	}
}
