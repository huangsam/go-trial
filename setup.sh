#!/bin/bash
set -eu

brew install golangci-lint gotestsum mockery staticcheck

go install golang.org/x/tools/cmd/godoc@latest
