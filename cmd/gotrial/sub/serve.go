package sub

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/huangsam/go-trial/internal/util"
	"github.com/huangsam/go-trial/pkg/abstraction"
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

		app.Get("/rectangle-size", func(c *fiber.Ctx) error {
			width, err := strconv.ParseFloat(c.Query("width", "1.0"), 64)
			if errors.Is(err, strconv.ErrSyntax) {
				return c.JSON(map[string]error{"error": err})
			}
			height, err := strconv.ParseFloat(c.Query("height", "1.0"), 64)
			if errors.Is(err, strconv.ErrSyntax) {
				return c.JSON(map[string]error{"error": err})
			}
			rect := abstraction.Rectangle{Width: width, Height: height}
			size := abstraction.Classify(&rect)
			payload := map[string]any{
				"area":      rect.Area(),
				"perimeter": rect.Perimeter(),
				"shape":     rect,
				"size":      size.String(),
			}
			return c.JSON(payload)
		})

		app.Get("/error", func(c *fiber.Ctx) error {
			return errors.New("What is going on with the world?")
		})

		return util.GracefulShutdown(app, c.String("addr"))
	},
}
