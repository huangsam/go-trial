package datastructure_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/datastructure"
	"github.com/stretchr/testify/assert"
)

func TestPositiveSliceContains(t *testing.T) {
	keys := datastructure.PositiveSliceContains(3)
	assert.NotNil(t, keys)
}

func TestSliceValuesArePositive(t *testing.T) {
	data := datastructure.PositiveSlice
	for i := 0; i < len(data); i++ {
		assert.LessOrEqual(t, 0, data[i])
	}
}
