package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"pingdirectory": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
	valid := true
	if os.Getenv("PING_SYNC_BASE_URL") == "" {
		t.Errorf("PING_SYNC_BASE_URL is required to run acceptance/integration tests")
		valid = false
	}
	if os.Getenv("PING_SYNC_DN_USER") == "" {
		t.Errorf("PING_SYNC_DN_USER is required to run acceptance/integration tests")
		valid = false
	}
	if os.Getenv("PING_SYNC_DN_PASS") == "" {
		t.Errorf("PING_SYNC_DN_PASS is required to run acceptance/integration tests")
		valid = false
	}
	if !valid {
		t.Fatal()
	}
}
