package abstraction_test

import (
	"sync"
	"testing"

	"github.com/huangsam/go-trial/lesson/abstraction"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunETL(t *testing.T) {
	// Test data
	inputData := []int{1, 2, 3, 4, 5}
	expectedTransformedData := []string{"1", "2", "3", "4", "5"}

	// Extractor function
	extractor := func() <-chan int {
		ch := make(chan int, len(inputData))
		for _, v := range inputData {
			ch <- v
		}
		close(ch)
		return ch
	}

	// Transformer function
	transformer := func(input int) string {
		return string(rune(input + '0'))
	}

	// Loader function
	var loadedData []string
	var mu sync.Mutex
	loader := func(data string) {
		mu.Lock()
		defer mu.Unlock()
		loadedData = append(loadedData, data)
	}

	// Run ETL
	abstraction.RunETL(extractor, transformer, loader)

	// Assertions
	require.Len(t, loadedData, len(expectedTransformedData))
	assert.ElementsMatch(t, expectedTransformedData, loadedData)
}
