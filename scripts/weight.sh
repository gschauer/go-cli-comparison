#!/usr/bin/env bash
set -euo pipefail

declare -r GOWEIGHT_BIN="${GOPATH:-$(go env GOPATH)}/bin/goweight"

[[ -f "${GOWEIGHT_BIN}" ]] ||
  (cd "${TMPDIR:-/tmp}" && go get -v github.com/jondot/goweight)

cd "$(dirname "$0")/../"
[[ -d tmp ]] || mkdir tmp
for f in ./cmd/*/main.go; do
  f="$(dirname "$f")"
  echo "Determining weight of $f..."
  "${GOWEIGHT_BIN}" "$f" >"tmp/$(basename "$f").log"
done
