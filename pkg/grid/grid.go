package grid

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/repoguard/pm-vault/pkg/ledger"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	pvDir string
	width int
	height int
}

func NewModel(pvDir string) Model {
	return Model{pvDir: pvDir}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	if msg == tea.KeyCtrlC {
		return m, tea.Quit
	}
	return m, nil
}

func (m Model) View() string {
	var b strings.Builder
	b.WriteString(lipgloss.NewStyle().Bold(true).Render("🎴 PM-Vault Binder Grid"))
	b.WriteString("\n\n")

	if m.width == 0 || m.height == 0 {
		b.WriteString("Loading...")
		return b.String()
	}

	entries := loadEntries(m.pvDir)
	if len(entries) == 0 {
		b.WriteString("No cards in inventory. Use 'pm add' to add cards.")
		b.WriteString("\n")
		b.WriteString(lipgloss.NewStyle().Faint(true).Render("Press Ctrl+C to exit"))
		return b.String()
	}

	b.WriteString(fmt.Sprintf("Portfolio: %d items\n\n", len(entries)))

	for i, e := range entries {
		if i > 0 && i%3 == 0 {
			b.WriteString("\n")
		}
		cardStyle := lipgloss.NewStyle().Width(20).Border(lipgloss.RoundedBorder()).Padding(0, 1)
		holo := holoGradient(e.Condition)
		b.WriteString(holo.Render(cardStyle.Render(fmt.Sprintf("%s\n%s", e.Name, e.Condition))))
		b.WriteString("  ")
	}

	b.WriteString("\n\n")
	b.WriteString(lipgloss.NewStyle().Faint(true).Render("Press Ctrl+C to exit"))
	return b.String()
}

func loadEntries(pvDir string) []ledger.Entry {
	path := filepath.Join(pvDir, "ledger.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var entries []ledger.Entry
	json.Unmarshal(data, &entries)
	return entries
}

func holoGradient(condition string) lipgloss.Style {
	switch condition {
	case "Near-Mint":
		return lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	case "Lightly-Played":
		return lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
	case "Moderately-Played":
		return lipgloss.NewStyle().Foreground(lipgloss.Color("208"))
	case "Heavily-Played":
		return lipgloss.NewStyle().Foreground(lipgloss.Color("166"))
	case "Damaged":
		return lipgloss.NewStyle().Foreground(lipgloss.Color("124"))
	default:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
	}
}
