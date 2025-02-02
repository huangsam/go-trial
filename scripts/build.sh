#!/bin/bash
set -eu

root="$(git rev-parse --show-toplevel)"
cmd="gotrial"

go build -o "$root/bin/$cmd" "github.com/huangsam/go-trial/cmd/$cmd"
