package concurrency

import (
	"context"
	"time"
)

// WaitForSum waits for a duration and returns the sum of the numbers sent to the channel.
func WaitForSum(duration time.Duration) int {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), duration)
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
