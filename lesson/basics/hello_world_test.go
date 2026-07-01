package basics_test

import (
	"testing"

	"github.com/huangsam/go-trial/lesson/basics"
	"github.com/stretchr/testify/assert"
)

func TestWorld(t *testing.T) {
	expected := "Hello World"
	actual := basics.GreetWorld()
	assert.Equal(t, expected, actual)
}

func TestName(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Alice", "Hello Alice"},
		{"Bob", "Hello Bob"},
		{"", "Hello "},
	}
	for _, tc := range testCases {
		actual := basics.GreetName(tc.input)
		assert.Equal(t, tc.expected, actual)
	}
}

func BenchmarkWorld(b *testing.B) {
	for b.Loop() {
		basics.GreetWorld()
	}
}

func BenchmarkName(b *testing.B) {
	name := "Alice"
	for b.Loop() {
		basics.GreetName(name)
	}
}
