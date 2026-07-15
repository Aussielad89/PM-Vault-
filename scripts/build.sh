#!/usr/bin/env bash
set -euo pipefail

echo "Building pm..."

mkdir -p bin
go build -o bin/pm ./cmd/pm

echo "Build complete: bin/pm"
./bin/pm --version
