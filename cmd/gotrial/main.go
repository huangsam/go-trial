// Package main is the entry point for the go-trial application.
package main

import (
	"os"

	"github.com/huangsam/go-trial/cmd/gotrial/sub"
	"github.com/huangsam/go-trial/internal/util"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

// init initializes the logger for the application.
func init() {
	log.Logger = util.NewLogger()
}

// main is the entry point of the application.
func main() {
	cmd := &cli.App{
		Usage:       "Try Go in action! 🔥",
		Description: "This command supports ad-hoc runs and client/server interactions.",
		Commands: [](*cli.Command){
			sub.DemoCommand,
			sub.HTTPCommand,
			sub.GrpcCommand,
		},
	}

	if err := cmd.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("Cannot run command line")
	}
}
