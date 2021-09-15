package local

import (
	"flag"
	"os"
	"testing"

	_ "github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/logging"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
