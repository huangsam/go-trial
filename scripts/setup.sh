#!/bin/bash
set -eu

# default: Install dependencies for local development.
# ci: Install dependencies for CI/CD.
mode="${1:-default}"

case "$mode" in
    default)
        brew install golangci-lint mockery protobuf
        go install golang.org/x/tools/cmd/godoc@latest
        go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
        ;;
    ci)
        PB_REL='https://github.com/protocolbuffers/protobuf/releases'
        PB_VER='31.0'
        PB_ARC='linux-x86_64'
        PB_DST="$HOME/.protoc"
        if [[ ! -d "$PB_DST" ]]; then
            curl -L "$PB_REL/download/v$PB_VER/protoc-$PB_VER-$PB_ARC.zip" -o protoc.zip
            unzip protoc.zip -d "$PB_DST"
        fi
        echo "Link $PB_DST/bin to PATH variable"
        export PATH="$PB_DST/bin:$PATH"
        go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
        ;;
    *)
        echo "Invalid mode '$mode' detected" && exit 1 ;;
esac

# Install dependencies
go mod download

# Generate protobuf files
go generate ./api
