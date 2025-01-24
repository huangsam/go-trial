package concurrency

import (
	"sync"
	"time"
)

// update updates a portion of the given integer array concurrently.
func update(m *[100]int, from int, to int, mu *sync.Mutex) {
	for i := from; i < to; i++ {
		time.Sleep(10 * time.Millisecond)
		mu.Lock() // Acquire the mutex
		(*m)[i] = i * 2
		mu.Unlock() // Release the mutex
	}
}

// GetAnswersWithWaitGroup concurrently updates portions of a 100-element integer array
// using a combination of goroutines and sequential calls to the update function.
// It utilizes a WaitGroup to ensure all concurrent operations complete before returning
// the updated array, and it uses a mutex to prevent race conditions during updates.
func GetAnswersWithWaitGroup() [100]int {
	var wg sync.WaitGroup
	var answers [100]int
	var mu sync.Mutex
	wg.Add(answersRoutineCount)
	for i := 0; i < answersRoutineCount; i++ {
		go func(i int, mup *sync.Mutex) {
			update(&answers, i*10, (i+1)*10, mup)
			wg.Done() // Signal that the goroutine has finished
		}(i, &mu)
	}
	wg.Wait() // Wait for all goroutines to finish
	return answers
}

// GetAnswersWithChannels concurrently updates portions of a 100-element integer array.
// It utilizes goroutines to update 10 elements each, and channels to signal completion.
func GetAnswersWithChannels() [100]int {
	var answers [100]int
	done := make(chan struct{}, answersRoutineCount) // Channel to track completion

	// Worker function
	update := func(start, end int) {
		time.Sleep(10 * time.Millisecond)
		for i := start; i < end; i++ {
			answers[i] = i * 2
		}
		done <- struct{}{} // Signal completion
	}

	// Spawn all goroutines
	for i := 0; i < answersRoutineCount; i++ {
		go update(i*10, (i+1)*10)
	}

	// Wait for all goroutines to complete
	for i := 0; i < answersRoutineCount; i++ {
		<-done
	}

	return answers
}
