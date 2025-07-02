package reporters

import (
	"io"

	"github.com/shanejonas/openrpc-linter/types"
)

type Reporter interface {
	Format(results []types.RuleFunctionResult, totalRules int, output io.Writer) error
}
