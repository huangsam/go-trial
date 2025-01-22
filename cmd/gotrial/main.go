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
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info(basicintro.GreetWorld())
	logger.Info(basicintro.GreetName("Peter"))

	circle := abstraction.Circle{Radius: 6}
	size := abstraction.Classify(circle)
	logger.Info(fmt.Sprintf("Circle size is %v", size))

	answers := concurrency.GetAnswersWithChannels()
	logger.Info("Retrieved answers with channels", "answers", answers)
}
