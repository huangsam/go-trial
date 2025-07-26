package sub

import (
	"fmt"
	"time"

	"github.com/huangsam/go-trial/lesson/abstraction"
	"github.com/huangsam/go-trial/lesson/basicintro"
	"github.com/huangsam/go-trial/lesson/concurrency"
	"github.com/huangsam/go-trial/lesson/realworld"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/context"
)

// DemoCommand is a command to run a demo.
var DemoCommand *cli.Command = &cli.Command{
	Name:        "demo",
	Usage:       "Run demo with some functions",
	Description: "This command runs functions from multiple packages.",
	Action: func(*cli.Context) error {
		// basicintro
		log.Info().Msg(basicintro.GreetWorld())

		// abstraction
		circle := abstraction.Circle{Radius: 6}
		size := abstraction.Classify(circle)
		log.Info().Msgf("Circle size is %v", size)

		// concurrency
		answers := concurrency.GetAnswersWithChannels()
		log.Info().Interface("answers", answers).Msg("Got answers with channels")
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()
		count := concurrency.RateLimitCounter(ctx, 10, 50*time.Millisecond)
		log.Info().Int("count", count).Msg("Got rate limit count")

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
