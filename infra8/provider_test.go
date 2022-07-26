package infra8

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"infra8": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("CF_SERVER_IP"); v == "" {
		t.Fatal("Server IP must be set for acceptance tests")
	}

	if v := os.Getenv("CF_USER_NAME"); v == "" {
		t.Fatal("Username must be set for acceptance tests")
	}

	if v := os.Getenv("CF_PASSWORD"); v == "" {
		t.Fatal("Password must be set for acceptance tests")
	}
}
