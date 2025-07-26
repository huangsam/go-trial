package realworld

import "time"

// CompareTime compares two time values.
func CompareTime(time1 time.Time, time2 time.Time) int {
	switch {
	case time1.Before(time2):
		return -1
	case time1.After(time2):
		return 1
	default:
		return 0
	}
}
