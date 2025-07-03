package concurrency

import "github.com/huangsam/go-trial/internal/lesson"

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

// MultiStagePipelineSimple creates a multi-stage pipeline by using chained channels.
//
// It generates a range of integers, squares them, doubles them, and sums the results.
// This shows the basic form of a pipeline with a single input and a final output.
//
// Kudos to https://go.dev/blog/pipelines for the original code!
func MultiStagePipelineSimple(from int, to int) int {
	sum := 0
	for n := range double(square(lesson.Range(from, to))) {
		sum += n
	}
	return sum
}

// MultiStagePipelineMerge extends the simple pipeline by using merged channels.
//
// It creates multiple square and double channels, merges them, and then sums the output.
// The number of square and double channels can be adjusted to control the level of concurrency.
// This demonstrates the flexibility of the pipeline approach.
func MultiStagePipelineMerge(from int, to int) int {
	sum := 0

	in := lesson.Range(from, to)

	// Create a merged channel of square channels
	squareChans := make([]<-chan int, squareChannelCount)
	for i := range squareChans {
		squareChans[i] = square(in)
	}
	squareMergedChan := lesson.Merge(squareChans...)

	// Create a merged channel of double channels
	doubleChans := make([]<-chan int, doubleChannelCount)
	for i := range doubleChans {
		doubleChans[i] = double(squareMergedChan)
	}
	doubleMergedChan := lesson.Merge(doubleChans...)

	// Sum the results from the merged double channels
	for n := range doubleMergedChan {
		sum += n
	}

	return sum
}
