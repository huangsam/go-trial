package basics_test

import (
	"testing"

	"github.com/huangsam/go-trial/lesson/basics"
	"github.com/stretchr/testify/assert"
)

func TestPositiveSliceContains(t *testing.T) {
	keys := basics.PositiveSliceContains(3)
	assert.NotNil(t, keys)
}

func TestPositiveSliceIsPositive(t *testing.T) {
	flag := basics.PositiveSliceIsPositive()
	assert.True(t, flag)
}
