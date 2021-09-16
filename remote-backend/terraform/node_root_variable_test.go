package terraform

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/addrs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/configs"
)

func TestNodeRootVariableExecute(t *testing.T) {
	ctx := new(MockEvalContext)

	n := &NodeRootVariable{
		Addr: addrs.InputVariable{Name: "foo"},
		Config: &configs.Variable{
			Name: "foo",
		},
	}

	diags := n.Execute(ctx, walkApply)
	if diags.HasErrors() {
		t.Fatalf("unexpected error: %s", diags.Err())
	}

}
