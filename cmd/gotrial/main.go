package main

import (
	"github.com/rs/zerolog/log"
	basic "io.huangsam/trial/pkg/basicintro"
)

func main() {
	log.Info().Msg(basic.GreetWorld())
	log.Info().Msg(basic.GreetName("Peter"))
}
