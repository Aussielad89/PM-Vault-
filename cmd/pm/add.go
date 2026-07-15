package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/repoguard/pm-vault/pkg/ledger"
	"github.com/spf13/cobra"
)

var (
	name     string
	set      string
	number   string
	condition string
	grade    string
	acquired float64
	date     string
	binder   string
)

func newAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a card or sealed product to your inventory",
		Long:  `Records a new acquisition in the local ledger with condition, cost basis, and binder assignment.`,
		Run:   runAdd,
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "Card or product name (required)")
	cmd.Flags().StringVarP(&set, "set", "s", "", "Set / product line")
	cmd.Flags().StringVarP(&number, "number", "N", "", "Card number / SKU")
	cmd.Flags().StringVarP(&condition, "condition", "c", "Near-Mint", "Condition: Near-Mint, Lightly-Played, Moderately-Played, Heavily-Played, Damaged")
	cmd.Flags().StringVarP(&grade, "grade", "g", "", "Professional grade (e.g., PSA 10, BGS 9.5)")
	cmd.Flags().Float64VarP(&acquired, "acquired", "a", 0.0, "Acquisition cost in configured currency")
	cmd.Flags().StringVarP(&date, "date", "d", "", "Acquisition date (YYYY-MM-DD, default today)")
	cmd.Flags().StringVarP(&binder, "binder", "b", "default", "Binder name")
	cmd.MarkFlagRequired("name")
	return cmd
}

func runAdd(cmd *cobra.Command, args []string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get working directory: %v\n", err)
		os.Exit(1)
	}

	pvDir := filepath.Join(cwd, ".pokevault")
	if _, err := os.Stat(pvDir); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, ".pokevault/ not found. Run 'pm init' first.\n")
		os.Exit(1)
	}

	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	entry := ledger.Entry{
		ID:        generateID(),
		Name:      name,
		Set:       set,
		Number:    number,
		Condition: condition,
		Grade:     grade,
		Acquired:  acquired,
		Date:      date,
		Binder:    binder,
		AddedAt:   time.Now().Format(time.RFC3339),
	}

	l, err := ledger.Open(pvDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open ledger: %v\n", err)
		os.Exit(1)
	}
	defer l.Close()

	if err := l.Append(entry); err != nil {
		fmt.Fprintf(os.Stderr, "failed to save entry: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Added %s [%s] to binder '%s'\n", name, condition, binder)
}

func generateID() string {
	return fmt.Sprintf("card-%d", time.Now().UnixNano())
}
