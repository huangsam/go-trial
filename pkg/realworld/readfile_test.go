package realworld_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/realworld"
	"github.com/stretchr/testify/assert"
)

func TestReadLines(t *testing.T) {
	lines, err := realworld.ReadLines("testdata/test.txt")
	assert.Nil(t, err)
	assert.Equal(t, []string{"Hello, world!", "This is a test file."}, lines)
}
