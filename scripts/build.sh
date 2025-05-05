#!/bin/bash
set -eu

# default: Build the binary for local development.
# docker: Build the Docker image for the binary.
mode="${1:-default}"

root="$(git rev-parse --show-toplevel)"
cmd="gotrial"

case "$mode" in
    default)
        go build -o "$root/bin/$cmd" "github.com/huangsam/go-trial/cmd/$cmd" ;;
    docker)
        docker build "$root" -f "builds/docker/Dockerfile.$cmd" -t "huangsam/$cmd:latest" ;;
esac
