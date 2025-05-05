package endpoint

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// LogUnaryRequest is a gRPC unary server interceptor that logs the request method.
var LogUnaryRequest grpc.UnaryServerInterceptor = func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	log.Info().Msgf("gRPC request: %s", info.FullMethod)
	return handler(ctx, req)
}

// LogStreamRequest is a gRPC stream server interceptor that logs the request method.
var LogStreamRequest grpc.StreamServerInterceptor = func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Info().Msgf("gRPC stream request: %s", info.FullMethod)
	return handler(srv, ss)
}
