package datastructure_test

import (
	"slices"
	"testing"

	"github.com/huangsam/go-trial/pkg/datastructure"
	"github.com/stretchr/testify/assert"
)

func TestGetFruitNames(t *testing.T) {
	expected := []string{"apple", "banana", "cherry"}
	names := slices.Collect(datastructure.GetFruitNames())
	assert.ElementsMatch(t, expected, names, "Expected fruit names to match")
}

func TestMapCanAccessValues(t *testing.T) {
	data := datastructure.FruitNumberMap
	assert.Equal(t, 1, data["apple"], "Expected apple value to be 1")
	assert.Equal(t, 2, data["banana"], "Expected banana value to be 2")
	assert.Equal(t, 3, data["cherry"], "Expected cherry value to be 3")
	_, ok := data["bogus"]
	assert.False(t, ok, "Expected bogus to be missing")
}
