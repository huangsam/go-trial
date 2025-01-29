package main

import (
	"context"
	"os"

	"github.com/huangsam/go-trial/cmd/gotrial/tier1"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

// main is the entry point of the application.
func main() {
	cmd := &cli.Command{
		Usage:       "Try Go in action! 🔥",
		Description: "This binary has multiple commands to choose from.",
		Commands: [](*cli.Command){
			tier1.DemoCommand,
			tier1.ScrapeCommand,
			tier1.ServeCommand,
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Err(err).Msg("Cannot run command line")
	}
}
