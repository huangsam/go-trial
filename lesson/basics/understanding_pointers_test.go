package basics_test

import (
	"testing"

	"github.com/huangsam/go-trial/lesson/basics"
	"github.com/stretchr/testify/assert"
)

func TestPointerSetOne(t *testing.T) {
	var n int
	basics.PointerSetOne(&n)
	assert.Equal(t, 1, n)
}
