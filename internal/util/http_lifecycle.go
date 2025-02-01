package util

import (
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// GracefulShutdown shuts down the HTTP server gracefully
func GracefulShutdown(app *fiber.App, addr string) error {
	go func() {
		if err := app.Listen(addr); err != nil {
			log.Error().Msg("HTTP server error")
		}
		log.Info().Msg("Stop accepting connections")
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	return app.Shutdown()
}
