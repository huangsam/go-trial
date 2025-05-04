#!/bin/bash
set -eu

mode="${1:-default}"

case "$mode" in
    default)
        brew install golangci-lint mockery protobuf
        go install golang.org/x/tools/cmd/godoc@latest
        go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
        ;;
    ci)
        curl -sSL https://github.com/protocolbuffers/protobuf/releases/download/v30.2/protoc-30.2-linux-x86_64.zip -o protoc.zip
        unzip protoc.zip -d "$HOME/.protoc"
        chmod +x "$HOME/.protoc/bin/protoc"
        go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
        ;;
    *)
        echo "Invalid mode '$mode' detected" && exit 1 ;;
esac

# Install dependencies
go mod download

# Generate protobuf files
protoc --go_out=. --go-grpc_out=. pkg/endpoint/proto/echo.proto
