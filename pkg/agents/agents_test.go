package agents_test

import (
	"testing"
	"time"

	"github.com/repoguard/pm-vault/pkg/agents"
)

func TestJuryRunProducesReport(t *testing.T) {
	entries := []agents.Entry{
		{Name: "Charizard", Set: "Base Set", Condition: "Near-Mint", Acquired: 150},
		{Name: "Umbreon", Set: "Evolving Skies", Condition: "Lightly-Played", Acquired: 45},
	}

	jury := agents.NewJury("llama3")
	report, err := jury.Run(entries)
	if err != nil {
		t.Fatalf("jury run: %v", err)
	}

	if len(report.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(report.Items))
	}
	if report.PortfolioValue <= 0 {
		t.Fatalf("expected positive portfolio value, got %f", report.PortfolioValue)
	}
	if len(report.Logs) != 3 {
		t.Fatalf("expected 3 jury logs, got %d", len(report.Logs))
	}
}

func TestConditionArbitration(t *testing.T) {
	base := 100.0
	adjusted := agents.ArbitrateCondition(base, "Near-Mint")
	if adjusted != base {
		t.Fatalf("Near-Mint should not reduce price, got %f", adjusted)
	}

	adjusted = agents.ArbitrateCondition(base, "Lightly-Played")
	if adjusted != 85.0 {
		t.Fatalf("expected 85.0 for Lightly-Played, got %f", adjusted)
	}
}

func TestConfidenceCalculation(t *testing.T) {
	conf := agents.CalcConfidence(100, 100, 100)
	if conf != 99 {
		t.Fatalf("expected 99 confidence for identical prices, got %d", conf)
	}

	conf = agents.CalcConfidence(100, 50, 50)
	if conf < 40 {
		t.Fatalf("expected minimum 40 confidence, got %d", conf)
	}
}

func TestReportTimestamps(t *testing.T) {
	entries := []agents.Entry{{Name: "X", Set: "Y"}}
	jury := agents.NewJury("llama3")
	report, err := jury.Run(entries)
	if err != nil {
		t.Fatalf("jury run: %v", err)
	}

	_, err = time.Parse(time.RFC3339, report.GeneratedAt)
	if err != nil {
		t.Fatalf("invalid generated_at timestamp: %v", err)
	}
}
