package main

import (
	"github.com/KubeOpereator/terraform-provider-fusioncompute/fusioncompute"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: fusioncompute.Provider})
}
