package realworld

import (
	"time"
)

// CompareTime compares two time values.
func CompareTime(time1 time.Time, time2 time.Time) int {
	if time1.Before(time2) {
		return -1
	} else if time1.After(time2) {
		return 1
	} else {
		return 0
	}
}
