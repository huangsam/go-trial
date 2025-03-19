# Phase 1: Build binary
FROM golang:1.24-bookworm AS build

WORKDIR /app
COPY . /app

ENV CGO_ENABLED=0
ENV GOOS=linux

USER root

RUN bash scripts/build.sh

# Phase 2: Host binary
FROM alpine:latest

WORKDIR /
COPY --from=build /app/bin/gotrial /bin/gotrial

EXPOSE 8080

ENTRYPOINT ["/bin/gotrial"]
