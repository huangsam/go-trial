package util

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// ZerologMiddleware writes a log entry for each HTTP request
func ZerologMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Info().
			Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Msg("Got request")
	}
}
