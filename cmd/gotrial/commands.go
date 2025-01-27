package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/huangsam/go-trial/pkg/abstraction"
	"github.com/huangsam/go-trial/pkg/basicintro"
	"github.com/huangsam/go-trial/pkg/concurrency"
	"github.com/urfave/cli/v3"
)

// demoCommand is a command to run a demo.
var demoCommand *cli.Command = &cli.Command{
	Name:        "demo",
	Usage:       "Run demo with some pkg functions",
	Description: "This command runs functions from multiple packages.",
	Action: func(ctx context.Context, c *cli.Command) error {
		slog.Debug(basicintro.GreetWorld())

		slog.Info(basicintro.GreetName("Peter"))

		circle := abstraction.Circle{Radius: 6}
		size := abstraction.Classify(circle)
		slog.Warn(fmt.Sprintf("Circle size is %v", size))

		answers := concurrency.GetAnswersWithChannels()
		slog.Error("Retrieved answers with channels", "answers", answers)

		return nil
	},
}

// serverCommand is a command to run an HTTP server.
var serverCommand *cli.Command = &cli.Command{
	Name:        "server",
	Usage:       "Run server with HTTP responses",
	Description: "This command runs HTTP server with one endpoint.",
	Action: func(ctx context.Context, c *cli.Command) error {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, World!")
		})
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
		return nil
	},
}
