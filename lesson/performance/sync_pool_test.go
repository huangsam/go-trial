package performance_test

import (
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/huangsam/go-trial/lesson/performance"
	"github.com/stretchr/testify/assert"
)

func TestFormatMessageWithPool(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "name", input: "Go", expected: "Hello, Go!"},
		{name: "empty name", input: "", expected: "Hello, !"},
		{name: "large name", input: strings.Repeat("x", 2048), expected: "Hello, " + strings.Repeat("x", 2048) + "!"},
		{name: "after large name", input: "Go", expected: "Hello, Go!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, performance.FormatMessageWithPool(tt.input))
		})
	}
}

func TestFormatMessageWithPoolConcurrent(t *testing.T) {
	const callCount = 100
	results := make([]string, callCount)
	var wg sync.WaitGroup

	for i := range callCount {
		wg.Go(func() {
			results[i] = performance.FormatMessageWithPool(fmt.Sprint(i))
		})
	}
	wg.Wait()

	for i, result := range results {
		assert.Equal(t, fmt.Sprintf("Hello, %d!", i), result)
	}
}

var benchmarkMessage string

func BenchmarkFormatMessage(b *testing.B) {
	const name = "Go"

	b.Run("Sprintf", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			benchmarkMessage = fmt.Sprintf("Hello, %s!", name)
		}
	})
	b.Run("SyncPool", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			benchmarkMessage = performance.FormatMessageWithPool(name)
		}
	})
}
