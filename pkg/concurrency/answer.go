package concurrency

import (
	"sync"
	"time"
)

// answersRoutineCount is the number of goroutines to use to update the answers array.
const answersRoutineCount int = 10

// GetAnswersWithWaitGroup updates a 100-element array using goroutines and a WaitGroup.
func GetAnswersWithWaitGroup() [100]int {
	var answers [100]int
	var wg sync.WaitGroup
	update := func(from int, to int) {
		for i := from; i < to; i++ {
			time.Sleep(10 * time.Millisecond)
			answers[i] = i * 2
		}
	}
	wg.Add(answersRoutineCount)
	for i := 0; i < answersRoutineCount; i++ {
		go func(i int) {
			update(i*10, (i+1)*10) // Skip mutex due to isolation
			wg.Done()              // Signal that the goroutine has finished
		}(i)
	}
	wg.Wait() // Wait for all goroutines to finish
	return answers
}

// GetAnswersWithChannels updates a 100-element array using goroutines and a channel.
func GetAnswersWithChannels() [100]int {
	var answers [100]int
	done := make(chan struct{}, answersRoutineCount) // Channel to track completion
	update := func(start, end int) {
		time.Sleep(10 * time.Millisecond)
		for i := start; i < end; i++ {
			answers[i] = i * 2
		}
		done <- struct{}{} // Signal completion
	}
	for i := 0; i < answersRoutineCount; i++ {
		go update(i*10, (i+1)*10)
	}
	for i := 0; i < answersRoutineCount; i++ {
		<-done
	}
	return answers
}
