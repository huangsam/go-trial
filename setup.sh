#!/bin/bash
set -eu

mode="${1:-default}"

declare -a brew_queue
declare -a go_queue

check_and_add_brew() {
    [[ -n "$(which "$1")" ]] && return
    brew_queue+=("$1")
}

check_and_add_go() {
    [[ -n "$(which "$1")" ]] && return
    go_queue+=("$2")
}

case "$mode" in
    "default")
        brew_list=("golangci-lint" "gotestsum" "mockery")
        for brew in "${brew_list[@]}"; do
            check_and_add_brew "$brew"
        done
        check_and_add_go "godoc" "golang.org/x/tools/cmd/godoc@latest"
        ;;
    "ci")
        check_and_add_go "gotestsum" "gotest.tools/gotestsum@latest" ;;
    *)
        echo "Invalid mode '$mode' detected" && exit 1 ;;
esac

# https://stackoverflow.com/a/15780028/2748860
[[ -n "${brew_queue[*]:-}" ]] && brew install "${brew_queue[@]}"
[[ -n "${go_queue[*]:-}" ]] && go install "${go_queue[@]}"
