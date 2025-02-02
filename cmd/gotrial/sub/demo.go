package sub

import (
	"github.com/huangsam/go-trial/pkg/abstraction"
	"github.com/huangsam/go-trial/pkg/basicintro"
	"github.com/huangsam/go-trial/pkg/concurrency"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

// DemoCommand is a command to run a demo.
var DemoCommand *cli.Command = &cli.Command{
	Name:        "demo",
	Usage:       "Run demo with some pkg functions",
	Description: "This command runs functions from multiple packages.",
	Action: func(c *cli.Context) error {
		log.Debug().Msg(basicintro.GreetWorld())

		log.Info().Msg(basicintro.GreetName("Peter"))

		circle := abstraction.Circle{Radius: 6}
		size := abstraction.Classify(circle)
		log.Warn().Msgf("Circle size is %v", size)

		answers := concurrency.GetAnswersWithChannels()
		log.Error().Interface("answers", answers).Msg("Got answers with channels")

		return nil
	},
}
