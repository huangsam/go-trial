package concurrency

import (
	"fmt"
	"time"
)

// Timezone demonstrates working with different timezones.
func Timezone() (time.Time, time.Time) {
	// Create a time in UTC
	utcTime := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)

	// Convert to EST (Eastern Standard Time)
	est, _ := time.LoadLocation("America/New_York")
	estTime := utcTime.In(est)

	return utcTime, estTime
}

// DateComponents demonstrates extracting date components.
func DateComponents() (int, time.Month, int) {
	t := time.Date(2024, 3, 15, 14, 30, 0, 0, time.UTC)
	return t.Year(), t.Month(), t.Day()
}

// TimeFormatting demonstrates various time formatting options.
func TimeFormatting() map[string]string {
	t := time.Date(2024, 1, 15, 14, 30, 45, 0, time.UTC)

	formats := map[string]string{
		"RFC3339":  t.Format(time.RFC3339),
		"ANSIC":    t.Format(time.ANSIC),
		"UnixDate": t.Format(time.UnixDate),
		"RubyDate": t.Format(time.RubyDate),
		"Kitchen":  t.Format(time.Kitchen),
		"Custom1":  t.Format("02/01/2006 03:04 PM"),
		"Custom2":  t.Format("2006-01-02 15:04:05"),
	}

	return formats
}

// TimeUntil demonstrates calculating time until a future event.
func TimeUntil(future time.Time) (time.Duration, string) {
	now := time.Now()
	if future.Before(now) || future.Equal(now) {
		return 0, "event already passed"
	}
	duration := future.Sub(now)
	return duration, fmt.Sprintf("%v remaining", duration.Round(time.Second))
}

// TimeSince demonstrates calculating time elapsed since a past event.
func TimeSince(past time.Time) (time.Duration, string) {
	now := time.Now()
	if past.After(now) {
		return 0, "event not yet occurred"
	}
	duration := now.Sub(past)
	return duration, fmt.Sprintf("%v elapsed", duration.Round(time.Second))
}

// StartOfDay demonstrates getting the start of a day.
func StartOfDay() (time.Time, error) {
	t, err := time.Parse("2006-01-02", "2024-03-15")
	if err != nil {
		return time.Time{}, err
	}

	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return start, nil
}

// EndOfDay demonstrates getting the end of a day.
func EndOfDay() (time.Time, error) {
	t, err := time.Parse("2006-01-02", "2024-03-15")
	if err != nil {
		return time.Time{}, err
	}

	end := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
	return end, nil
}

// Weekday demonstrates getting the day of week.
func Weekday() (time.Weekday, string) {
	t := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	weekday := t.Weekday()
	return weekday, weekday.String()
}
