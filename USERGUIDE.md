# User guide

This content is for developers who are want to experiment with the repository at
a deeper level.

## Setup the developer environment

```shell
# Setup tools
bash scripts/setup.sh

# Run tests
bash scripts/test.sh

# Run linters
bash scripts/lint.sh
```

## See the docs

To see the codebase docs via CLI:

```shell
# Show docs for package
go doc github.com/huangsam/go-trial/pkg/basicintro

# Show docs for package entity
go doc github.com/huangsam/go-trial/pkg/basicintro.Person
```

To see the codebase docs via web server:

```shell
# Run godoc server
godoc -http :8080
```

Then open up a browser on <http://localhost:8080>.

<img src="images/godoc-server.png" alt="Godoc server" width="500px" />

## Generate mocks

To generate mocks for testing, add interfaces to `.mockery.yaml`.

Then run the following command:

```shell
# Run mockery to generate mocks
mockery
```

## Run with Docker

Here are the steps to run the `gotrial` CLI as a container:

```shell
# Run Docker build
bash scripts/build.sh docker

# Run Docker container
docker run --rm -it huangsam/gotrial:latest
```
