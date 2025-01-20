package datastructure_test

import (
	"testing"

	"io.huangsam/trial/pkg/datastructure"
)

func TestMapCanAccessValues(t *testing.T) {
	data := datastructure.SampleFruitNumberMap
	if data["apple"] != 1 {
		t.Errorf("Expected apple value to be 1, got %d", data["apple"])
	}
	if data["banana"] != 2 {
		t.Errorf("Expected banana value to be 2, got %d", data["banana"])
	}
	if data["cherry"] != 3 {
		t.Errorf("Expected cherry value to be 3, got %d", data["cherry"])
	}
	if _, ok := data["bogus"]; ok {
		t.Error("Expected bogus to be missing")
	}
}
