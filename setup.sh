#!/bin/bash
set -eu

mode="${1:-default}"

if [[ "$mode" == "default" ]]; then
    brew install golangci-lint gotestsum mockery staticcheck
    go install golang.org/x/tools/cmd/godoc@latest
elif [[ "$mode" == "ci" ]]; then
    go install gotest.tools/gotestsum@latest
fi
