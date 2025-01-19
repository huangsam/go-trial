#!/bin/bash
set -eu

# Default mode is "default"
mode="${1:-default}"

# Run tests for all packages
if [[ "$mode" == "default" ]]; then
    go test ./...

# Run tests and report coverage for all packages
elif [[ "$mode" == "cover" ]]; then
    go test -cover ./...

# Run benchmarks for all packages
elif [[ "$mode" == "bench" ]]; then
    go test -bench=. ./...

# Run tests for the "concurrency" package 10 times
elif [[ "$mode" == "concurrency" ]]; then
    go test -count=10 io.huangsam/trial/pkg/concurrency

# Exit abruptly
else
    echo "Invalid mode '$mode' detected"
    exit 1
fi
