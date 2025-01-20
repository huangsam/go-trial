package datastructure_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"io.huangsam/trial/pkg/datastructure"
)

func TestMapCanAccessValues(t *testing.T) {
	data := datastructure.FruitNumberMap
	assert.Equal(t, 1, data["apple"], "Expected apple value to be 1")
	assert.Equal(t, 2, data["banana"], "Expected banana value to be 2")
	assert.Equal(t, 3, data["cherry"], "Expected cherry value to be 3")
	_, ok := data["bogus"]
	assert.False(t, ok, "Expected bogus to be missing")
}
