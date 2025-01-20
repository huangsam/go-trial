package concurrency_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/concurrency"
	"github.com/stretchr/testify/assert"
)

func TestGetAnswersWithWaitGroup(t *testing.T) {
	results := concurrency.GetAnswersWithWaitGroup()
	for i := 0; i < 100; i++ {
		assert.Equal(t, i*2, results[i])
	}
}

func BenchmarkGetAnswersWithWaitGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = concurrency.GetAnswersWithWaitGroup()
	}
}
