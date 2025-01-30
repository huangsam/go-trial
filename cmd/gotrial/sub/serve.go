package sub

import (
	"context"
	"fmt"
	"net/http"

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
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Info().
				Str("method", r.Method).
				Str("host", r.Host).
				Str("path", r.URL.Path).
				Msg("Got request")

			fmt.Fprintf(w, "Hello, World!")
		})
		if err := http.ListenAndServe(c.String("port"), nil); err != nil {
			return err
		}
		return nil
	},
}
