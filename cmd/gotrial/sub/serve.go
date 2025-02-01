package sub

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
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
			Name:  "port",
			Value: ":8080",
			Usage: "HTTP server port",
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.Use(gin.Recovery(), util.ZerologMiddleware())

		router.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Hello world")
		})

		router.GET("/health", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Health check")
		})

		server := &http.Server{Addr: c.String("port"), Handler: router}

		if err := util.GracefulShutdown(server); err != nil {
			log.Error().Err(err).Msg("Shutdown error")
			return err
		}

		log.Info().Msg("Shutdown success")
		return nil
	},
}
