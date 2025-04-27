package util_test

import (
	"testing"
	"time"

	"github.com/huangsam/go-trial/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()

	go func() {
		defer close(ch2)
		ch2 <- 3
		ch2 <- 4
	}()

	go func() {
		defer close(ch3)
		ch3 <- 5
		ch3 <- 6
	}()

	merged := util.Merge(ch1, ch2, ch3)

	var result []int
	for val := range merged {
		result = append(result, val)
	}

	expected := []int{1, 2, 3, 4, 5, 6}
	assert.ElementsMatch(t, expected, result)
}

func TestMergeEmptyChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
	}()

	go func() {
		defer close(ch2)
	}()

	merged := util.Merge(ch1, ch2)

	var result []int
	for val := range merged {
		result = append(result, val)
	}

	assert.Empty(t, result)
}

func TestMergeWithDelay(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		time.Sleep(100 * time.Millisecond)
		ch1 <- 2
	}()

	go func() {
		defer close(ch2)
		ch2 <- 3
		time.Sleep(50 * time.Millisecond)
		ch2 <- 4
	}()

	merged := util.Merge(ch1, ch2)

	var result []int
	for val := range merged {
		result = append(result, val)
	}

	expected := []int{1, 2, 3, 4}
	assert.ElementsMatch(t, expected, result)
}

func TestRange(t *testing.T) {
	t.Run("Range from 1 to 5", func(t *testing.T) {
		expected := []int{1, 2, 3, 4, 5}
		var result []int
		for n := range util.Range(1, 5) {
			result = append(result, n)
		}
		assert.Equal(t, expected, result)
	})

	t.Run("Range from -2 to 2", func(t *testing.T) {
		expected := []int{-2, -1, 0, 1, 2}
		var result []int
		for n := range util.Range(-2, 2) {
			result = append(result, n)
		}
		assert.Equal(t, expected, result)
	})

	t.Run("Range from 5 to 5", func(t *testing.T) {
		expected := []int{5}
		var result []int
		for n := range util.Range(5, 5) {
			result = append(result, n)
		}
		assert.Equal(t, expected, result)
	})

	t.Run("Range from 5 to 1", func(t *testing.T) {
		var result []int
		for n := range util.Range(5, 1) {
			result = append(result, n)
		}
		assert.Empty(t, result)
	})
}
