# Go Trial

[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/huangsam/go-trial/ci.yml)](https://github.com/huangsam/go-trial/actions)
[![License](https://img.shields.io/github/license/huangsam/go-trial)](https://github.com/huangsam/go-trial/blob/main/LICENSE)

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

To learn more, please consult the [user guide](USERGUIDE.md).

## Learning path

This learning path is designed to take you from Go beginner to proficient
in a matter of weeks. Each section contains practical examples to read
and experiment with.

### Core concepts

- [Functions, structs, pointers](pkg/basicintro/)
- [Common data structures](pkg/datastructure/)
- [Goroutines, channels, synchronization](pkg/concurrency/)
- [Interfaces, generics, mocking](pkg/abstraction/)

### Practical concepts

- [Files, I/O, time](pkg/realworld/)
- [HTTP endpoints](pkg/endpoint/)

### Apply your knowledge!

Ready to showcase your Go skills? Dive into a side project and create something amazing!
Check out [Namigo](https://github.com/huangsam/namigo) for inspiration. 🚀

## Useful resources

### GitHub projects

- [avelino/awesome-go](https://github.com/avelino/awesome-go): Awesome Go frameworks, libraries, software.
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout): Standard Go project layout.

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
