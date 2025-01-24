// Package concurrency has examples using goroutines, channels.
package concurrency

var (
	// Number of goroutines to use for updating the answers array.
	answersRoutineCount int = 10

	// Number of channels to use for the square function.
	squareChannelCount int = 4

	// Number of channels to use for the double function.
	doubleChannelCount int = 2
)
