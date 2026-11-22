package performance_test

import (
	"testing"

	"github.com/huangsam/go-trial/lesson/performance"
	"github.com/stretchr/testify/assert"
)

func TestDoubleInts(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected []int
	}{
		{name: "nil input", values: nil, expected: nil},
		{name: "values", values: []int{1, 2, 3, 4}, expected: []int{2, 4, 6, 8}},
	}

	doubleFunctions := []struct {
		name     string
		function func([]int) []int
	}{
		{name: "append", function: performance.DoubleIntsWithAppend},
		{name: "preallocated", function: performance.DoubleIntsPreallocated},
	}

	for _, tt := range tests {
		for _, doubleFunction := range doubleFunctions {
			t.Run(tt.name+"/"+doubleFunction.name, func(t *testing.T) {
				doubled := doubleFunction.function(tt.values)
				assert.Equal(t, tt.expected, doubled)

				if len(doubled) > 0 {
					doubled[0] = 0
					assert.Equal(t, 1, tt.values[0])
				}
			})
		}
	}
}

var benchmarkDoubledInts []int

func BenchmarkDoubleInts(b *testing.B) {
	values := make([]int, 1024)
	for i := range values {
		values[i] = i
	}

	benchmarks := []struct {
		name     string
		function func([]int) []int
	}{
		{name: "append", function: performance.DoubleIntsWithAppend},
		{name: "preallocated", function: performance.DoubleIntsPreallocated},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				benchmarkDoubledInts = benchmark.function(values)
			}
		})
	}
}
