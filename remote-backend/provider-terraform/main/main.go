package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/builtin/providers/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/grpcwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/tfplugin5"
)

func main() {
	// Provide a binary version of the internal terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(terraform.NewProvider())
		},
	})
}
