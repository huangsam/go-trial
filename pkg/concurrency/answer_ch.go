package concurrency

// GetAnswersWithChannels concurrently updates portions of a 100-element integer array.
// It utilizes goroutines to update 10 elements each, and channels to signal completion.
func GetAnswersWithChannels() [100]int {
	var answers [100]int
	ch := make(chan int, 10) // Channel to track completion

	// Worker function
	update := func(start, end int) {
		for i := start; i < end; i++ {
			answers[i] = i * 2
		}
		ch <- 0 // Signal completion
	}

	// Spawn 10 goroutines
	for i := 0; i < 10; i++ {
		go update(i*10, (i+1)*10)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-ch
	}

	return answers
}
