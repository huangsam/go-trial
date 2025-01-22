package datastructure

import "slices"

// Contains checks if PositiveSlice contains a specific value.
func PositiveSliceContains(value int) bool {
	return slices.Contains(PositiveSlice, value)
}
