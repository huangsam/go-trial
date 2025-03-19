package sub

import (
	"github.com/huangsam/go-trial/internal/model"
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
	Description: "This command runs an HTTP server.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "addr",
			Value: ":8080",
			Usage: "HTTP address",
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
		e.Use(middleware.Recover())
		e.Use(util.ZerologMiddleware)

		acc := model.UserAccount{Username: c.String("user"), Password: c.String("pass")}
		authMiddleware := util.SetupBasicAuth(acc)

		e.GET("/", endpoint.HelloHandler)
		e.GET("/error", endpoint.ErrorHandler)
		e.GET("/rectangle-size", endpoint.RectangleSizeHandler)
		e.GET("/secret", endpoint.HelloHandler, authMiddleware)

		return util.RunEcho(e, c.String("addr"))
	},
}
