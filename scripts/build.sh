#!/bin/bash
set -eu

mode="${1:-default}"

root="$(git rev-parse --show-toplevel)"
cmd="gotrial"

case "$mode" in
    default)
        go build -o "$root/bin/$cmd" "github.com/huangsam/go-trial/cmd/$cmd" ;;
    docker)
        docker build "$root" -f "builds/docker/Dockerfile.$cmd" -t "huangsam/$cmd:latest" ;;
esac
