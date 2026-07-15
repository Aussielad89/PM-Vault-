# PM-Vault

![Go Version](https://img.shields.io/badge/Go-1.25%2B-blue)
![Python Version](https://img.shields.io/badge/Python-3.10%2B-blue)
![License](https://img.shields.io/badge/License-MIT-green)

**PM-Vault** is a local-first, zero-telemetry command-line application designed for collectors to digitize, price-track, and simulate investment scenarios for physical cards and sealed products using an agentic multi-source market evaluator.

Rather than trusting a single external marketplace app that tracks user inventory telemetry or pushes biased ads, PM-Vault runs entirely in the terminal, managing inventory via localized JSON transaction ledgers.

## Key Features

- **Multi-Agent Valuation Jury (The Price Consensus)**
  - Scraper Agent: Queries decentralized open APIs (TCGdex, open card databases, localized raw market scrapers)
  - Condition Arbitrator: Simulates variance in condition adjustments (e.g., 15–20% liquidity penalty for Near-Mint vs. Lightly Played)
  - Investment Analyst: Runs historical volatility algorithms to track long-term set value trends, outputting a cryptographic inventory health matrix
- **Git-Style Binder Snapstocks**
  - Saves virtual card binders, grading conditions, and acquisition costs inside a hidden version-controlled workspace directory (`.pokevault/`)
  - Commit changes locally to track visual changes chronologically (e.g., switching from a 3x3 to a 4x4 matrix visualizer)
- **High-Fidelity Terminal Grid View**
  - Advanced layout rendering to project full-color terminal cell art
  - Holo-foil gradient effects using true-color ANSI parsing
  - High-density text tables tracking population reports, card variants, and current portfolio spreads

## Tech Stack

- **Go 1.25+** — Core CLI, ledger, snapstock, and grid engine
- **Bubble Tea** — Terminal UI framework for binder visualization
- **Lip Gloss** — True-color ANSI styling and holo-foil gradients
- **Cobra** — CLI command routing
- **JSONL / JSON** — Local-first transaction storage

## Installation

```bash
git clone https://github.com/Aussielad89/PM-Vault-.git
cd PM-Vault-

go build -o bin/pm ./cmd/pm
./bin/pm --version
```

## Usage

```bash
pm init                  # Initialize .pokevault/ workspace
pm add -n "Charizard" -s "Base Set" -c "Near-Mint" -a 150.00   # Add card
pm price                 # Run multi-agent valuation jury
pm snapshot -m "Binder restructure to 4x4"  # Create snapstock
pm grid                  # Launch high-fidelity terminal grid view
```

## Suggested Topics

`collectibles` `cli` `local-first` `tcg` `golang` `bubble-tea` `agentic-ai` `portfolio` `zero-telemetry` `terminal-ui`
