package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/mikelaws/terraform-provider-cloudforms/cloudforms"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		//Call provider
		ProviderFunc: cloudforms.Provider})
}
