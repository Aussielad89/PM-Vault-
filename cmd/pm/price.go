package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/repoguard/pm-vault/pkg/agents"
	"github.com/repoguard/pm-vault/pkg/ledger"
	"github.com/spf13/cobra"
)

var (
	cardID string
	model  string
)

func newPriceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "price",
		Short: "Run multi-agent valuation jury on inventory",
		Long:  `Queries market sources, applies condition arbitrage, and outputs a consensus health matrix.`,
		Run:   runPrice,
	}
	cmd.Flags().StringVarP(&cardID, "id", "i", "", "Specific card ID to price (default: all)")
	cmd.Flags().StringVarP(&model, "model", "m", "llama3", "Ollama model for analyst agent")
	return cmd
}

func runPrice(cmd *cobra.Command, args []string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get working directory: %v\n", err)
		os.Exit(1)
	}

	pvDir := filepath.Join(cwd, ".pokevault")
	l, err := ledger.Open(pvDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open ledger: %v\n", err)
		os.Exit(1)
	}
	defer l.Close()

	entries, err := l.All()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read ledger: %v\n", err)
		os.Exit(1)
	}

	if len(entries) == 0 {
		fmt.Println("Inventory is empty. Add cards with 'pm add' first.")
		return
	}

	agentEntries := make([]agents.Entry, len(entries))
	for i, e := range entries {
		agentEntries[i] = agents.Entry{
			Name:      e.Name,
			Set:       e.Set,
			Condition: e.Condition,
			Grade:     e.Grade,
			Acquired:  e.Acquired,
			Date:      e.Date,
		}
	}

	fmt.Printf("Running valuation jury on %d entries...\n", len(entries))

	jury := agents.NewJury(model)
	result, err := jury.Run(agentEntries)
	if err != nil {
		fmt.Fprintf(os.Stderr, "jury error: %v\n", err)
		return
	}

	printReport(result)
}

func printReport(result agents.Report) {
	fmt.Println("# PM-Vault Valuation Report")
	fmt.Printf("**Generated:** %s\n\n", result.GeneratedAt)
	fmt.Printf("**Model:** %s\n\n", result.Model)
	fmt.Println("---")
	fmt.Println()

	if len(result.Items) == 0 {
		fmt.Println("No items priced.")
		return
	}

	fmt.Println("## Consensus Prices\n")
	for _, item := range result.Items {
		fmt.Printf("- **%s** (%s) — `%s`\n", item.Name, item.Set, item.Condition)
		fmt.Printf("  - Scraper: $%.2f | Arbitrator: $%.2f | Analyst: $%.2f\n", item.ScraperPrice, item.ArbitratorPrice, item.AnalystPrice)
		fmt.Printf("  - Consensus: **$%.2f** (confidence: %d%%)\n", item.Consensus, item.Confidence)
		fmt.Printf("  - Volatility: %.1f%% | Trend: %s\n", item.Volatility, item.Trend)
		fmt.Println()
	}

	fmt.Println("## Portfolio Health\n")
	fmt.Printf("- Total Value: **$%.2f**\n", result.PortfolioValue)
	fmt.Printf("- 90-Day Change: **%+.2f%%**\n", result.Change90d)
	fmt.Printf("- Diversification Score: **%d/100**\n", result.Diversification)
	fmt.Println()

	fmt.Println("## Jury Deliberation\n")
	for _, log := range result.Logs {
		fmt.Printf("- **[%s]:** %s\n", strings.ToUpper(log.Agent), log.Text)
	}
}
