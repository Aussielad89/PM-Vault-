package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.1.0-dev"

var rootCmd = &cobra.Command{
	Use:   "pm",
	Short: "PM-Vault — Local-first card portfolio tracker",
	Long: `PM-Vault is a zero-telemetry, terminal-native CLI for collectors.
Digitize, price-track, and simulate investment scenarios for physical cards
and sealed products using an agentic multi-source market evaluator.`,
	Version: version,
}

func init() {
	rootCmd.AddCommand(newInitCmd())
	rootCmd.AddCommand(newAddCmd())
	rootCmd.AddCommand(newPriceCmd())
	rootCmd.AddCommand(newSnapshotCmd())
	rootCmd.AddCommand(newGridCmd())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
