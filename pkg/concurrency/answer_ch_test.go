package concurrency_test

import (
	"testing"

	"io.huangsam/trial/pkg/concurrency"
)

func TestGetAnswersWithChannels(t *testing.T) {
	results := concurrency.GetAnswersWithChannels()
	for i := 0; i < 100; i++ {
		if results[i] != i*2 {
			t.Errorf("Expected %d at index %d but got %d", i*2, i, results[0])
		}
	}
}

func BenchmarkGetAnswersWithChannels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = concurrency.GetAnswersWithChannels()
	}
}
