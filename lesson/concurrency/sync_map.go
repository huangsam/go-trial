package concurrency

import (
	"sync"
	"sync/atomic"
)

// CountWordsSyncMap counts words concurrently using sync.Map and atomic counters.
// sync.Map is most useful for write-once/read-many or disjoint-key workloads;
// a regular map with a mutex is often simpler for other cases.
func CountWordsSyncMap(words []string) map[string]int64 {
	var counts sync.Map
	var wg sync.WaitGroup

	for _, word := range words {
		wg.Go(func() {
			counter, ok := counts.Load(word)
			if !ok {
				counter, _ = counts.LoadOrStore(word, &atomic.Int64{})
			}
			counter.(*atomic.Int64).Add(1)
		})
	}

	wg.Wait()

	result := make(map[string]int64)
	counts.Range(func(key, value any) bool {
		result[key.(string)] = value.(*atomic.Int64).Load()
		return true
	})
	return result
}
