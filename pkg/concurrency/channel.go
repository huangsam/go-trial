package concurrency

import (
	"sync"
	"time"
)

// processTask receives work requests on the 'taskQueue' channel and sends results on the 'resultChannel' channel.
func processTask(taskQueue <-chan int, resultChannel chan<- int) {
	for task := range taskQueue {
		time.Sleep(10 * time.Millisecond)
		resultChannel <- task * 2
	}
}

// distributeTasks distributes work requests to the taskQueue and collects results from the resultChannel.
func distributeTasks(m *[100]int, from int, to int, taskQueue chan<- int, resultChannel <-chan int) {
	for i := from; i < to; i++ {
		taskQueue <- i
	}
	close(taskQueue)

	for i := from; i < to; i++ {
		(*m)[i] = <-resultChannel
	}
}

// GetAnswersWithChannels concurrently updates portions of a 100-element integer array
// using channels for communication and synchronization between goroutines.
// Each set of tasks has its own dedicated taskQueue and resultChannel to ensure
// that results are received and stored in the correct order.
func GetAnswersWithChannels() [100]int {
	var wg sync.WaitGroup
	var answers [100]int
	for i := 0; i < 10; i++ {
		taskQueue := make(chan int, 10)
		resultChannel := make(chan int, 10) // Create a result channel for each set of tasks
		wg.Add(1)
		go processTask(taskQueue, resultChannel)
		go func(i int) {
			distributeTasks(&answers, i*10, (i+1)*10, taskQueue, resultChannel)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return answers
}
