package realworld_test

import (
	"testing"
	"time"

	"github.com/huangsam/go-trial/pkg/realworld"
)

func TestCompareDatetime(t *testing.T) {
	now := time.Now()
	previous := now.Add(-time.Hour * 24)

	testCases := []struct {
		datetime1 time.Time
		datetime2 time.Time
		expected  int
		message   string
	}{
		{previous, now, -1, "datetime1 should be before datetime2"},
		{now, previous, 1, "datetime1 should be after datetime2"},
		{now, now, 0, "datetime1 should be equal to datetime2"},
	}

	for _, tc := range testCases {
		if result := realworld.CompareDatetime(tc.datetime1, tc.datetime2); result != tc.expected {
			t.Error(tc.message)
		}
	}
}
