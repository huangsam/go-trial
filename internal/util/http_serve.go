package util

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

// ZerologMiddleware emits a log for each incoming HTTP request.
var ZerologMiddleware echo.MiddlewareFunc = middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	LogURI:    true,
	LogStatus: true,
	LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
		log.Info().
			Str("uri", v.URI).
			Int("status", v.Status).
			Msg("Got request")

		return nil
	},
})
