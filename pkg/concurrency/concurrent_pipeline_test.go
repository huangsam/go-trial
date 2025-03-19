package concurrency_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/concurrency"
	"github.com/stretchr/testify/assert"
)

func TestMultiStagePipeline(t *testing.T) {
	tests := []struct {
		name        string
		start, end  int
		expectedSum int
	}{
		{"Small range", 1, 10, 770},
		{"Medium range", 1, 1000, 667667000},
		{"Large range", 1, 1000000, 666667666667000000},
	}
	for _, tt := range tests {
		t.Run(tt.name+" simple", func(t *testing.T) {
			actualSum := concurrency.MultiStagePipelineSimple(tt.start, tt.end)
			assert.Equal(t, tt.expectedSum, actualSum)
		})
		t.Run(tt.name+" merge", func(t *testing.T) {
			actualSum := concurrency.MultiStagePipelineMerge(tt.start, tt.end)
			assert.Equal(t, tt.expectedSum, actualSum)
		})
	}
}

func BenchmarkMultiStagePipeline(b *testing.B) {
	from, to := 1, 1000
	b.Run("Simple", func(b *testing.B) {
		for b.Loop() {
			_ = concurrency.MultiStagePipelineSimple(from, to)
		}
	})
	b.Run("Merge", func(b *testing.B) {
		for b.Loop() {
			_ = concurrency.MultiStagePipelineMerge(from, to)
		}
	})
}
