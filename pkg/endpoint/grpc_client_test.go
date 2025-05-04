package endpoint_test

import (
	"context"
	"testing"

	"github.com/huangsam/go-trial/pkg/endpoint"
	pb "github.com/huangsam/go-trial/pkg/endpoint/proto"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type mockEchoerClient struct {
	mock.Mock
}

func (m *mockEchoerClient) EchoOnce(ctx context.Context, req *pb.EchoRequest, opts ...grpc.CallOption) (*pb.EchoResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.EchoResponse), args.Error(1)
}

func (m *mockEchoerClient) EchoStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[pb.EchoRequest, pb.EchoResponse], error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(grpc.BidiStreamingClient[pb.EchoRequest, pb.EchoResponse]), args.Error(1)
}

func TestEchoOnceWithClient(t *testing.T) {
	client := new(mockEchoerClient)
	resp := &pb.EchoResponse{Message: endpoint.HelloValue}
	client.On("EchoOnce", mock.Anything, &pb.EchoRequest{Message: endpoint.HelloValue}).Return(resp, nil)

	err := endpoint.EchoOnceWithClient(context.Background(), client)
	require.NoError(t, err)
	client.AssertExpectations(t)
}

type mockBidiStreamingClient struct {
	mock.Mock
}

func (c *mockBidiStreamingClient) CloseSend() error {
	args := c.Called()
	return args.Error(0)
}

func (c *mockBidiStreamingClient) Context() context.Context {
	args := c.Called()
	return args.Get(0).(context.Context)
}

func (c *mockBidiStreamingClient) Header() (metadata.MD, error) {
	args := c.Called()
	return args.Get(0).(metadata.MD), args.Error(1)
}

func (c *mockBidiStreamingClient) Recv() (*pb.EchoResponse, error) {
	args := c.Called()
	return args.Get(0).(*pb.EchoResponse), args.Error(1)
}

func (c *mockBidiStreamingClient) RecvMsg(m any) error {
	args := c.Called(m)
	return args.Error(0)
}

func (c *mockBidiStreamingClient) Send(req *pb.EchoRequest) error {
	args := c.Called(req)
	return args.Error(0)
}

func (c *mockBidiStreamingClient) SendMsg(m any) error {
	args := c.Called(m)
	return args.Error(0)
}

func (c *mockBidiStreamingClient) Trailer() metadata.MD {
	args := c.Called()
	return args.Get(0).(metadata.MD)
}

func TestEchoManyWithClient(t *testing.T) {
	client := new(mockEchoerClient)
	stream := new(mockBidiStreamingClient)
	resp := &pb.EchoResponse{Message: endpoint.DoneValue}
	client.On("EchoStream", mock.Anything, mock.Anything).Return(stream, nil)
	stream.On("CloseSend").Return(nil)
	stream.On("Recv").Return(resp, nil)
	stream.On("Send", &pb.EchoRequest{Message: endpoint.HelloValue}).Return(nil)
	err := endpoint.EchoManyWithClient(context.Background(), client)
	require.NoError(t, err)
	client.AssertExpectations(t)
	stream.AssertExpectations(t)
}
