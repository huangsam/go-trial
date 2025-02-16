package sub

import (
	"net/http"
	"time"

	"github.com/huangsam/go-trial/internal/util"
	"github.com/huangsam/go-trial/pkg/endpoint"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		&cli.StringFlag{
			Name:  "user",
			Value: "admin",
			Usage: "Login username",
		},
		&cli.StringFlag{
			Name:  "pass",
			Value: "admin",
			Usage: "Login password",
		},
	},
	Action: func(c *cli.Context) error {
		e := echo.New()
		e.Use(util.ZerologMiddleware)
		e.Use(middleware.Recover())

		authMiddleware := util.SetupBasicAuth(c.String("user"), c.String("pass"))

		e.GET("/", endpoint.HelloHandler)
		e.GET("/error", endpoint.ErrorHandler)
		e.GET("/rectangle-size", endpoint.RectangleSizeHandler)
		e.GET("/secret", endpoint.HelloHandler, authMiddleware)

		srv := &http.Server{
			Addr:         c.String("addr"),
			Handler:      e,
			ReadTimeout:  c.Duration("rw"),
			WriteTimeout: c.Duration("rw"),
		}

		return util.GracefulShutdown(srv, c.Duration("shutdown"))
	},
}
