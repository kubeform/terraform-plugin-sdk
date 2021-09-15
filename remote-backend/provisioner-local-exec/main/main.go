package main

import (
	localexec "github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/builtin/provisioners/local-exec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/grpcwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/tfplugin5"
)

func main() {
	// Provide a binary version of the internal terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProvisionerFunc: func() tfplugin5.ProvisionerServer {
			return grpcwrap.Provisioner(localexec.New())
		},
	})
}
