package util

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
