#!/bin/bash
set -eu

mode="${1:-default}"

case "$mode" in
    "default")
        golangci-lint run ;;
    "fix")
        golangci-lint run --fix ;;
    "format")
        golangci-lint fmt ;;
    *)
        echo "Invalid mode '$mode' detected" && exit 1 ;;
esac
