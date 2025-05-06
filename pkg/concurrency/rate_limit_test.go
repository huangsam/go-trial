package concurrency_test

import (
	"testing"
	"time"

	"github.com/huangsam/go-trial/pkg/concurrency"
	"github.com/stretchr/testify/assert"
)

func TestRateLimitCounter(t *testing.T) {
	limit := 5 // Burst of up to 5 increments

	// Since the function runs for 1 second, we expect the result to be approximately
	// the number of increments possible within that time frame.

	t.Run("Limit5Rate50ms", func(t *testing.T) {
		rate := 50 * time.Millisecond
		result := concurrency.RateLimitCounter(limit, rate)

		expectedMin := 20 // 1 second / 50ms = 20 increments
		expectedMax := 25 // Allow for some buffer due to timing variations
		assert.GreaterOrEqual(t, result, expectedMin)
		assert.LessOrEqual(t, result, expectedMax)
	})

	t.Run("Limit5Rate100ms", func(t *testing.T) {
		rate := 100 * time.Millisecond
		result := concurrency.RateLimitCounter(limit, rate)

		expectedMin := 10 // 1 second / 100ms = 10 increments
		expectedMax := 15 // Allow for some buffer due to timing variations
		assert.GreaterOrEqual(t, result, expectedMin)
		assert.LessOrEqual(t, result, expectedMax)
	})

	t.Run("Limit5Rate200ms", func(t *testing.T) {
		rate := 200 * time.Millisecond
		result := concurrency.RateLimitCounter(limit, rate)

		expectedMin := 5  // 1 second / 200ms = 5 increments
		expectedMax := 10 // Allow for some buffer due to timing variations
		assert.GreaterOrEqual(t, result, expectedMin)
		assert.LessOrEqual(t, result, expectedMax)
	})
}
