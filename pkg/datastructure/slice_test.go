package datastructure_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/datastructure"
	"github.com/stretchr/testify/assert"
)

func TestSliceValuesArePositive(t *testing.T) {
	data := datastructure.PositiveSlice
	for i := 0; i < len(data); i++ {
		assert.LessOrEqual(t, 0, data[i])
	}
}
