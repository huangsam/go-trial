package main

import (
	"github.com/rs/zerolog/log"
	"io.huangsam/trial/pkg/abstraction"
	"io.huangsam/trial/pkg/basicintro"
)

// main is the entry point of the application.
func main() {
	log.Info().Msg(basicintro.GreetWorld())
	log.Info().Msg(basicintro.GreetName("Peter"))

	circle := abstraction.Circle{Radius: 6}
	logCircleSize(circle)
}

// logCircleSize logs the size of the given circle.
func logCircleSize(circle abstraction.Circle) {
	size := abstraction.Classify(circle)
	log.Info().Msgf("Circle size is %v", size)
}
