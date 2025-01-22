package basicintro_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/basicintro"
	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{5, 8},
		{10, 89},
	}
	for _, tc := range testCases {
		t.Run("Iterative", func(t *testing.T) {
			actual := basicintro.FibonacciIterative(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
		t.Run("Recursive", func(t *testing.T) {
			actual := basicintro.FibonacciRecursive(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkFibonacci(b *testing.B) {
	b.Run("Iterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			basicintro.FibonacciIterative(40)
		}
	})
	b.Run("Recursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			basicintro.FibonacciRecursive(40)
		}
	})
}
