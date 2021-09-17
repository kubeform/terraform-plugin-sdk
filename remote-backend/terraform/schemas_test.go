package terraform

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/addrs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/configs/configschema"
)

func simpleTestSchemas() *Schemas {
	provider := simpleMockProvider()
	provisioner := simpleMockProvisioner()

	return &Schemas{
		Providers: map[addrs.Provider]*ProviderSchema{
			addrs.NewDefaultProvider("test"): provider.ProviderSchema(),
		},
		Provisioners: map[string]*configschema.Block{
			"test": provisioner.GetSchemaResponse.Provisioner,
		},
	}
}
