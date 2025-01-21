package basicintro_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/basicintro"
	"github.com/stretchr/testify/assert"
)

func TestPointerSetOne(t *testing.T) {
	var n int
	basicintro.PointerSetOne(&n)
	assert.Equal(t, 1, n)
}
