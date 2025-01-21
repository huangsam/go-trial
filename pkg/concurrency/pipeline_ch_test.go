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
		t.Run(tt.name, func(t *testing.T) {
			actualSum := concurrency.MultiStagePipeline(tt.start, tt.end)
			assert.Equal(t, tt.expectedSum, actualSum)
		})
	}
}

func BenchmarkMultiStagePipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = concurrency.MultiStagePipeline(1, 1000)
	}
}
