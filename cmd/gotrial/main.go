package main

import (
	"context"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

// main is the entry point of the application.
func main() {
	cmd := &cli.Command{
		Usage:       "Try Go in action! 🔥",
		Description: "This binary has multiple commands to choose from.",
		Commands: [](*cli.Command){
			demoCommand,
			serverCommand,
			scrapeCommand,
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Err(err).Msg("Cannot run command line")
	}
}
