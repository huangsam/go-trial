package basic

import (
	"testing"
)

func TestSliceValuesArePositive(t *testing.T) {
	positiveSlice := []int32{1, 2, 3, 4}
	for i := 0; i < len(positiveSlice); i++ {
		if positiveSlice[i] <= 0 {
			t.Errorf("Found non-positive integer %d", positiveSlice[i])
		}
	}
}
