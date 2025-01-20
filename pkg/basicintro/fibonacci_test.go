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
		actual := basicintro.Fibonacci(tc.input)
		assert.Equal(t, tc.expected, actual)
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		basicintro.Fibonacci(40)
	}
}
