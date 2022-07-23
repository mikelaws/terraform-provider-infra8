package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/mikelaws/terraform-provider-infra8/infra8"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: infra8.Provider})
}
