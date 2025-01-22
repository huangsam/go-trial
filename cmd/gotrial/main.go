package main

import (
	"github.com/huangsam/go-trial/pkg/abstraction"
	"github.com/huangsam/go-trial/pkg/basicintro"
	"github.com/huangsam/go-trial/pkg/concurrency"
	"github.com/rs/zerolog/log"
)

// main is the entry point of the application.
func main() {
	log.Info().Msg(basicintro.GreetWorld())
	log.Info().Msg(basicintro.GreetName("Peter"))

	circle := abstraction.Circle{Radius: 6}
	logCircleSize(circle)

	concurrency.GetAnswersWithChannels()
}

// logCircleSize logs the size of the given circle.
func logCircleSize(circle abstraction.Circle) {
	size := abstraction.Classify(circle)
	log.Info().Msgf("Circle size is %v", size)
}
