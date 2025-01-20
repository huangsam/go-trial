package datastructure_test

import (
	"testing"

	"io.huangsam/trial/pkg/datastructure"
)

func TestSliceValuesArePositive(t *testing.T) {
	data := datastructure.SamplePositiveSlice
	for i := 0; i < len(data); i++ {
		if data[i] <= 0 {
			t.Errorf("Found non-positive integer %d", data[i])
		}
	}
}
