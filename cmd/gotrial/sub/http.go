package sub

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/huangsam/go-trial/internal/lesson"
	"github.com/huangsam/go-trial/internal/model"
	"github.com/huangsam/go-trial/pkg/endpoint"
	"github.com/urfave/cli/v2"
)

// HTTPCommand is a command to run an HTTP server.
var HTTPCommand *cli.Command = &cli.Command{
	Name:        "http",
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
		r := chi.NewRouter()
		r.Use(lesson.ZeroLogger)
		r.Use(middleware.Recoverer)

		acc := model.UserAccount{Username: c.String("user"), Password: c.String("pass")}
		authMiddleware := lesson.BasicAuth(acc)

		r.Get("/", endpoint.HelloHandler)
		r.Get("/error", endpoint.ErrorHandler)
		r.Get("/rectangle-size", endpoint.RectangleSizeHandler)
		r.Get("/circle-size", endpoint.CircleSizeHandler)
		r.With(authMiddleware).Get("/secret", endpoint.HelloHandler)

		return lesson.RunServer(c.String("addr"), r)
	},
}
