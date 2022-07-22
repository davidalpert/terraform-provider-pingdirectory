package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccExampleDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccExampleDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.pingdirectory_sync_location.test", "id", "aws-east-1"),
					resource.TestCheckResourceAttr("data.pingdirectory_sync_location.test", "name", "aws-east-1"),
					//resource.TestCheckResourceAttr("data.pingdirectory_sync_location.test", "description", "TBD"),
				),
			},
		},
	})
}

// name = "AWS-East-1"
const testAccExampleDataSourceConfig = `
resource "pingdirectory_sync_location" "test" {
  name = "aws-east-1"
  description = "AWS Region 1"
}

data "pingdirectory_sync_location" "test" {
  id = pingdirectory_sync_location.test.id
}
`
