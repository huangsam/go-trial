package concurrency_test

import (
	"testing"

	"io.huangsam/trial/pkg/concurrency"
)

func TestGetAnswersWithChannels(t *testing.T) {
	// Get the results from the function
	results := concurrency.GetAnswersWithChannels()

	// Compare the results with the expected values
	if results[0] != 0 {
		t.Errorf("Expected 0 at index 0 but got %d", results[0])
	}
	if results[99] != 198 {
		t.Errorf("Expected 198 at index 0 but got %d", results[99])
	}
}

func BenchmarkGetAnswersWithChannels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = concurrency.GetAnswersWithChannels()
	}
}
