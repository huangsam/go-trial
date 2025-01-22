package realworld

import (
	"time"
)

// CompareDatetime compares two datetime values.
func CompareDatetime(datetime1 time.Time, datetime2 time.Time) int {
	if datetime1.Before(datetime2) {
		return -1
	} else if datetime1.After(datetime2) {
		return 1
	} else {
		return 0
	}
}
