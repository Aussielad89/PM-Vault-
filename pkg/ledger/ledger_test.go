package ledger_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/repoguard/pm-vault/pkg/ledger"
)

func TestOpenAndAppend(t *testing.T) {
	dir := t.TempDir()
	l, err := ledger.Open(dir)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	defer l.Close()

	entry := ledger.Entry{
		Name:     "Test Card",
		Set:      "Test Set",
		Acquired: 10.0,
	}
	if err := l.Append(entry); err != nil {
		t.Fatalf("append: %v", err)
	}

	entries, err := l.All()
	if err != nil {
		t.Fatalf("all: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(entries))
	}
	if entries[0].Name != "Test Card" {
		t.Fatalf("unexpected name: %s", entries[0].Name)
	}
}

func TestPersistsAcrossReopens(t *testing.T) {
	dir := t.TempDir()
	{
		l, err := ledger.Open(dir)
		if err != nil {
			t.Fatalf("open: %v", err)
		}
		l.Append(ledger.Entry{Name: "A"})
		l.Append(ledger.Entry{Name: "B"})
		l.Close()
	}

	l, err := ledger.Open(dir)
	if err != nil {
		t.Fatalf("reopen: %v", err)
	}
	defer l.Close()

	entries, err := l.All()
	if err != nil {
		t.Fatalf("all: %v", err)
	}
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}
}

func TestLedgerJSONFile(t *testing.T) {
	dir := t.TempDir()
	l, err := ledger.Open(dir)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	l.Append(ledger.Entry{Name: "Card1", Set: "Set1"})
	l.Close()

	data, err := os.ReadFile(filepath.Join(dir, "ledger.json"))
	if err != nil {
		t.Fatalf("read ledger.json: %v", err)
	}
	if len(data) == 0 {
		t.Fatal("ledger.json is empty")
	}
}
