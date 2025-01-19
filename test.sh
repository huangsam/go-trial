#!/bin/bash
set -eu

display_help() {
    cat <<'EOF'
Usage: [mode]

mode:
    default: Run tests for all packages.
    cover: Run tests and report coverage for all packages.
    bench: Run benchmarks for all packages.
    concurrency: Run tests for the "concurrency" package 10 times.
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
    "concurrency")
        go test -count=10 io.huangsam/trial/pkg/concurrency ;;
    *)
        echo "Invalid mode '$mode' detected" && exit 1 ;;
esac
