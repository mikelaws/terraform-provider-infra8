package infra8

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider : Defines provider schema
// Contains registry of Data sources and Resources
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"ip": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The IP of Cloudforms service",
				DefaultFunc: schema.EnvDefaultFunc("CF_SERVER_IP", nil),
			},

			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The UserName of ManageIQ service",
				DefaultFunc: schema.EnvDefaultFunc("CF_USER_NAME", nil),
			},

			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The Password of ManageIQ service",
				DefaultFunc: schema.EnvDefaultFunc("CF_PASSWORD", nil),
			},
		},

		//Supported Data Source by this provider
		DataSourcesMap: map[string]*schema.Resource{
			"infra8_service":          dataSourceServiceDetail(),
			"infra8_service_template": dataSourceServiceTemplate(),
		},

		//Supported Resources by this provider
		ResourcesMap: map[string]*schema.Resource{
			"infra8_service_request": resourceServiceRequest(),
		},

		//Provider configuration
		ConfigureContextFunc: providerConfigure,
	}
}

// providerConfigure : This funtion will read provider module form '.tf' file store data into config structure
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	config, err := CFConnect(d)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	log.Println("[DEBUG] Connecting to Cloudforms...")
	return config, diags
}
