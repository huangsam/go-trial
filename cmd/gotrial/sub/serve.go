package sub

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
		server := &http.Server{Addr: c.String("Port")}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Info().
				Str("method", r.Method).
				Str("host", r.Host).
				Str("path", r.URL.Path).
				Msg("Got request")

			fmt.Fprintf(w, "Hello, World!")
		})

		go func() {
			if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				log.Err(err).Msg("HTTP server error")
			}
			log.Info().Msg("Stop serving new connections")
		}()

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownRelease()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Error().Err(err).Msg("Shutdown error")
			return err
		}

		log.Info().Msg("Shutdown complete")
		return nil
	},
}
