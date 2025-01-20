package datastructure_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"io.huangsam/trial/pkg/datastructure"
)

func TestSliceValuesArePositive(t *testing.T) {
	data := datastructure.SamplePositiveSlice
	for i := 0; i < len(data); i++ {
		assert.LessOrEqual(t, 0, data[i])
	}
}
