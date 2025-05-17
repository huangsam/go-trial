package endpoint_test

import (
	"context"
	"io"
	"testing"

	pb "github.com/huangsam/go-trial/api/protobuf"
	"github.com/huangsam/go-trial/pkg/endpoint"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestEchoOnce(t *testing.T) {
	server := &endpoint.EchoerServer{}
	req := &pb.EchoRequest{Message: "Test Message"}

	resp, err := server.EchoOnce(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, "Test Message", resp.Message)
}

type mockStream struct {
	grpc.ServerStream
	recvMessages []*pb.EchoRequest
	sendMessages []*pb.EchoResponse
	recvIndex    int
}

func (m *mockStream) Recv() (*pb.EchoRequest, error) {
	if m.recvIndex >= len(m.recvMessages) {
		return nil, io.EOF
	}
	msg := m.recvMessages[m.recvIndex]
	m.recvIndex++
	return msg, nil
}

func (m *mockStream) Send(resp *pb.EchoResponse) error {
	m.sendMessages = append(m.sendMessages, resp)
	return nil
}

func TestEchoStream(t *testing.T) {
	server := &endpoint.EchoerServer{}
	mock := &mockStream{
		recvMessages: []*pb.EchoRequest{
			{Message: "Hello"},
			{Message: "World"},
		},
		sendMessages: []*pb.EchoResponse{},
	}

	err := server.EchoStream(mock)
	require.ErrorIs(t, err, io.EOF)

	expectedMessages := []*pb.EchoResponse{
		{Message: "Hello"},
		{Message: "Hello"},
		{Message: "Hello"},
		{Message: "Done"},
		{Message: "World"},
		{Message: "World"},
		{Message: "World"},
		{Message: "Done"},
	}

	assert.Equal(t, expectedMessages, mock.sendMessages)
}

func TestLogServerUnaryInfo(t *testing.T) {
	expectedMessage := "Test Message"
	mockHandler := func(ctx context.Context, req any) (any, error) {
		return &pb.EchoResponse{Message: expectedMessage}, nil
	}
	info := &grpc.UnaryServerInfo{
		FullMethod: "/pkg.endpoint.Echoer/EchoOnce",
	}
	req := &pb.EchoRequest{Message: "Test Message"}
	ctx := context.Background()

	resp, err := endpoint.LogServerUnaryInfo(ctx, req, info, mockHandler)
	require.NoError(t, err)
	echoResp, ok := resp.(*pb.EchoResponse)
	require.True(t, ok)
	assert.Equal(t, expectedMessage, echoResp.Message)
}

func TestLogServerStreamInfo(t *testing.T) {
	mockHandler := func(any, grpc.ServerStream) error {
		return nil
	}
	info := &grpc.StreamServerInfo{
		FullMethod:     "/pkg.endpoint.Echoer/EchoStream",
		IsClientStream: true,
		IsServerStream: true,
	}
	mockStream := &mockStream{}

	err := endpoint.LogServerStreamInfo(nil, mockStream, info, mockHandler)
	require.NoError(t, err)
}
