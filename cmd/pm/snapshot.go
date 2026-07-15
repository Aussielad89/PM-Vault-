package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/repoguard/pm-vault/pkg/snapstock"
	"github.com/spf13/cobra"
)

var message string

func newSnapshotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "snapshot",
		Short: "Create a git-backed binder snapstock",
		Long:  `Commits the current binder layout, conditions, and portfolio state to .pokevault/history/ as a versioned snapstock.`,
		Run:   runSnapshot,
	}
	cmd.Flags().StringVarP(&message, "message", "m", time.Now().Format("2006-01-02 15:04"), "Snapstock commit message")
	return cmd
}

func runSnapshot(cmd *cobra.Command, args []string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get working directory: %v\n", err)
		os.Exit(1)
	}

	pvDir := filepath.Join(cwd, ".pokevault")
	ss := snapstock.New(pvDir)

	if err := ss.Commit(message); err != nil {
		fmt.Fprintf(os.Stderr, "snapshot failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Snapstock created: %s\n", message)
}
