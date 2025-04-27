package basicintro_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/basicintro"
	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"Zero", 0, 1},
		{"One", 1, 1},
		{"Two", 2, 2},
		{"Three", 3, 3},
		{"Five", 5, 8},
		{"Ten", 10, 89},
	}
	for _, tc := range testCases {
		t.Run(tc.name+"Iterative", func(t *testing.T) {
			actual := basicintro.FibonacciIterative(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
		t.Run(tc.name+"Recursive", func(t *testing.T) {
			actual := basicintro.FibonacciRecursive(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkFibonacci(b *testing.B) {
	n := 40
	b.Run("Iterative", func(b *testing.B) {
		for b.Loop() {
			basicintro.FibonacciIterative(n)
		}
	})
	b.Run("Recursive", func(b *testing.B) {
		for b.Loop() {
			basicintro.FibonacciRecursive(n)
		}
	})
}
