package sub

import (
	"context"
	"fmt"
	"time"

	"github.com/huangsam/go-trial/lesson/basics"
	"github.com/huangsam/go-trial/lesson/concurrency"
	"github.com/huangsam/go-trial/lesson/practical"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

// DemoCommand is a command to run a demo.
var DemoCommand *cli.Command = &cli.Command{
	Name:        "demo",
	Usage:       "Run demo with some functions",
	Description: "This command runs functions from multiple packages.",
	Action: func(ctx context.Context, _ *cli.Command) error {
		// basics
		log.Info().Msg(basics.GreetWorld())

		circle := basics.Circle{Radius: 6}
		size := basics.Classify(circle)
		log.Info().Msgf("Circle size is %v", size)

		// concurrency
		answers := concurrency.GetAnswersWithChannels()
		log.Info().Interface("answers", answers).Msg("Got answers with channels")
		timeCtx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
		defer cancel()
		count := concurrency.RateLimitCounter(timeCtx, 10, 50*time.Millisecond)
		log.Info().Int("count", count).Msg("Got rate limit count")

		// practical
		car := practical.Car{Make: "Honda", Model: "Civic", ModelYear: 2025, WheelCount: 4, MileCount: 1234}
		content, err := practical.RenderCarInfo(&car)
		if err != nil {
			return err
		}
		fmt.Println(content)

		return nil
	},
}
