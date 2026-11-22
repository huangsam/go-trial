//go:generate protoc --go_out=. --go-grpc_out=. echo.proto
//go:generate protoc --go_out=. extended.proto

// Package api contains protobuf definitions for gRPC.
package api
