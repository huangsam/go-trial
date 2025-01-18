# Go Trial

There's no better time to master Golang than now! 🔥

Previous trials include: [Python], [Java], [C++]

## Getting started

Here are the commands you need to know:

```shell
# Run build
./build.sh

# Run tests
./test.sh

# Run benchmarks
./test.sh bench
```

### Documentation

To see the documentation over a local web server:

```shell
# Install godoc
go install golang.org/x/tools/cmd/godoc@latest

# Run godoc server
godoc -http :8080
```

## Goals

Here is a checklist of things I want to cover:

- [x] [Intro to functions, structs](pkg/basicintro/)
- [x] [Data structures](pkg/datastructure/)
- [ ] [Goroutines, channels, sync primitives](pkg/concurrency/)
- [ ] [Interfaces, generics, mocking](pkg/abstraction/)

## Useful resources

- <https://github.com/avelino/awesome-go>
- <https://github.com/golang-standards/project-layout>
- <https://github.com/golangci/golangci-lint>
- <https://github.com/uber-go/mock>
- <https://go.dev/tour/list>
- <https://gobyexample.com>
- <https://go.dev/talks/2014/organizeio.slide#9>
- <https://go.dev/blog/organizing-go-code>
- <https://go.dev/blog/package-names>

[Python]: https://github.com/huangsam/ultimate-python
[Java]: https://github.com/huangsam/java-trial
[C++]: https://github.com/huangsam/cpp-trial
