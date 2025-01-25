package concurrency_test

import (
	"testing"
	"time"

	"github.com/huangsam/go-trial/pkg/concurrency"
	"github.com/stretchr/testify/assert"
)

func TestSumUntil(t *testing.T) {
	result := concurrency.SumUntil(time.Millisecond*350, 2)
	assert.True(t, result == 12 || result == 20, "The result should be 12 or 20 but got %d", result)
}

func BenchmarkSumUntil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concurrency.SumUntil(time.Millisecond*350, 1)
	}
}
