package cmd_test

import (
	"testing"

	"github.com/huangsam/go-trial/internal/cmd"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestLoggerBuilder_WithLevel(t *testing.T) {
	builder := cmd.NewLoggerBuilder()
	logger := builder.WithLevel("debug").Build()
	assert.Equal(t, zerolog.DebugLevel, logger.GetLevel())
}

func TestLoggerBuilder_WithEmpty(t *testing.T) {
	builder := cmd.NewLoggerBuilder()
	logger := builder.WithLevel("").WithMode("").Build()
	assert.Equal(t, zerolog.InfoLevel, logger.GetLevel())
}
