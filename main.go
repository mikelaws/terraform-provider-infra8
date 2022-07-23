package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/mikelaws/terraform-provider-infra8/cloudforms"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cloudforms.Provider})
}
