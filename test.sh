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
        go test github.com/huangsam/go-trial/pkg/... ;;
    "cover")
        go test -cover github.com/huangsam/go-trial/pkg/... ;;
    "bench")
        go test -bench=. github.com/huangsam/go-trial/pkg/... ;;
    "race")
        go test -race github.com/huangsam/go-trial/pkg/concurrency ;;
    *)
        echo "Invalid mode '$mode' detected" && exit 1 ;;
esac
