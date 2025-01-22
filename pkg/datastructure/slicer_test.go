package datastructure_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/datastructure"
	"github.com/stretchr/testify/assert"
)

func TestPositiveSliceContains(t *testing.T) {
	keys := datastructure.PositiveSliceContains(3)
	assert.NotNil(t, keys)
}

func TestPositiveSliceIsPositive(t *testing.T) {
	flag := datastructure.PositiveSliceIsPositive()
	assert.True(t, flag)
}
