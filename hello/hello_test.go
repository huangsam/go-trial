package hello

import (
	"testing"
)

func TestWorld(t *testing.T) {
	expected := "Hello world"
	actual := World()

	if actual != expected {
		t.Errorf("World() = %v, want %v", actual, expected)
	}
}

func TestName(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Alice", "Hello Alice\n"},
		{"Bob", "Hello Bob\n"},
		{"", "Hello \n"},
	}

	for _, tc := range testCases {
		actual := Name(tc.input)

		if actual != tc.expected {
			t.Errorf("Name(%q) = %q, want %q", tc.input, actual, tc.expected)
		}
	}
}

func BenchmarkWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		World()
	}
}

func BenchmarkName(b *testing.B) {
	name := "Alice"
	for i := 0; i < b.N; i++ {
		Name(name)
	}
}
