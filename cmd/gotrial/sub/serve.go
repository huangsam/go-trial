package sub

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/huangsam/go-trial/internal/util"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

// ServeCommand is a command to run an HTTP server.
var ServeCommand *cli.Command = &cli.Command{
	Name:        "serve",
	Usage:       "Run simple HTTP server",
	Description: "This command runs an HTTP server with one endpoint.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "addr",
			Value: ":8080",
			Usage: "HTTP address",
		},
		&cli.DurationFlag{
			Name:  "timeout",
			Value: 5 * time.Second,
			Usage: "HTTP read timeout",
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		app := fiber.New(fiber.Config{ReadTimeout: c.Duration("timeout")})
		app.Use(fiberzerolog.New(fiberzerolog.Config{Logger: &log.Logger}))

		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello world")
		})

		app.Get("/stack", func(c *fiber.Ctx) error {
			return c.JSON(c.App().Stack())
		})

		app.Get("/error", func(c *fiber.Ctx) error {
			return errors.New("What is going on with the world?")
		})

		return util.GracefulShutdown(app, c.String("addr"))
	},
}
