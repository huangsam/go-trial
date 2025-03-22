#!/bin/bash
set -eu

brew_packages=()
go_packages=()

install_brew_if_missing() {
    local package="$1"
    if ! command -v "$package" &> /dev/null; then
        brew_packages+=("$package")
    fi
}

install_go_if_missing() {
    local command="$1"
    local package="$2"
    if ! command -v "$command" &> /dev/null; then
        go_packages+=("$package")
    fi
}

# Define packages to install
install_brew_if_missing "golangci-lint"
install_brew_if_missing "mockery"
install_go_if_missing "godoc" "golang.org/x/tools/cmd/godoc@latest"

if [[ ${#brew_packages[@]} -gt 0 ]]; then
    echo "Install brew packages: ${brew_packages[*]}"
    brew install "${brew_packages[@]}"
fi

if [[ ${#go_packages[@]} -gt 0 ]]; then
    echo "Install go packages: ${go_packages[*]}"
    go install "${go_packages[@]}"
fi

echo "Installation checks completed."
