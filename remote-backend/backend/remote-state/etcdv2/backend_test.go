package etcdv2

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/backend"
)

func TestBackend_impl(t *testing.T) {
	var _ backend.Backend = new(Backend)
}
