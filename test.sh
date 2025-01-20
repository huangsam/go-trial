#!/bin/bash
set -eu

display_help() {
    cat <<'EOF'
Usage: [mode]

mode:
    default: Run tests for all packages.
    cover: Run tests and report coverage for all packages.
    bench: Run benchmarks for all packages.
    race: Run tests with race detection for the concurrency package.
EOF
}

help="${1:-none}"

if [[ "$help" == "-h" || "$help" == "--help" ]]; then
    display_help
    exit 0
fi

mode="${1:-default}"

case "$mode" in
    "default")
        go test ./... ;;
    "cover")
        go test -cover ./... ;;
    "bench")
        go test -bench=. ./... ;;
    "race")
        go test -race -count=10 github.com/huangsam/go-trial/pkg/concurrency ;;
    *)
        echo "Invalid mode '$mode' detected" && exit 1 ;;
esac
