package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/grpcwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/plugin"
	simple "github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/provider-simple"
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/tfplugin5"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(simple.Provider())
		},
	})
}
