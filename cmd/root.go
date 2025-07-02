package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openrpc-linter",
	Short: "A linter for OpenRPC documents",
	Long:  "Fast, extensible linter for OpenRPC documents",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
} 