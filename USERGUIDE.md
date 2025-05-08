# User guide

This content is for developers who want to experiment with the repository at
a deeper level.

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

# Run Docker container as a CLI
docker run --rm -it huangsam/gotrial:latest

# Run Docker container as an HTTP server
docker run --rm -it -p 8080:8080 huangsam/gotrial:latest http

# Run Docker container as a gRPC server
docker run --rm -it -p 50051:50051 huangsam/gotrial:latest grpc serve
```
