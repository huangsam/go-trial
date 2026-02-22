package datastructure

import "slices"

var positiveSlice = []int{1, 2, 3, 4, 5}

// PositiveSliceContains checks if positiveSlice contains a specific value.
func PositiveSliceContains(value int) bool {
	return slices.Contains(positiveSlice, value)
}

// PositiveSliceIsPositive checks if all values in positiveSlice are positive.
func PositiveSliceIsPositive() bool {
	result := true
	for i := range positiveSlice {
		result = result && positiveSlice[i] > 0
	}
	return result
}
