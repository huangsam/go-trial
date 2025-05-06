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

	// Fill the channel with a value on every tick
	go func() {
		defer close(ch)
		input := 1
		for range ticker.C {
			ch <- input * factor
			input += 1
		}
	}()

	// Sum the numbers from the channel until the context is done
	for num := range ch {
		if ctx.Err() != nil {
			break
		}
		sum += num
	}
	return sum
}
