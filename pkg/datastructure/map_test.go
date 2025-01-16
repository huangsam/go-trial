package datastructure_test

import (
	"testing"
)

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
