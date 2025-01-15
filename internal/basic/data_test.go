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

func TestMapCanAccessValues(t *testing.T) {
	myMap := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}
	if myMap["apple"] != 1 {
		t.Errorf("Expected apple value to be 1, got %d", myMap["apple"])
	}
	if myMap["banana"] != 2 {
		t.Errorf("Expected banana value to be 2, got %d", myMap["banana"])
	}
	if myMap["cherry"] != 3 {
		t.Errorf("Expected cherry value to be 3, got %d", myMap["cherry"])
	}
	if _, ok := myMap["bogus"]; ok {
		t.Error("Expected bogus to be missing")
	}
}
