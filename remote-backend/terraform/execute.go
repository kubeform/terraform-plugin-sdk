package terraform

import "github.com/hashicorp/terraform-plugin-sdk/v2/remote-backend/tfdiags"

// GraphNodeExecutable is the interface that graph nodes must implement to
// enable execution.
type GraphNodeExecutable interface {
	Execute(EvalContext, walkOperation) tfdiags.Diagnostics
}
