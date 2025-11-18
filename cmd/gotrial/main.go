// Package main is the entry point for the go-trial application.
package main

import (
	"context"
	"os"

	"github.com/huangsam/go-trial/cmd/gotrial/sub"
	"github.com/huangsam/go-trial/internal/cmd"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

// init initializes the logger for the application.
func init() {
	log.Logger = cmd.NewLogger()
}

// main is the entry point of the application.
func main() {
	cmd := &cli.Command{
		Usage:       "Try Go in action! 🔥",
		Description: "This command supports ad-hoc runs and client/server interactions.",
		Commands: [](*cli.Command){
			sub.DemoCommand,
			sub.HTTPCommand,
			sub.GrpcCommand,
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal().Err(err).Msg("Cannot run command line")
	}
}
