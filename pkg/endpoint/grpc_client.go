package endpoint

import (
	"context"

	"github.com/huangsam/go-trial/internal/util"
	pb "github.com/huangsam/go-trial/pkg/endpoint/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// EchoOnceWithClient demonstrates how to use a gRPC client to call the EchoOnce method.
func EchoOnceWithClient(ctx context.Context, client pb.EchoerClient) error {
	resp, err := client.EchoOnce(ctx, &pb.EchoRequest{Message: HelloValue})
	if err != nil {
		return err
	}
	log.Info().Msgf("Echo response: %s", resp.Message)
	return nil
}

// EchoManyWithClient demonstrates how to use a gRPC client to call the EchoStream method.
func EchoManyWithClient(ctx context.Context, client pb.EchoerClient) error {
	stream, err := client.EchoStream(ctx)
	if err != nil {
		return err
	}
	defer util.Dismiss(stream.CloseSend)
	err = stream.Send(&pb.EchoRequest{Message: HelloValue})
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Info().Msgf("Echo response: %s", resp.Message)
		if resp.Message == DoneValue {
			break
		}
	}
	if err := stream.Send(&pb.EchoRequest{Message: DoneValue}); err != nil {
		return err
	}
	return nil
}

// LogClientUnaryInfo is a gRPC unary client interceptor that logs request info.
func LogClientUnaryInfo(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Debug().Str("method", method).Msg("Send gRPC unary client request")
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

// LogClientStreamInfo is a gRPC stream client interceptor that logs request info.
func LogClientStreamInfo(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Debug().Str("method", method).Msg("Send gRPC stream client request")
	stream, err := streamer(ctx, desc, cc, method, opts...)
	return stream, err
}
