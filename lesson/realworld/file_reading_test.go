package realworld_test

import (
	"testing"

	"github.com/huangsam/go-trial/lesson/realworld"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadLinesSuccess(t *testing.T) {
	lines, err := realworld.ReadLines(fixturesPath + "/test.txt")
	require.NoError(t, err)
	assert.Equal(t, []string{"Hello, world!", "This is a test file."}, lines)
}

func TestReadLinesFailure(t *testing.T) {
	lines, err := realworld.ReadLines(fixturesPath + "/nonexistent.txt")
	require.Error(t, err)
	assert.Nil(t, lines)
}
