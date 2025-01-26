package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/urfave/cli/v3"
)

// main is the entry point of the application.
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	cmd := &cli.Command{
		Commands: [](*cli.Command){
			demoCommand,
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		panic(err)
	}
}
