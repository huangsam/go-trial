package sub

import (
	"context"
	"net"
	"time"

	"github.com/huangsam/go-trial/internal/util"
	"github.com/huangsam/go-trial/pkg/endpoint"
	pb "github.com/huangsam/go-trial/pkg/endpoint/proto"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcCommand is a catch-all command for gRPC.
var GrpcCommand *cli.Command = &cli.Command{
	Name:        "grpc",
	Usage:       "Play with gRPC",
	Description: "This command supports gRPC server and client interactions.",
	Subcommands: []*cli.Command{
		GrpcServeCommand,      // run in first terminal
		GrpcEchoOnceCommand,   // run in second terminal
		GrpcEchoStreamCommand, // run in second terminal
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "addr",
			Value: ":50051", // default gRPC address
			Usage: "gRPC address",
		},
	},
}

// GrpcServeCommand is a command to run a gRPC server.
var GrpcServeCommand *cli.Command = &cli.Command{
	Name:        "serve",
	Usage:       "Run gRPC server",
	Description: "This command runs the Echoer gRPC server.",
	Action: func(c *cli.Context) error {
		addr := c.String("addr")
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to listen")
		}
		server := grpc.NewServer(
			grpc.UnaryInterceptor(endpoint.LogServerUnaryInfo),
			grpc.StreamInterceptor(endpoint.LogServerStreamInfo),
		)
		pb.RegisterEchoerServer(server, &endpoint.EchoerServer{})
		log.Info().Msgf("gRPC server listening on %s", addr)
		return server.Serve(lis)
	},
}

// requestTimeout is the timeout for gRPC requests.
var requestTimeout = 5 * time.Second

// GrpcEchoOnceCommand is a command to run a gRPC client.
var GrpcEchoOnceCommand *cli.Command = &cli.Command{
	Name:        "echo-once",
	Usage:       "Call Echoer server once",
	Description: "This command calls the Echoer gRPC server once.",
	Action: func(c *cli.Context) error {
		addr := c.String("addr")
		conn, err := grpc.NewClient(addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithUnaryInterceptor(endpoint.LogClientUnaryInfo),
		)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to connect to server")
		}
		defer util.Dismiss(conn.Close)
		client := pb.NewEchoerClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()
		return endpoint.EchoOnceWithClient(ctx, client)
	},
}

// GrpcEchoStreamCommand is a command to run a gRPC client.
var GrpcEchoStreamCommand *cli.Command = &cli.Command{
	Name:        "echo-stream",
	Usage:       "Call Echoer server with stream",
	Description: "This command calls the Echoer gRPC server with a stream.",
	Action: func(c *cli.Context) error {
		addr := c.String("addr")
		conn, err := grpc.NewClient(addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithStreamInterceptor(endpoint.LogClientStreamInfo),
		)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to connect to server")
		}
		defer util.Dismiss(conn.Close)
		client := pb.NewEchoerClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()
		return endpoint.EchoManyWithClient(ctx, client)
	},
}
