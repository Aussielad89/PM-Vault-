# PM-Vault

[![Go Version](https://img.shields.io/badge/Go-1.25%2B-blue)](https://go.dev/dl/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![CI](https://github.com/Aussielad89/PM-Vault-/actions/workflows/ci.yml/badge.svg)](https://github.com/Aussielad89/PM-Vault-/actions/workflows/ci.yml)

**PM-Vault** is a local-first, zero-telemetry command-line application designed for collectors to digitize, price-track, and simulate investment scenarios for physical cards and sealed products using an agentic multi-source market evaluator.

Rather than trusting a single external marketplace app that tracks user inventory telemetry or pushes biased ads, PM-Vault runs entirely in the terminal, managing inventory via localized JSON transaction ledgers.

---

## Table of Contents

- [Why PM-Vault?](#why-pm-vault)
- [Key Features](#key-features)
  - [Multi-Agent Valuation Jury](#multi-agent-valuation-jury)
  - [Git-Style Binder Snapstocks](#git-style-binder-snapstocks)
  - [High-Fidelity Terminal Grid View](#high-fidelity-terminal-grid-view)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Usage](#usage)
  - [Initialize a Vault](#initialize-a-vault)
  - [Add Cards](#add-cards)
  - [Run Valuation Jury](#run-valuation-jury)
  - [Create a Snapstock](#create-a-snapstock)
  - [Launch Grid View](#launch-grid-view)
- [Configuration](#configuration)
- [Security Model](#security-model)
- [Project Structure](#project-structure)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

---

## Why PM-Vault?

Modern collectors face three problems with existing apps:

1. **Privacy erosion** — marketplace apps track inventory, purchase habits, and portfolio value for ad targeting.
2. **Single-source pricing** — relying on one price index ignores condition variance, regional differences, and liquidity discounts.
3. **No local versioning** — when you rearrange a physical binder or change a grading label, there is no way to chronologically track those layout changes.

PM-Vault fixes all three by running **agentic pricing, local git-backed snapstocks, and terminal-native visualization** entirely on your machine.

---

## Key Features

### Multi-Agent Valuation Jury

Instead of pulling a single standard index price, PM-Vault fires off a dedicated local execution script using a specialized agent trio:

- **Scraper Agent** — Connects via a local gateway to query decentralized open APIs (TCGdex, open card databases, or localized raw market scrapers).
- **Condition Arbitrator** — Simulates variance in condition adjustments (e.g., subtracting a 15–20% liquidity penalty for Near-Mint vs. Lightly-Played shifts).
- **Investment Analyst** — Runs historical volatility algorithms to track long-term set value trends, outputting a cryptographic inventory health matrix.

The three agents return individual price estimates, and a consensus score with confidence percentage is computed.

### Git-Style Binder Snapstocks

Saves your virtual card binders, grading conditions, and acquisition costs inside a hidden version-controlled workspace directory (`.pokevault/`).

If you trade cards or rearrange a physical binder layout (e.g., switching from a 3x3 to a 4x4 matrix visualizer), you can commit the changes locally to track visual changes chronologically over time.

### High-Fidelity Terminal Grid View

Utilizes advanced layout rendering to project full-color terminal cell art, holo-foil gradient effects using true-color ANSI parsing, and high-density text tables tracking population reports, card variants, and current portfolio spreads.

---

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Core CLI | Go 1.25+ |
| TUI Framework | Charmbracelet Bubble Tea |
| Styling | Lip Gloss (true-color ANSI) |
| CLI Framework | Cobra |
| Storage | JSON / JSONL |
| Versioning | Git (via snapstock) |
| Agents | Go-native jury (Python optional) |

---

## Installation

### Prerequisites

- Go 1.25 or later
- Git (for snapstocks)

### Build from Source

```bash
git clone https://github.com/Aussielad89/PM-Vault-.git
cd PM-Vault-

go mod tidy
go build -o bin/pm ./cmd/pm

# Verify
./bin/pm --version
```

Or use convenience scripts:

```bash
./scripts/build.sh        # Linux / macOS
scripts\build.bat         # Windows
```

---

## Usage

### Initialize a Vault

```bash
pm init
pm init --force   # Overwrite existing .pokevault/
```

This creates:

```
.pokevault/
├── ledger.json          # Append-only transaction ledger
├── snapshots.jsonl      # Snapstock event log
├── config.yaml          # Condition penalties, currency, agent settings
└── binders/
    └── default.json     # Default binder layout
```

### Add Cards

```bash
pm add -n "Charizard" -s "Base Set" -N "4/102" -c "Near-Mint" -g "PSA 9" -a 150.00
pm add -n "Umbreon VMAX" -s "Evolving Skies" -c "Lightly-Played" -a 45.50
```

### Run Valuation Jury

```bash
pm price
pm price --model qwen2.5:7b   # Use a different Ollama model (optional)
```

Sample output:

```markdown
# PM-Vault Valuation Report
**Generated:** 2026-07-15T20:00:00Z
**Model:** llama3

---

## Consensus Prices

- **Charizard** (Base Set) — `Near-Mint`
  - Scraper: $180.00 | Arbitrator: $180.00 | Analyst: $198.00
  - Consensus: **$186.00** (confidence: 88%)
  - Volatility: 12.4% | Trend: Upward

## Portfolio Health

- Total Value: **$542.30**
- 90-Day Change: **+3.2%**
- Diversification Score: **64/100**
```

### Create a Snapstock

```bash
pm snapshot -m "Binder restructure to 4x4"
pm snapshot -m "Added Umbreon VMAX after grading"
```

This commits the current `.pokevault/` state into `.pokevault/history/<timestamp>/` via git.

### Launch Grid View

```bash
pm grid
```

Opens an interactive Bubble Tea TUI showing your portfolio as a color-coded binder grid with holo-foil gradients.

---

## Configuration

Edit `.pokevault/config.yaml`:

```yaml
pm-vault:
  version: 1
  default_condition: Near-Mint
  currency: USD
  agents:
    scraper:
      sources:
        - tcgdex
        - pricecharting
    arbitrator:
      condition_penalty:
        Near-Mint: 0.00
        Lightly-Played: 0.15
        Moderately-Played: 0.30
        Heavily-Played: 0.50
        Damaged: 0.70
    analyst:
      volatility_window: 90
```

---

## Security Model

PM-Vault runs entirely locally:

1. **No telemetry or cloud calls** — all data stays in `.pokevault/`
2. **Zero-knowledge architecture** — ledger and binders are plain JSON files tracked by git
3. **Local-only agents** — valuation jury runs without external API dependencies by default
4. **Optional market integration** — scraper agent queries open APIs only when explicitly configured

---

## Project Structure

```
PM-Vault-/
├── .github/workflows/ci.yml   # GitHub Actions CI
├── cmd/pm/
│   ├── main.go                # CLI entry point
│   ├── init.go                # Vault initialization
│   ├── add.go                 # Card acquisition entry
│   ├── price.go               # Valuation jury orchestrator
│   ├── snapshot.go            # Git-backed snapstock
│   └── grid.go                # Bubble Tea grid launcher
├── pkg/
│   ├── ledger/                # JSON transaction storage
│   ├── agents/                # Scraper, Arbitrator, Analyst jury
│   ├── snapstock/             # History versioning
│   └── grid/                  # TUI model and ANSI rendering
├── scripts/
│   ├── build.sh               # Linux/macOS build script
│   └── build.bat              # Windows build script
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

---

## Roadmap

- [ ] Real HTTP gateway for TCGdex / PriceCharting integrations
- [ ] Sealed product support (booster boxes, elite trainer boxes)
- [ ] Portfolio export to CSV / PDF
- [ ] Condition photography metadata (image hash + local path)
- [ ] GitHub Action for automated portfolio valuation checks
- [ ] VS Code extension for inline grid preview

---

## Contributing

Contributions are welcome. Please open an issue or PR.

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- [Charmbracelet](https://github.com/charmbracelet) for Bubble Tea and Lip Gloss
- [TCGdex](https://tcgdex.dev) for open card database APIs
- The collectibles and terminal-native tooling communities
