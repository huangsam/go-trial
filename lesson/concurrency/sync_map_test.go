package concurrency_test

import (
	"testing"

	"github.com/huangsam/go-trial/lesson/concurrency"
	"github.com/stretchr/testify/assert"
)

func TestCountWordsSyncMap(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected map[string]int64
	}{
		{
			name:     "repeated words",
			words:    []string{"go", "sync", "go", "map", "sync", "go"},
			expected: map[string]int64{"go": 3, "sync": 2, "map": 1},
		},
		{
			name:     "empty input",
			words:    nil,
			expected: map[string]int64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, concurrency.CountWordsSyncMap(tt.words))
		})
	}
}
