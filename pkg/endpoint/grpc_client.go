package endpoint

import (
	"context"

	"github.com/huangsam/go-trial/internal/util"
	pb "github.com/huangsam/go-trial/pkg/endpoint/proto"
	"github.com/rs/zerolog/log"
)

// HelloValue is a value used in the EchoOnce and EchoStream methods.
const HelloValue = "Hello"

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
