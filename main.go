package main

import (
	"flag"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/mikelaws/terraform-provider-infra8/cloudforms"
)

func main() {
	debugFlag := flag.Bool("debug", false, "Start provider in debug mode.")
	flag.Parse()

	plugin.Serve(&plugin.ServeOpts{
		//Call provider
		ProviderFunc: cloudforms.Provider})
}
