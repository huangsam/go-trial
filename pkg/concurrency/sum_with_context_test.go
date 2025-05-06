package concurrency_test

import (
	"context"
	"testing"
	"time"

	"github.com/huangsam/go-trial/pkg/concurrency"
	"github.com/stretchr/testify/assert"
)

func TestSumUntil(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 350*time.Millisecond)
	defer cancel()
	result := concurrency.SumUntil(ctx, 2)
	assert.GreaterOrEqual(t, result, 12)
	assert.LessOrEqual(t, result, 20)
}

func BenchmarkSumUntil(b *testing.B) {
	for b.Loop() {
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		concurrency.SumUntil(ctx, 2)
		cancel()
	}
}
