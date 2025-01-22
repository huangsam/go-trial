package utils_test

import (
	"testing"

	"github.com/huangsam/go-trial/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	t.Run("Range from 1 to 5", func(t *testing.T) {
		expected := []int{1, 2, 3, 4, 5}
		var result []int
		for n := range utils.Range(1, 5) {
			result = append(result, n)
		}
		assert.Equal(t, expected, result)
	})

	t.Run("Range from -2 to 2", func(t *testing.T) {
		expected := []int{-2, -1, 0, 1, 2}
		var result []int
		for n := range utils.Range(-2, 2) {
			result = append(result, n)
		}
		assert.Equal(t, expected, result)
	})

	t.Run("Range from 5 to 5", func(t *testing.T) {
		expected := []int{5}
		var result []int
		for n := range utils.Range(5, 5) {
			result = append(result, n)
		}
		assert.Equal(t, expected, result)
	})

	t.Run("Range from 5 to 1", func(t *testing.T) {
		var result []int
		for n := range utils.Range(5, 1) {
			result = append(result, n)
		}
		assert.Empty(t, result)
	})
}
