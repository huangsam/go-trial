#!/bin/bash
set -eu

mode="${1:-default}"

if [[ "$mode" == "bench" ]]; then
    go test -bench=. ./...
elif [[ "$mode" == "concurrency" ]]; then
    go test -count=10 io.huangsam/trial/pkg/concurrency
else
    go test ./...
fi
