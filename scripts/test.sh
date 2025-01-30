#!/bin/bash
set -eu

display_help () {
    cat <<'EOF'
Usage: [mode]

mode:
    default: Run tests for all packages with caching.
    bench: Run benchmarks for all packages.
    cover: Run tests and report coverage for all packages.
    race: Run tests with race detection for the concurrency package.
EOF
}

help="${1:-none}"

if [[ "$help" == "-h" || "$help" == "--help" ]]; then
    display_help
    exit 0
fi

mode="${1:-default}"

selector=(
    "github.com/huangsam/go-trial/pkg/..."
    "github.com/huangsam/go-trial/internal/..."
)

case "$mode" in
    "default")
        go test "${selector[@]}" ;;
    "bench")
        go test -bench=. "${selector[@]}" ;;
    "cover")
        go test -cover "${selector[@]}" ;;
    "race")
        go test -race github.com/huangsam/go-trial/pkg/concurrency ;;
    *)
        echo "Invalid mode '$mode' detected" && exit 1 ;;
esac
