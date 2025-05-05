package util

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

// LoggerBuilder is a builder for creating zerolog.Logger instances.
type LoggerBuilder struct {
	writer io.Writer
	level  zerolog.Level
}

// NewLoggerBuilder creates a new LoggerBuilder with default values.
func NewLoggerBuilder() *LoggerBuilder {
	return &LoggerBuilder{
		writer: os.Stdout,         // default writer
		level:  zerolog.InfoLevel, // default level
	}
}

// WithMode sets the output writer based on mode.
func (b *LoggerBuilder) WithMode(mode string) *LoggerBuilder {
	switch mode {
	case "console":
		b.writer = zerolog.NewConsoleWriter()
	case "stderr":
		b.writer = os.Stderr
	default:
		b.writer = os.Stdout
	}
	return b
}

// WithLevel sets the log level.
func (b *LoggerBuilder) WithLevel(level string) *LoggerBuilder {
	switch level {
	case "trace":
		b.level = zerolog.TraceLevel
	case "debug":
		b.level = zerolog.DebugLevel
	case "info":
		b.level = zerolog.InfoLevel
	case "warn":
		b.level = zerolog.WarnLevel
	case "error":
		b.level = zerolog.ErrorLevel
	default:
		b.level = zerolog.InfoLevel
	}
	return b
}

// Build creates and returns the configured logger.
func (b *LoggerBuilder) Build() zerolog.Logger {
	return zerolog.New(b.writer).
		With().Timestamp().Logger().
		Level(b.level)
}

// NewLogger creates a new logger instance based on LOG_MODE and LOG_LEVEL.
//
// The logger is set to output to os.Stdout by default, but can be changed
// based on the LOG_MODE environment variable.
//
// The log level is InfoLevel by default, but can be changed based on the
// LOG_LEVEL environment variable.
func NewLogger() zerolog.Logger {
	return NewLoggerBuilder().
		WithMode(os.Getenv("LOG_MODE")).
		WithLevel(os.Getenv("LOG_LEVEL")).
		Build()
}
