#!/bin/bash
set -eu

# default: Run tests for all packages with caching.
# bench: Run benchmarks for all packages.
# cover: Run tests and report coverage for all packages.
# html: Run tests and report coverage in HTML format.
# race: Run tests with race detection for the concurrency package.
mode="${1:-default}"

selector=(
    "github.com/huangsam/go-trial/pkg/..."
    "github.com/huangsam/go-trial/internal/..."
)

testout="coverage.out"
testhtml="coverage.html"

case "$mode" in
    default)
        go test "${selector[@]}" ;;
    bench)
        go test -bench=. "${selector[@]}" ;;
    cover)
        go test -cover "${selector[@]}" ;;
    html)
        go test -cover -coverprofile="$testout" "${selector[@]}"
        go tool cover -html="$testout" -o "$testhtml" ;;
    race)
        go test -race github.com/huangsam/go-trial/pkg/concurrency ;;
    *)
        echo "Invalid mode '$mode' detected" && exit 1 ;;
esac
