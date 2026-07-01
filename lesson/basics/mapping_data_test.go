package basics_test

import (
	"slices"
	"testing"

	"github.com/huangsam/go-trial/lesson/basics"
	"github.com/stretchr/testify/assert"
)

func TestGetFruitNames(t *testing.T) {
	expected := []string{"apple", "banana", "cherry"}
	names := slices.Collect(basics.GetFruitNames())
	assert.ElementsMatch(t, expected, names)
}

func TestFruitNameExists(t *testing.T) {
	assert.True(t, basics.FruitNameExists("apple"))
	assert.False(t, basics.FruitNameExists("durian"))
}
