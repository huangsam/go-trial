package endpoint_test

import (
	"context"
	"io"
	"testing"

	"github.com/huangsam/go-trial/pkg/endpoint"
	pb "github.com/huangsam/go-trial/pkg/endpoint/proto"
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
