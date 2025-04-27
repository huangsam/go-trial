package sub

import (
	"fmt"

	"github.com/huangsam/go-trial/pkg/abstraction"
	"github.com/huangsam/go-trial/pkg/basicintro"
	"github.com/huangsam/go-trial/pkg/concurrency"
	"github.com/huangsam/go-trial/pkg/realworld"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

// DemoCommand is a command to run a demo.
var DemoCommand *cli.Command = &cli.Command{
	Name:        "demo",
	Usage:       "Run demo with some pkg functions",
	Description: "This command runs functions from multiple packages.",
	Action: func(c *cli.Context) error {
		// basicintro
		log.Info().Msg(basicintro.GreetWorld())

		// abstraction
		circle := abstraction.Circle{Radius: 6}
		size := abstraction.Classify(circle)
		log.Info().Msgf("Circle size is %v", size)

		// concurrency
		answers := concurrency.GetAnswersWithChannels()
		log.Info().Interface("answers", answers).Msg("Got answers with channels")

		// realworld
		car := realworld.Car{Make: "Honda", Model: "Civic", ModelYear: 2025, WheelCount: 4, MileCount: 1234}
		content, err := realworld.RenderCarInfo(&car)
		if err != nil {
			return err
		}
		fmt.Println(content)

		return nil
	},
}
