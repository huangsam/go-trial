package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/huangsam/go-trial/pkg/abstraction"
	"github.com/huangsam/go-trial/pkg/basicintro"
	"github.com/huangsam/go-trial/pkg/concurrency"
	"github.com/urfave/cli/v3"
)

// demoCommand is a CLI command
var demoCommand *cli.Command = &cli.Command{
	Name:  "demo",
	Usage: "Do the demo thing",
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
