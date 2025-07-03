package cmd_test

import (
	"errors"
	"testing"

	"github.com/huangsam/go-trial/internal/cmd"
	"github.com/stretchr/testify/assert"
)

func TestDismiss(t *testing.T) {
	// Test case where the function succeeds
	called := false
	cmd.Dismiss(func() error {
		called = true
		return nil
	})
	assert.True(t, called, "The function should have been called")

	// Test case where the function returns an error
	called = false
	cmd.Dismiss(func() error {
		called = true
		return errors.New("some error")
	})
	assert.True(t, called, "The function should have been called even if it returns an error")
}
