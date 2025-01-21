package concurrency

import "sync"

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

// merge combines multiple channels into a single channel.
func merge(cs []<-chan int) <-chan int {
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

// MultiStagePipeline creates a multi-stage pipeline that generates a range of integers,
// squares each integer, and then doubles the result. The function returns the sum of
// all the doubled squared integers.
func MultiStagePipelineSimple(from int, to int) int {
	sum := 0
	for n := range double(square(generate(from, to))) {
		sum += n
	}
	return sum
}

// MultiStagePipeline creates a multi-stage pipeline that generates a range of integers,
// squares each integer, and then doubles the result. The function returns the sum of
// all the doubled squared integers.
func MultiStagePipelineMerge(from int, to int) int {
	sum := 0

	in := generate(from, to)
	squareOut := merge([]<-chan int{square(in), square(in)})
	doubleOut := merge([]<-chan int{double(squareOut), double(squareOut)})

	for n := range doubleOut {
		sum += n
	}

	return sum
}
