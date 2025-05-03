package concurrency

import (
	"context"
	"time"
)

// SumUntil sums numbers until the timeout is reached.
func SumUntil(timeout time.Duration, factor int) int {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	ticker := time.NewTicker(100 * time.Millisecond)
	ch := make(chan int)
	sum := 0
	defer cancel()
	go func() {
		defer close(ch) // Close the channel when goroutine ends
		input := 1
		for {
			select {
			case <-ctx.Done(): // Return when timeout is reached
				return
			case <-ticker.C: // Send the number to the channel every 100ms
				ch <- factor * input
			}
			input += 1
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
