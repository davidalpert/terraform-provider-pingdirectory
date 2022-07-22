package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccExampleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccSyncLocationResourceConfig("aws-east-1", "AWS Region 1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("pingdirectory_sync_location.test", "id", "aws-east-1"),
					resource.TestCheckResourceAttr("pingdirectory_sync_location.test", "name", "aws-east-1"),
					resource.TestCheckResourceAttr("pingdirectory_sync_location.test", "description", "AWS Region 1"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "pingdirectory_sync_location.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccSyncLocationResourceConfig("aws-east-1", "AWS Region 1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("pingdirectory_sync_location.test", "id", "aws-east-1"),
					resource.TestCheckResourceAttr("pingdirectory_sync_location.test", "name", "aws-east-1"),
					resource.TestCheckResourceAttr("pingdirectory_sync_location.test", "description", "AWS Region 1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccSyncLocationResourceConfig(name, description string) string {
	return fmt.Sprintf(`
resource "pingdirectory_sync_location" "test" {
  name = %[1]q
  description = %[2]q
}
`, name, description)
}
