package util

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/huangsam/go-trial/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

// SetupBasicAuth sets up basic authentication middleware for an Echo web server.
//
// It uses the provided username and password to authenticate requests.
// Returns an Echo middleware function that checks the provided credentials against
// predefined admin credentials (AdminUser and AdminPass).
// If the credentials are valid, the request is allowed to proceed; otherwise, an error is returned.
func SetupBasicAuth(accounts ...model.UserAccount) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(u, p string, c echo.Context) (bool, error) {
		// Use a cache or a database in production
		for _, account := range accounts {
			if u == account.Username && p == account.Password {
				return true, nil
			}
		}
		return false, errors.New("invalid user credentials")
	})
}

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

// ZerologMiddleware emits a log for each incoming HTTP request.
var ZerologMiddleware echo.MiddlewareFunc = middleware.RequestLoggerWithConfig(
	middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("uri", v.URI).
				Int("status", v.Status).
				Dur("latency", v.Latency).
				Msg("Got request")

			return nil
		},
	})
