package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/grpcwrap"
	plugin "github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/plugin6"
	simple "github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/provider-simple-v6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/tfplugin6"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin6.ProviderServer {
			return grpcwrap.Provider6(simple.Provider())
		},
	})
}
