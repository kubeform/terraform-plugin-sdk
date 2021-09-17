package terraform

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/addrs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/configs"
)

// GraphNodeAttachProvider is an interface that must be implemented by nodes
// that want provider configurations attached.
type GraphNodeAttachProvider interface {
	// ProviderName with no module prefix. Example: "aws".
	ProviderAddr() addrs.AbsProviderConfig

	// Sets the configuration
	AttachProvider(*configs.Provider)
}
