package reporters

import (
	"io"
	"openrpc-linter/types"
)

type Reporter interface {
	Format(results []types.RuleFunctionResult, totalRules int, output io.Writer) error
}
