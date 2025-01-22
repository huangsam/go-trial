package concurrency

import (
	"sync"

	"github.com/huangsam/go-trial/internal/utils"
)

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

// MultiStagePipelineSimple creates a multi-stage pipeline that generates a range of integers,
// squares them, and then doubles them. It then sums the output.
func MultiStagePipelineSimple(from int, to int) int {
	sum := 0
	for n := range double(square(utils.Range(from, to))) {
		sum += n
	}
	return sum
}

// MultiStagePipelineMerge extends the simple pipeline with a merge function.
func MultiStagePipelineMerge(from int, to int) int {
	sum := 0

	in := utils.Range(from, to)

	// Create an array of square channels
	squareChans := make([]<-chan int, 4)
	for i := range squareChans {
		squareChans[i] = square(in)
	}
	squareOut := merge(squareChans)

	// Create an array of double channels
	doubleChans := make([]<-chan int, 2)
	for i := range doubleChans {
		doubleChans[i] = double(squareOut)
	}
	doubleOut := merge(doubleChans)

	for n := range doubleOut {
		sum += n
	}

	return sum
}
