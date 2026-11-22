# --- Configuration ---

GO             ?= go
GOLANGCI_LINT  ?= golangci-lint
BINARY_NAME    ?= gotrial
BIN_DIR        ?= bin
MAIN_PACKAGE   ?= ./cmd/gotrial
TEST_PACKAGES  ?= ./lesson/... ./internal/...

PROTOC_VERSION ?= 32.0
PROTOC_ARCHIVE ?= linux-x86_64
PROTOC_DIR     ?= $(HOME)/.protoc

.DEFAULT_GOAL := build

.PHONY: all build docker-build clean setup setup-ci setup-go test bench coverage coverage-html test-race lint lint-fix format help

all: build

build:
	@mkdir -p "$(BIN_DIR)"
	$(GO) build -o "$(BIN_DIR)/$(BINARY_NAME)" "$(MAIN_PACKAGE)"

docker-build:
	docker build . -f "builds/docker/Dockerfile.$(BINARY_NAME)" -t "huangsam/$(BINARY_NAME):latest"

clean:
	rm -f "$(BIN_DIR)/$(BINARY_NAME)" coverage.out coverage.html

setup:
	brew install golangci-lint mockery protobuf
	$(GO) install golang.org/x/tools/cmd/godoc@latest
	$(MAKE) setup-go

setup-ci:
	@if [ ! -x "$(PROTOC_DIR)/bin/protoc" ]; then \
		mkdir -p "$(PROTOC_DIR)"; \
		curl -fsSL "https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-$(PROTOC_ARCHIVE).zip" -o "$(PROTOC_DIR)/protoc.zip"; \
		unzip -o "$(PROTOC_DIR)/protoc.zip" -d "$(PROTOC_DIR)"; \
	fi
	PATH="$(PROTOC_DIR)/bin:$$PATH" $(MAKE) setup-go

setup-go:
	$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	$(GO) mod download
	$(GO) generate ./api

test:
	$(GO) test $(TEST_PACKAGES)

bench:
	$(GO) test -bench=. $(TEST_PACKAGES)

coverage:
	$(GO) test -cover $(TEST_PACKAGES)

coverage-html:
	$(GO) test -cover -coverprofile=coverage.out $(TEST_PACKAGES)
	$(GO) tool cover -html=coverage.out -o coverage.html

test-race:
	$(GO) test -race ./lesson/concurrency

lint:
	$(GOLANGCI_LINT) run

lint-fix:
	$(GOLANGCI_LINT) run --fix

format:
	$(GOLANGCI_LINT) fmt

help:
	@echo "Go Trial development targets"
	@echo "  make (or make build)  Build the CLI into $(BIN_DIR)/$(BINARY_NAME)"
	@echo "  make docker-build     Build the CLI Docker image"
	@echo "  make setup            Install local development tools and generate code"
	@echo "  make setup-ci         Install CI generators and generate code"
	@echo "  make test             Run tests in $(TEST_PACKAGES)"
	@echo "  make bench            Run benchmarks"
	@echo "  make coverage         Run tests with coverage"
	@echo "  make coverage-html    Generate coverage.out and coverage.html"
	@echo "  make test-race        Run concurrency tests with the race detector"
	@echo "  make lint             Run golangci-lint"
	@echo "  make lint-fix         Run golangci-lint with automatic fixes"
	@echo "  make format           Format code with golangci-lint"
	@echo "  make clean            Remove local build and coverage artifacts"
