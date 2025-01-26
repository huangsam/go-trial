package concurrency

import (
	"context"
	"time"
)

// sumKey is the key for storing sumInfo instances into context.
type sumKey struct{}

// sumInfo stores the delay and factor values.
type sumInfo struct {
	delay  time.Duration
	factor int
}

// newContext returns a new context with the sumInfo value.
func newContext(ctx context.Context, sumInfoValue sumInfo) context.Context {
	return context.WithValue(ctx, sumKey{}, sumInfoValue)
}

// SumUntil sums numbers until the timeout is reached.
func SumUntil(timeout time.Duration, factor int) int {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	ctx = newContext(ctx, sumInfo{100 * time.Millisecond, factor})
	ch := make(chan int)
	sum := 0
	defer cancel()
	go func() {
		defer close(ch) // Close the channel when goroutine ends
		ctxSumInfo := ctx.Value(sumKey{}).(sumInfo)
		channelInput := 1
		for {
			select {
			case <-ctx.Done(): // Return when timeout is reached
				return
			case ch <- channelInput * ctxSumInfo.factor: // Send the number to the channel
				time.Sleep(ctxSumInfo.delay)
			}
			channelInput += 1
		}
	}()
	for {
		select {
		case <-ctx.Done(): // Return sum when timeout is reached
			return sum
		case num := <-ch: // Add the number received from the channel to the sum
			sum += num
		}
	}
}
