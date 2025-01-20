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
// the updated array. The function uses a mutex to protect concurrent access to the shared array.
func GetAnswersWithWaitGroup() [100]int {
	var wg sync.WaitGroup
	var answers [100]int
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, mu *sync.Mutex) {
			update(&answers, i*10, (i+1)*10, mu)
			wg.Done() // Signal that the goroutine has finished
		}(i, &mu)
	}
	wg.Wait() // Wait for all goroutines to finish
	return answers
}
