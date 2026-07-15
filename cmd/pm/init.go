package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var force bool

func newInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize PM-Vault in the current directory",
		Long:  `Creates the .pokevault/ workspace, ledger, and configuration files.`,
		Run:   runInit,
	}
	cmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing .pokevault/ directory")
	return cmd
}

func runInit(cmd *cobra.Command, args []string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get working directory: %v\n", err)
		os.Exit(1)
	}

	pvDir := filepath.Join(cwd, ".pokevault")
	if _, err := os.Stat(pvDir); err == nil && !force {
		fmt.Fprintf(os.Stderr, ".pokevault/ already exists. Use --force to overwrite.\n")
		os.Exit(1)
	}

	if err := os.MkdirAll(filepath.Join(pvDir, "binders"), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create .pokevault/: %v\n", err)
		os.Exit(1)
	}

	files := map[string]string{
		"ledger.json":          "[]\n",
		"snapshots.jsonl":      "",
		"config.yaml":          "pm-vault:\n  version: 1\n  default_condition: Near-Mint\n  currency: USD\n  agents:\n    scraper:\n      sources:\n        - tcgdex\n        - pricecharting\n    arbitrator:\n      condition_penalty:\n        Near-Mint: 0.00\n        Lightly-Played: 0.15\n        Moderately-Played: 0.30\n        Heavily-Played: 0.50\n        Damaged: 0.70\n    analyst:\n      volatility_window: 90\n",
		filepath.Join("binders", "default.json"): "{\n  \"name\": \"Default Binder\",\n  \"layout\": \"3x3\",\n  \"cards\": []\n}\n",
	}

	for name, content := range files {
		path := filepath.Join(pvDir, name)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "failed to create %s: %v\n", path, err)
			os.Exit(1)
		}
	}

	fmt.Printf("PM-Vault initialized at %s\n", pvDir)
}
