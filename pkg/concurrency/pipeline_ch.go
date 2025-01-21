package concurrency

// generate creates a channel of integers from a range.
func generate(from int, to int) <-chan int {
	out := make(chan int)
	go func() {
		for n := from; n <= to; n++ {
			out <- n
		}
		close(out)
	}()
	return out
}

// square squares the input.
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// double doubles the input.
func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

// MultiStagePipeline creates a multi-stage pipeline that generates a range of integers,
// squares each integer, and then doubles the result. The function returns the sum of
// all the doubled squared integers.
func MultiStagePipeline(from int, to int) int {
	sum := 0
	for n := range double(square(generate(from, to))) {
		sum += n
	}
	return sum
}
