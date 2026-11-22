package sub

import (
	"context"
	"net/http"

	"github.com/huangsam/go-trial/internal/lesson"
	"github.com/huangsam/go-trial/internal/model"
	"github.com/huangsam/go-trial/lesson/practical"
	"github.com/urfave/cli/v3"
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
	Action: func(ctx context.Context, _ *cli.Command) error {
		mux := http.NewServeMux()
		mux.HandleFunc("GET /{$}", practical.HelloHandler)
		mux.HandleFunc("GET /error", practical.ErrorHandler)
		mux.HandleFunc("GET /rectangle-size", practical.RectangleSizeHandler)
		mux.HandleFunc("GET /circle-size", practical.CircleSizeHandler)

		acc := model.UserAccount{Username: ctx.Value("user").(string), Password: ctx.Value("pass").(string)}
		authMiddleware := lesson.BasicAuth(acc)
		mux.Handle("GET /secret", authMiddleware(http.HandlerFunc(practical.HelloHandler)))

		return lesson.RunServer(ctx.Value("addr").(string), lesson.ZeroLogger(mux))
	},
}
