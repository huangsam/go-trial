package concurrency_test

import (
	"context"
	"testing"
	"time"

	"github.com/huangsam/go-trial/lesson/concurrency"
	"github.com/stretchr/testify/assert"
)

func TestRateLimitCounter(t *testing.T) {
	limit := 5                         // Fixed burst limit
	duration := 400 * time.Millisecond // Fixed duration

	// Since the function runs for 1 second, we expect the result to be approximately
	// the number of increments possible within that time frame.

	t.Run("Rate50ms", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), duration)
		defer cancel()
		result := concurrency.RateLimitCounter(ctx, limit, 50*time.Millisecond)

		assert.GreaterOrEqual(t, result, 8)    // 400ms / 50ms = 8 increments
		assert.LessOrEqual(t, result, 8+limit) // allow for some buffer
	})

	t.Run("Rate100ms", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), duration)
		defer cancel()
		result := concurrency.RateLimitCounter(ctx, limit, 100*time.Millisecond)

		assert.GreaterOrEqual(t, result, 4)    // 400ms / 100ms = 4 increments
		assert.LessOrEqual(t, result, 4+limit) // allow for some buffer
	})

	t.Run("Rate200ms", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), duration)
		defer cancel()
		result := concurrency.RateLimitCounter(ctx, limit, 200*time.Millisecond)

		assert.GreaterOrEqual(t, result, 2)    // 400ms / 200ms = 2 increments
		assert.LessOrEqual(t, result, 2+limit) // allow for some buffer
	})
}
