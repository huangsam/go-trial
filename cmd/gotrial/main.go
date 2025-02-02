package main

import (
	"os"

	"github.com/huangsam/go-trial/cmd/gotrial/sub"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

// main is the entry point of the application.
func main() {
	cmd := &cli.App{
		Usage:       "Try Go in action! 🔥",
		Description: "This binary has multiple commands to choose from.",
		Commands: [](*cli.Command){
			sub.DemoCommand,
			sub.ScrapeCommand,
			sub.ServeCommand,
		},
	}

	if err := cmd.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("Cannot run command line")
	}
}
