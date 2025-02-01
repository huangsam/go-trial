package sub

import (
	"context"
	"net/http"

	"github.com/huangsam/go-trial/internal/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
			Name:  "port",
			Value: ":8080",
			Usage: "HTTP server port",
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		e := echo.New()
		e.Use(middleware.Recover(), util.ZerologMiddleware())

		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello world")
		})

		server := &http.Server{Addr: c.String("port"), Handler: e}

		if err := util.GracefulShutdown(server); err != nil {
			log.Error().Err(err).Msg("Shutdown error")
			return err
		}

		log.Info().Msg("Shutdown success")
		return nil
	},
}
