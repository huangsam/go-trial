#!/bin/bash
set -eu

mode="${1:-default}"

check_and_install_brew() {
    [[ -n "$(which "$1")" ]] && return
    brew install "$1"
}

check_and_install_go() {
    [[ -n "$(which "$1")" ]] && return
    go install "$2"
}

case "$mode" in
    "default")
        brew_list=("golangci-lint" "gotestsum" "mockery" "staticcheck")
        for brew in "${brew_list[@]}"; do
            check_and_install_brew "$brew"
        done
        check_and_install_go "godoc" "golang.org/x/tools/cmd/godoc@latest"
        ;;
    "ci")
        check_and_install_go "gotestsum" "gotest.tools/gotestsum@latest" ;;
    *)
        echo "Invalid mode '$mode' detected" && exit 1 ;;
esac
