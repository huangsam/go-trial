package sub

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/huangsam/go-trial/internal/util"
	"github.com/huangsam/go-trial/pkg/endpoint"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
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
			Name:  "rw",
			Value: 5 * time.Second,
			Usage: "HTTP read/write timeout",
		},
		&cli.DurationFlag{
			Name:  "shutdown",
			Value: 10 * time.Second,
			Usage: "HTTP shutdown timeout",
		},
	},
	Action: func(c *cli.Context) error {
		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
			Output: &log.Logger,
			Formatter: func(params gin.LogFormatterParams) string {
				return fmt.Sprintf("%s - %s - %d - %v", params.Method, params.Request.URL, params.StatusCode, params.Latency)
			},
		}))
		router.Use(gin.Recovery())

		router.GET("/", endpoint.HelloHandler)
		router.GET("/error", endpoint.ErrorHandler)
		router.GET("/rectangle-size", endpoint.RectangleSizeHandler)

		srv := &http.Server{
			Addr:         c.String("addr"),
			Handler:      router,
			ReadTimeout:  c.Duration("rw"),
			WriteTimeout: c.Duration("rw"),
		}

		return util.GracefulShutdown(srv, c.Duration("shutdown"))
	},
}
