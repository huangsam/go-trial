package concurrency_test

import (
	"testing"

	"io.huangsam/trial/pkg/concurrency"
)

func TestGetAnswersWithWaitGroup(t *testing.T) {
	results := concurrency.GetAnswersWithWaitGroup()
	if results[0] != 0 {
		t.Errorf("Expected 0 at index 0 but got %d", results[0])
	}
	if results[99] != 198 {
		t.Errorf("Expected 198 at index 0 but got %d", results[99])
	}
}

func BenchmarkGetAnswersWithWaitGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = concurrency.GetAnswersWithWaitGroup()
	}
}
