package endpoint

import (
	"context"

	pb "github.com/huangsam/go-trial/pkg/endpoint/proto"
	"github.com/rs/zerolog/log"
)

// EchoOnceWithClient demonstrates how to use a gRPC client to call the EchoOnce method.
func EchoOnceWithClient(ctx context.Context, client pb.EchoerClient) error {
	resp, err := client.EchoOnce(ctx, &pb.EchoRequest{Message: "Hello"})
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
	err = stream.Send(&pb.EchoRequest{Message: "Hello"})
	if err != nil {
		return err
	}
	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		resp, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Info().Msgf("Echo response: %s", resp.Message)
	}
}
