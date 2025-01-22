package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/huangsam/go-trial/pkg/abstraction"
	"github.com/huangsam/go-trial/pkg/basicintro"
	"github.com/huangsam/go-trial/pkg/concurrency"
)

// main is the entry point of the application.
func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	slog.Info(basicintro.GreetWorld())
	slog.Info(basicintro.GreetName("Peter"))

	circle := abstraction.Circle{Radius: 6}
	size := abstraction.Classify(circle)
	slog.Info(fmt.Sprintf("Circle size is %v", size))

	answers := concurrency.GetAnswersWithChannels()
	slog.Info("Retrieved answers with channels", "answers", answers)
}
