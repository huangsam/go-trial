package endpoint

import (
	"context"
	"time"

	pb "github.com/huangsam/go-trial/pkg/endpoint/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// EchoerServer is a struct that implements the Echoer service defined in the proto file.
type EchoerServer struct {
	pb.UnimplementedEchoerServer
}

// EchoOnce implements the unary RPC method.
func (s *EchoerServer) EchoOnce(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: req.Message}, nil
}

// EchoStream implements the server streaming RPC method.
func (s *EchoerServer) EchoStream(stream grpc.BidiStreamingServer[pb.EchoRequest, pb.EchoResponse]) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		if req.Message == DoneValue {
			continue
		}
		for range 3 {
			resp := &pb.EchoResponse{Message: req.Message}
			if err := stream.Send(resp); err != nil {
				return err
			}
			time.Sleep(50 * time.Millisecond)
		}
		if err := stream.Send(&pb.EchoResponse{Message: DoneValue}); err != nil {
			return err
		}
	}
}

// LogServerUnaryInfo is a gRPC unary server interceptor that logs request info.
var LogServerUnaryInfo grpc.UnaryServerInterceptor = func(
	ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (resp any, err error) {
	log.Debug().Str("method", info.FullMethod).Msg("Handle gRPC unary request")
	return handler(ctx, req)
}

// LogServerStreamInfo is a gRPC stream server interceptor that logs request info.
var LogServerStreamInfo grpc.StreamServerInterceptor = func(
	srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler,
) error {
	log.Debug().Str("method", info.FullMethod).Msg("Handle gRPC stream request")
	return handler(srv, ss)
}
