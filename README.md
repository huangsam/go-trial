# Go Trial

Curated learning designed for Software Engineers proficient in [Python] / [Java] / [C++]
and need to quickly ramp up their Go skills for work.

> The best time to start was yesterday. The second best time is now.

There's no better time to master Golang than now! 🔥

## Getting started

Here are the commands you need to know:

```shell
# Run build
./build.sh

# Run tests
./test.sh
```

### Documentation

To see the documentation via CLI:

```shell
# Show docs for package
go doc io.huangsam/trial/pkg/basicintro

# Show docs for package entity
go doc io.huangsam/trial/pkg/basicintro.Person
```

To see the documentation via web server:

```shell
# Install godoc
go install golang.org/x/tools/cmd/godoc@latest

# Run godoc server
godoc -http :8080
```

Then open up a browser on <http://localhost:8080>.

## Goals

Here is a checklist of things I want to cover:

- [x] [Intro to functions, structs](pkg/basicintro/)
- [x] [Data structures](pkg/datastructure/)
- [x] [Goroutines, channels, sync primitives](pkg/concurrency/)
- [x] [Interfaces, generics, mocking](pkg/abstraction/)

## Useful resources

### GitHub projects

- [avelino/awesome-go](https://github.com/avelino/awesome-go): A curated list of awesome Go frameworks, libraries, and software.
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout): Standard Go project layout.
- [golangci/golangci-lint](https://github.com/golangci/golangci-lint): Fast linters runner for Go.
- [uber-go/mock](https://github.com/uber-go/mock): A mocking framework for Go.
- [vektra/mockery](https://github.com/vektra/mockery): A mock code autogenerator for Golang.

### Online tutorials

- [A Tour of Go](https://go.dev/tour/list): An interactive tour of Go.
- [Go by Example](https://gobyexample.com): Hands-on introduction to Go using annotated example programs.

### Blog posts

- [Organizing Go code](https://go.dev/blog/organizing-go-code): A blog post on how to organize Go code.
- [Package names](https://go.dev/blog/package-names): A blog post on how to name Go packages effectively.

[Python]: https://github.com/huangsam/ultimate-python
[Java]: https://github.com/huangsam/java-trial
[C++]: https://github.com/huangsam/cpp-trial
