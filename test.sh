#!/bin/bash
set -eu

mode="${1:-default}"

if [[ "$mode" == "bench" ]]; then
    go test -bench=. ./...
else
    go test ./...
fi
