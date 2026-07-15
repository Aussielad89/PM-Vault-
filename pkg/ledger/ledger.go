package ledger

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type Entry struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Set       string  `json:"set"`
	Number    string  `json:"number"`
	Condition string  `json:"condition"`
	Grade     string  `json:"grade"`
	Acquired  float64 `json:"acquired"`
	Date      string  `json:"date"`
	Binder    string  `json:"binder"`
	AddedAt   string  `json:"added_at"`
}

type Ledger struct {
	path    string
	entries []Entry
	mu      sync.Mutex
}

func Open(pvDir string) (*Ledger, error) {
	path := filepath.Join(pvDir, "ledger.json")
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("open ledger: %w", err)
	}
	defer f.Close()

	var entries []Entry
	if err := json.NewDecoder(f).Decode(&entries); err != nil {
		if err.Error() != "EOF" {
			return nil, fmt.Errorf("decode ledger: %w", err)
		}
	}

	return &Ledger{path: path, entries: entries}, nil
}

func (l *Ledger) Append(e Entry) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.entries = append(l.entries, e)
	return l.save()
}

func (l *Ledger) All() ([]Entry, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	out := make([]Entry, len(l.entries))
	copy(out, l.entries)
	return out, nil
}

func (l *Ledger) Close() {}

func (l *Ledger) save() error {
	data, err := json.MarshalIndent(l.entries, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(l.path, data, 0644)
}
