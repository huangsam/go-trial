package performance_test

import (
	"testing"

	"github.com/huangsam/go-trial/lesson/performance"
	"github.com/stretchr/testify/assert"
)

func TestProfileWorkload(t *testing.T) {
	assert.Zero(t, performance.ProfileWorkload(0))
	assert.NotZero(t, performance.ProfileWorkload(100))
	assert.NotEqual(t, performance.ProfileWorkload(10), performance.ProfileWorkload(11))
}

var profileChecksum uint64

func BenchmarkProfileWorkload(b *testing.B) {
	for b.Loop() {
		profileChecksum = performance.ProfileWorkload(1000)
	}
}
