package main

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/repoguard/pm-vault/pkg/grid"
	"github.com/spf13/cobra"
)

func newGridCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grid",
		Short: "Launch high-fidelity terminal binder grid",
		Long:  `Starts the Bubble Tea TUI for full-color card grid visualization with holo-foil gradients and portfolio spread tables.`,
		Run:   runGrid,
	}
	return cmd
}

func runGrid(cmd *cobra.Command, args []string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get working directory: %v\n", err)
		os.Exit(1)
	}

	pvDir := filepath.Join(cwd, ".pokevault")
	p := tea.NewProgram(grid.NewModel(pvDir), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "TUI error: %v\n", err)
		os.Exit(1)
	}
}
