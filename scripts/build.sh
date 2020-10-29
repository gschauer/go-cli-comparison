#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/../"
mkdir -p bin
go build -ldflags='-s -w' -o=bin/ ./...
