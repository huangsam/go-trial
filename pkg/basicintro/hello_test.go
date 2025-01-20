package basicintro_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/basicintro"
	"github.com/stretchr/testify/assert"
)

func TestWorld(t *testing.T) {
	expected := "Hello World"
	actual := basicintro.GreetWorld()
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
		actual := basicintro.GreetName(tc.input)
		assert.Equal(t, tc.expected, actual)
	}
}

func BenchmarkWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		basicintro.GreetWorld()
	}
}

func BenchmarkName(b *testing.B) {
	name := "Alice"
	for i := 0; i < b.N; i++ {
		basicintro.GreetName(name)
	}
}
