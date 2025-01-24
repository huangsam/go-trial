package concurrency

import (
	"context"
	"time"
)

// SumOneUntil returns the sum of 1s received until the timeout is reached.
func SumOneUntil(timeout time.Duration) int {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	ch := make(chan int)
	sum := 0
	defer cancel()
	go func() {
		defer close(ch) // Close the channel when the function returns
		for {
			select {
			case <-ctx.Done(): // Return when timeout is reached
				return
			case ch <- 1: // Send 1 to the channel every 100ms
				time.Sleep(100 * time.Millisecond)
			}
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
