package main

import (
	"flag"
	"log"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/mikelaws/terraform-provider-infra8/cloudforms"
)

func main() {
	debugFlag := flag.Bool("debug", false, "Start provider in debug mode.")
	flag.Parse()

	logFlags := log.Flags()
	logFlags = logFlags &^ (log.Ldate | log.Ltime)
	log.SetFlags(logFlags)

	err = plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cloudforms.Provider})

	if err != nil {
		log.Fatal(err)
	}
}
