# Contributing to PM-Vault

Thanks for your interest in improving PM-Vault. This is a local-first, zero-telemetry project, and we value privacy, performance, and terminal-native UX.

## Development Setup

```bash
git clone https://github.com/Aussielad89/PM-Vault-.git
cd PM-Vault-
go mod tidy
go build ./cmd/pm
go test ./...
```

## Project Structure

- `cmd/pm/` — CLI commands (init, add, price, snapshot, grid)
- `pkg/ledger/` — JSON transaction storage
- `pkg/agents/` — Multi-agent valuation jury
- `pkg/snapstock/` — Git-backed versioned binders
- `pkg/grid/` — Bubble Tea TUI and ANSI rendering
- `pkg/gateway/` — Local HTTP gateway for market APIs

## Commit Conventions

- `feat:` — new feature
- `fix:` — bug fix
- `docs:` — documentation only
- `refactor:` — code change that neither fixes a bug nor adds a feature
- `test:` — adding or updating tests
- `chore:` — build, CI, or tooling changes

## Pull Requests

1. Fork the repo
2. Create a feature branch
3. Ensure `go test ./...` passes
4. Open a PR with a clear description

## Code of Conduct

Be respectful. Privacy and collector sovereignty are core values of this project.
