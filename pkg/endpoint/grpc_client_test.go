package endpoint_test

import (
	"context"
	"testing"

	pb "github.com/huangsam/go-trial/api/protobuf"
	"github.com/huangsam/go-trial/pkg/endpoint"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	helloReq  = &pb.EchoRequest{Message: endpoint.HelloValue}
	helloResp = &pb.EchoResponse{Message: endpoint.HelloValue}
	doneReq   = &pb.EchoRequest{Message: endpoint.DoneValue}
	doneResp  = &pb.EchoResponse{Message: endpoint.DoneValue}
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
	client.On("EchoOnce", mock.Anything, helloReq).Return(helloResp, nil)

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
	client.On("EchoStream", mock.Anything, mock.Anything).Return(stream, nil)
	stream.On("CloseSend").Return(nil)
	stream.On("Recv").Return(helloResp, nil).Times(3) // 3 hello from server
	stream.On("Recv").Return(doneResp, nil).Once()    // 1 done from server
	stream.On("Send", helloReq).Once().Return(nil)    // 1 hello to server
	stream.On("Send", doneReq).Once().Return(nil)     // 1 done to server

	err := endpoint.EchoManyWithClient(context.Background(), client)
	require.NoError(t, err)
	client.AssertExpectations(t)
	stream.AssertExpectations(t)
}

func TestLogClientUnaryInfo(t *testing.T) {
	cc := new(grpc.ClientConn)
	invoker := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		return nil
	}
	ctx := context.Background()
	method := "/test.Service/TestMethod"

	err := endpoint.LogClientUnaryInfo(ctx, method, helloReq, helloResp, cc, invoker)
	require.NoError(t, err)
}

func TestLogClientStreamInfo(t *testing.T) {
	desc := &grpc.StreamDesc{StreamName: "TestStream"}
	cc := new(grpc.ClientConn)
	streamer := func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		return new(mockBidiStreamingClient), nil
	}
	ctx := context.Background()
	method := "/test.Service/TestMethod"

	clientStream, err := endpoint.LogClientStreamInfo(ctx, desc, cc, method, streamer)
	require.NoError(t, err)
	require.NotNil(t, clientStream)
}
