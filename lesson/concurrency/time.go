package concurrency

import (
	"fmt"
	"time"
)

// CurrentTime demonstrates getting the current time.
func CurrentTime() string {
	now := time.Now()
	return now.Format(time.RFC3339)
}

// ParseRFC3339 demonstrates parsing RFC3339 formatted time.
func ParseRFC3339() (time.Time, error) {
	timestamp := "2024-01-15T10:30:00Z"
	return time.Parse(time.RFC3339, timestamp)
}

// ParseCustomFormat demonstrates parsing custom time formats.
func ParseCustomFormat(dateStr string) (time.Time, error) {
	// Go uses a reference time: Mon Jan 2 15:04:05 MST 2006
	format := "2006-01-02"
	return time.Parse(format, dateStr)
}

// DurationParsing demonstrates parsing durations.
func DurationParsing() (time.Duration, error) {
	duration := "2h30m15s"
	d, err := time.ParseDuration(duration)
	if err != nil {
		return 0, err
	}
	return d, nil
}

// TimeComparison demonstrates comparing times.
func TimeComparison(t1, t2 time.Time) string {
	if t1.Before(t2) {
		return "t1 is before t2"
	} else if t1.After(t2) {
		return "t1 is after t2"
	}
	return "t1 equals t2"
}

// TimeArithmetic demonstrates adding/subtracting time.
func TimeArithmetic() (time.Time, time.Duration) {
	now := time.Now()
	inFiveMinutes := now.Add(5 * time.Minute)
	duration := inFiveMinutes.Sub(now)
	return inFiveMinutes, duration
}

// SleepDemonstration demonstrates time.Sleep.
func SleepDemonstration() string {
	start := time.Now()
	time.Sleep(10 * time.Millisecond)
	elapsed := time.Since(start)
	return fmt.Sprintf("Slept for %v", elapsed)
}

// UnixTimestamps demonstrates working with Unix timestamps.
func UnixTimestamps() (int64, int64) {
	now := time.Now()
	unixSeconds := now.Unix()
	unixMilli := now.UnixMilli()
	return unixSeconds, unixMilli
}
