package concurrency_test

import (
	"testing"

	"github.com/huangsam/go-trial/lesson/concurrency"
	"github.com/stretchr/testify/assert"
)

func TestGetAnswers(t *testing.T) {
	tests := []struct {
		name     string
		function func() [100]int
	}{
		{"Channels", concurrency.GetAnswersWithChannels},
		{"WaitGroup", concurrency.GetAnswersWithWaitGroup},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := tt.function()
			for i := 0; i < 100; i++ {
				assert.Equal(t, i*2, results[i])
			}
		})
	}
}

func BenchmarkGetAnswers(b *testing.B) {
	benchmarks := []struct {
		name     string
		function func() [100]int
	}{
		{"Channels", concurrency.GetAnswersWithChannels},
		{"WaitGroup", concurrency.GetAnswersWithWaitGroup},
	}

	for _, bb := range benchmarks {
		b.Run(bb.name, func(b *testing.B) {
			for b.Loop() {
				_ = bb.function()
			}
		})
	}
}
