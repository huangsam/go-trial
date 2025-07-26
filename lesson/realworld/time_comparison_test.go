package realworld_test

import (
	"testing"
	"time"

	"github.com/huangsam/go-trial/lesson/realworld"
)

func TestCompareTime(t *testing.T) {
	now := time.Now()
	previous := now.Add(-time.Hour * 24)
	testCases := []struct {
		time1    time.Time
		time2    time.Time
		expected int
		message  string
	}{
		{previous, now, -1, "time1 should be before time2"},
		{now, previous, 1, "time1 should be after time2"},
		{now, now, 0, "time1 should be equal to time2"},
	}
	for _, tc := range testCases {
		if result := realworld.CompareTime(tc.time1, tc.time2); result != tc.expected {
			t.Error(tc.message)
		}
	}
}
