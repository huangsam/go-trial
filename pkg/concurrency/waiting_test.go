package concurrency_test

import (
	"testing"
	"time"

	"github.com/huangsam/go-trial/pkg/concurrency"
	"github.com/stretchr/testify/assert"
)

func TestWaitForSum(t *testing.T) {
	result := concurrency.WaitForSum(time.Millisecond * 350)
	assert.True(t, result == 3 || result == 4)
}

func BenchmarkWaitForSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concurrency.WaitForSum(time.Millisecond * 350)
	}
}
