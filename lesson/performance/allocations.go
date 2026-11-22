package performance

// DoubleIntsWithAppend doubles values by growing a slice as it appends.
func DoubleIntsWithAppend(values []int) []int {
	if values == nil {
		return nil
	}

	var result []int
	for _, value := range values {
		result = append(result, value*2)
	}
	return result
}

// DoubleIntsPreallocated doubles values in a slice with the required capacity.
func DoubleIntsPreallocated(values []int) []int {
	if values == nil {
		return nil
	}

	result := make([]int, len(values))
	for i, value := range values {
		result[i] = value * 2
	}
	return result
}
