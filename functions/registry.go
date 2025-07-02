package functions

import "github.com/shanejonas/openrpc-linter/types"

var FunctionRegistry = make(map[string]types.RuleFunction)

func init() {
	RegisterFunctions()
}

func RegisterFunctions() {
	FunctionRegistry["truthy"] = &TruthyRule{}
}
