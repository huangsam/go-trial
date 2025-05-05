package concurrency_test

import (
	"testing"
	"time"

	"github.com/huangsam/go-trial/pkg/concurrency"
	"github.com/stretchr/testify/assert"
)

func TestRateLimitCounter(t *testing.T) {
	t.Run("RateLimitCounter with limit 5 and rate 100ms", func(t *testing.T) {
		limit := 5
		rate := 100 * time.Millisecond
		result := concurrency.RateLimitCounter(limit, rate)

		// Since the function runs for 1 second, we expect the result to be approximately
		// the number of increments possible within that time frame.
		expectedMin := 10 // 1 second / 100ms = 10 increments
		expectedMax := 15 // Allowing for some buffer due to timing variations

		assert.GreaterOrEqual(t, result, expectedMin)
		assert.LessOrEqual(t, result, expectedMax)
	})
}
