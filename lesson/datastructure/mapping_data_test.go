package datastructure_test

import (
	"slices"
	"testing"

	"github.com/huangsam/go-trial/lesson/datastructure"
	"github.com/stretchr/testify/assert"
)

func TestGetFruitNames(t *testing.T) {
	expected := []string{"apple", "banana", "cherry"}
	names := slices.Collect(datastructure.GetFruitNames())
	assert.ElementsMatch(t, expected, names)
}

func TestFruitNameExists(t *testing.T) {
	assert.True(t, datastructure.FruitNameExists("apple"))
	assert.False(t, datastructure.FruitNameExists("durian"))
}
