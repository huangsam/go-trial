package basicintro_test

import (
	"testing"

	"io.huangsam/trial/pkg/basicintro"
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
		if actual != tc.expected {
			t.Errorf("Fibonacci(%d) expected: %d, got: %d", tc.input, tc.expected, actual)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		basicintro.Fibonacci(40)
	}
}
