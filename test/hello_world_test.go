package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestHelloWorld(t *testing.T) {
	t.Parallel()

	// Generate a random DSN name to prevent a naming conflict
	uniqueID := random.UniqueId()
	uniqueName := fmt.Sprintf("Hello-World-%s", uniqueID)

	// Specify the test case folder and "-var" options
	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/hello-world/",
		Vars: map[string]interface{}{
			"resource_group": uniqueName,
			"dns-prefix":     uniqueName,
		},
	}

	// Terraform init, apply, output, and destroy
	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Get Application URL from output
	fqdn := terraform.Output(t, terraformOptions, "fqdn")
	query := fmt.Sprintf("http://%s", fqdn)

	// Validate the provisioned application
	http_helper.HttpGetWithCustomValidation(t, query, func(status int, content string) bool {
		return status == 200 &&
			strings.Contains(content, "Hello")
	})
}
