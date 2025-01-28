#!/bin/bash
set -eu

declare -a brew_queue
declare -a go_queue

check_and_queue_brew () {
    [[ -n "$(which "$1")" ]] || brew_queue+=("$1")
}

check_and_queue_go () {
    [[ -n "$(which "$1")" ]] || go_queue+=("$2")
}

check_and_queue_brew "golangci-lint"
check_and_queue_brew "mockery"
check_and_queue_go "godoc" "golang.org/x/tools/cmd/godoc@latest"

# https://stackoverflow.com/a/15780028/2748860
[[ -z "${brew_queue[*]:-}" ]] || brew install "${brew_queue[@]}"
[[ -z "${go_queue[*]:-}" ]] || go install "${go_queue[@]}"
