package concurrency

import (
	"context"
	"time"
)

// RateLimitCounter simulates a rate-limited counter.
//
// It creates a channel with a limit and a rate of replenishment.
// The counter will increment as long as the channel has space.
// The function will return the count of increments when the context is done.
func RateLimitCounter(ctx context.Context, limit int, rate time.Duration) int {
	// Fill the channel to its limit for bursty behavior
	limitChan := make(chan struct{}, limit)
	for range limit {
		limitChan <- struct{}{}
	}

	// Replenish the channel at the specified rate
	go func() {
		ticker := time.NewTicker(rate)
		for range ticker.C {
			limitChan <- struct{}{}
		}
	}()

	// Count the number of increments until the context times out
	result := 0
	for {
		select {
		case <-ctx.Done():
			return result
		case <-limitChan:
			result++
		}
	}
}
