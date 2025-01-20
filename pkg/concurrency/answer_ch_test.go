package concurrency_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/concurrency"
	"github.com/stretchr/testify/assert"
)

func TestGetAnswersWithChannels(t *testing.T) {
	results := concurrency.GetAnswersWithChannels()
	for i := 0; i < 100; i++ {
		assert.Equal(t, i*2, results[i])
	}
}

func BenchmarkGetAnswersWithChannels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = concurrency.GetAnswersWithChannels()
	}
}
