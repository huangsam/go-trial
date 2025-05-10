package concurrency_test

import (
	"context"
	"testing"
	"time"

	"github.com/huangsam/go-trial/pkg/concurrency"
	"github.com/stretchr/testify/assert"
)

func TestSumUntil(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	result := concurrency.SumUntil(ctx, 2)
	assert.GreaterOrEqual(t, result, 2) // At least 2 after a 100ms tick
	assert.LessOrEqual(t, result, 6)    // At most 2 + 4 after a 200ms timeout
}

func BenchmarkSumUntil(b *testing.B) {
	for b.Loop() {
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		concurrency.SumUntil(ctx, 2)
		cancel()
	}
}
