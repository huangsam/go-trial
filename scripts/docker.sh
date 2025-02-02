#!/bin/bash
set -eu

root="$(git rev-parse --show-toplevel)"

docker build "$root" -f builds/docker/Dockerfile -t huangsam/gotrial:latest
