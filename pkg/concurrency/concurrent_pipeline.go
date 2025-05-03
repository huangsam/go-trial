package concurrency

import "github.com/huangsam/go-trial/internal/util"

const (
	// squareChannelCount is the number of square channels.
	squareChannelCount = 4

	// doubleChannelCount is the number of double channels.
	doubleChannelCount = 2
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

// MultiStagePipelineSimple creates a multi-stage pipeline that generates a range of integers,
// squares them, and then doubles them. It then sums the output.
func MultiStagePipelineSimple(from int, to int) int {
	sum := 0
	for n := range double(square(util.Range(from, to))) {
		sum += n
	}
	return sum
}

// MultiStagePipelineMerge extends the simple pipeline with a merge function.
func MultiStagePipelineMerge(from int, to int) int {
	sum := 0

	in := util.Range(from, to)

	// Create a merged channel of square channels
	squareChans := make([]<-chan int, squareChannelCount)
	for i := range squareChans {
		squareChans[i] = square(in)
	}
	squareMergedChan := util.Merge(squareChans...)

	// Create a merged channel of double channels
	doubleChans := make([]<-chan int, doubleChannelCount)
	for i := range doubleChans {
		doubleChans[i] = double(squareMergedChan)
	}
	doubleMergedChan := util.Merge(doubleChans...)

	// Sum the results from the merged double channels
	for n := range doubleMergedChan {
		sum += n
	}

	return sum
}
