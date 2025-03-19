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

const shutdownWait = 10 * time.Second

var (
	ErrBadPassword = errors.New("bad password")
	ErrMissingUser = errors.New("missing user")
)

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

// SetupBasicAuth sets up basic authentication middleware for an Echo web server.
//
// It uses the provided username and password to authenticate requests.
// Returns an Echo middleware function that checks the provided credentials against
// predefined admin credentials (AdminUser and AdminPass).
// If the credentials are valid, the request is allowed to proceed; otherwise, an error is returned.
func SetupBasicAuth(accounts ...model.UserAccount) echo.MiddlewareFunc {
	accountSet := map[string]string{}
	for _, account := range accounts {
		accountSet[account.Username] = account.Password
	}
	return middleware.BasicAuth(func(u, p string, c echo.Context) (bool, error) {
		pass, ok := accountSet[u]
		if !ok {
			return false, ErrMissingUser
		}
		if pass != p {
			return false, ErrBadPassword
		}
		return true, nil
	})
}

// RunEcho runs an Echo server until an interrupt shuts it down.
func RunEcho(echo *echo.Echo, addr string) error {
	go func() {
		if err := echo.Start(addr); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("HTTP server error")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info().Msg("Stop HTTP server")
	ctx, cancel := context.WithTimeout(context.Background(), shutdownWait)
	defer cancel()
	return echo.Shutdown(ctx)
}
