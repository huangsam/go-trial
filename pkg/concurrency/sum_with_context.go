package concurrency

import (
	"context"
	"time"
)

// SumUntil sums numbers until the timeout is reached.
//
// It sends the numbers to a channel every 100ms and adds them to the sum.
// The function returns the sum when the timeout is reached.
// The function takes a timeout and a factor as arguments.
// The factor is used to multiply the numbers before adding them to the sum.
func SumUntil(ctx context.Context, factor int) int {
	ticker := time.NewTicker(100 * time.Millisecond)
	ch := make(chan int)
	sum := 0
	go func() {
		defer close(ch) // Close the channel when goroutine ends
		input := 1
		for {
			select {
			case <-ctx.Done(): // Return when context is done
				return
			case <-ticker.C: // Send the number to the channel every 100ms
				ch <- factor * input
			}
			input += 1
		}
	}()
	for {
		select {
		case <-ctx.Done(): // Return sum when context is done
			return sum
		case num := <-ch: // Add the number received from the channel to the sum
			sum += num
		}
	}
}
