package realworld_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/realworld"
	"github.com/stretchr/testify/assert"
)

func TestReadLinesSuccess(t *testing.T) {
	lines, err := realworld.ReadLines(fixturesPath + "/test.txt")
	assert.Nil(t, err)
	assert.Equal(t, []string{"Hello, world!", "This is a test file."}, lines)
}

func TestReadLinesFailure(t *testing.T) {
	lines, err := realworld.ReadLines(fixturesPath + "/nonexistent.txt")
	assert.NotNil(t, err)
	assert.Nil(t, lines)
}
