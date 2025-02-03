package util

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog/log"
)

// GracefulShutdown shuts down the HTTP server gracefully.
func GracefulShutdown(server *http.Server, timeout time.Duration) error {
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("HTTP server error")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info().Msg("Stop HTTP server")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return server.Shutdown(ctx)
}
