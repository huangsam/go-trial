package endpoint

import (
	"context"
	"time"

	pb "github.com/huangsam/go-trial/pkg/endpoint/proto"
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
		for range 4 {
			resp := &pb.EchoResponse{Message: req.Message}
			if err := stream.Send(resp); err != nil {
				return err
			}
			time.Sleep(50 * time.Millisecond)
		}
	}
}
