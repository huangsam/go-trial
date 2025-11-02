package lesson

import "sync"

// Merge combines multiple channels into a single channel.
func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start a goroutine to output from each input channel until they are all closed.
	for _, c := range cs {
		wg.Go(func() {
			for n := range c {
				out <- n
			}
		})
	}

	// Start a goroutine to close out once all the output goroutines are done.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Range creates a channel of integers from a range.
func Range(from int, to int) <-chan int {
	out := make(chan int)
	go func() {
		for n := from; n <= to; n++ {
			out <- n
		}
		close(out)
	}()
	return out
}
