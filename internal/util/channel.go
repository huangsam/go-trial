package util

import "sync"

// Merge combines multiple channels into a single channel.
func Merge(cs []<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// output copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	// Start a goroutine to output from each input channel until they are all closed.
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
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
