package terraform

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/addrs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/configs"
)

// GraphNodeAttachProviderMetaConfigs is an interface that must be implemented
// by nodes that want provider meta configurations attached.
type GraphNodeAttachProviderMetaConfigs interface {
	GraphNodeConfigResource

	// Sets the configuration
	AttachProviderMetaConfigs(map[addrs.Provider]*configs.ProviderMeta)
}
