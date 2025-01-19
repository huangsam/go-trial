package main

import (
	"github.com/rs/zerolog/log"
	"io.huangsam/trial/pkg/abstraction"
	basic "io.huangsam/trial/pkg/basicintro"
)

func main() {
	log.Info().Msg(basic.GreetWorld())
	log.Info().Msg(basic.GreetName("Peter"))

	circle := abstraction.Circle{Radius: 6}
	size := abstraction.Classify(circle)
	log.Info().Msgf("Circle size is %v", size)
}
