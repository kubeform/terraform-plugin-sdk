package terraform

import (
	backendInit "github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/backend/init"
)

func init() {
	// Initialize the backends
	backendInit.Init(nil)
}
