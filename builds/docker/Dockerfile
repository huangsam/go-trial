FROM golang:1.24-bookworm AS build

COPY . /app

WORKDIR /app

RUN bash scripts/build.sh

FROM debian:bookworm

COPY --from=build /app/bin/gotrial /bin/gotrial

ENTRYPOINT ["/bin/gotrial"]
