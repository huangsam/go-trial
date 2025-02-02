# Go Trial

[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/huangsam/go-trial/ci.yml)](https://github.com/huangsam/go-trial/actions)

Curated learning designed for Software Engineers proficient in [Python] / [Java] / [C++]
and need to quickly ramp up their Go skills for work.

> The best time to start was yesterday. The second best time is now.

There's no better time to master Golang than now! 🔥

<img src="images/gopher.jpeg" alt="Gopher" width="250px" />

## Getting started

Here are the commands you need to know:

```shell
# Setup tools
bash scripts/setup.sh

# Run build
bash scripts/build.sh

# Run tests
bash scripts/test.sh

# Run linters
bash scripts/lint.sh
```

### Run the binary

After the build runs successfully, the `gotrial` binary should exist at the root
directory of this repo. Here are some commands to get you started:

```shell
# Get help for global command
./gotrial help

# Get help for subcommand
./gotrial help demo

# Run a command
./gotrial demo
```

### See the docs

To see the documentation via CLI:

```shell
# Show docs for package
go doc github.com/huangsam/go-trial/pkg/basicintro

# Show docs for package entity
go doc github.com/huangsam/go-trial/pkg/basicintro.Person
```

To see the documentation via web server:

```shell
# Run godoc server
godoc -http :8080
```

Then open up a browser on <http://localhost:8080>.

<img src="images/godoc-server.png" alt="Godoc server" width="500px" />

### Generate mocks

To generate mocks for testing, add interfaces to `.mockery.yaml`.

Then run the following command:

```shell
# Run mockery to generate mocks
mockery
```

## Goals

Here is a checklist of things I want to cover:

### Core concepts

- [x] [Intro to functions, structs](pkg/basicintro/)
- [x] [Data structures](pkg/datastructure/)
- [x] [Goroutines, channels, sync primitives](pkg/concurrency/)
- [x] [Interfaces, generics, mocking](pkg/abstraction/)

### Applied concepts

- [x] [Config, I/O, time](pkg/realworld/)
- [x] [HTTP endpoints](pkg/endpoint/)

## Useful resources

### GitHub projects

#### Content

- [avelino/awesome-go](https://github.com/avelino/awesome-go): Awesome Go frameworks, libraries, software.
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout): Standard Go project layout.

#### Tools

- [golangci/golangci-lint](https://github.com/golangci/golangci-lint): Fast linters runner.
- [knadh/koanf](https://github.com/knadh/koanf): Simple configuration management.
- [rs/zerolog](https://github.com/rs/zerolog): Efficient JSON logger.
- [urfave/cli](https://github.com/urfave/cli): Simple, fast, fun CLIs.
- [vektra/mockery](https://github.com/vektra/mockery): Mock code autogenerator.

### Online tutorials

- [A Tour of Go](https://go.dev/tour/list): An interactive tour of Go.
- [Go by Example](https://gobyexample.com): Annotated programs in Go.

### Blog posts

- [Organizing Go code](https://go.dev/blog/organizing-go-code): General codebase conventions.
- [Package names](https://go.dev/blog/package-names): Package naming conventions.
- [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines): Pipelines with multiple stages.
- [Go Concurrency Patterns: Context](https://go.dev/blog/context): Context for downstream logic.
- [Structured Logging with slog](https://go.dev/blog/slog): Structured logging using slog.
- [4 Mocking Approaches for Go](https://www.twilio.com/en-us/blog/4-mocking-approaches-go): Different mocking techniques.
- [Proper HTTP shutdown in Go](https://dev.to/mokiat/proper-http-shutdown-in-go-3fji): Graceful HTTP shutdown.

[Python]: https://github.com/huangsam/ultimate-python
[Java]: https://github.com/huangsam/java-trial
[C++]: https://github.com/huangsam/cpp-trial
