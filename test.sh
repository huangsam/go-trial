#!/bin/bash
set -eu

# Default mode is "default"
mode="${1:-default}"

# Run benchmarks for all packages
if [[ "$mode" == "bench" ]]; then
    go test -bench=. ./...

# Run tests for the "concurrency" package 10 times
elif [[ "$mode" == "concurrency" ]]; then
    go test -count=10 io.huangsam/trial/pkg/concurrency

# Run tests for all packages
else
    go test ./...
fi
