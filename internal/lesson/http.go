package lesson

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/huangsam/go-trial/internal/model"
	"github.com/rs/zerolog/log"
)

// RunServer runs an HTTP server until an interrupt shuts it down.
//
// Specifically, it waits for a SIGINT or SIGTERM signal to be received.
// When such a signal is received, it gracefully shuts down the server
// within a reasonable timeout period.
func RunServer(addr string, handler http.Handler) error {
	alog := log.With().Str("addr", addr).Logger()
	alog.Info().Msg("Start HTTP server")

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			alog.Fatal().Err(err).Msg("HTTP server error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	alog.Info().Msg("Stop HTTP server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
}

// ZeroLogger emits a log for each incoming HTTP request.
//
// It logs the request URI, status code, and latency.
func ZeroLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)
		latency := time.Since(start)

		log.Debug().
			Str("uri", r.RequestURI).
			Int("status", ww.Status()).
			Dur("latency", latency).
			Msg("Got request")
	})
}

// BasicAuth sets up basic authentication middleware.
//
// Returns an HTTP middleware function that checks the provided credentials against
// the provided accounts. If the credentials are valid, the request is allowed
// to proceed; otherwise, an error is returned.
func BasicAuth(accounts ...model.UserAccount) func(http.Handler) http.Handler {
	accountSet := map[string]string{}
	for _, account := range accounts {
		accountSet[account.Username] = account.Password
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()
			if !ok {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`) // Prompt for credentials
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			storedPass, exists := accountSet[user]
			if !exists || storedPass != pass {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
