package concurrency

import (
	"testing"
	"time"
)

func TestTimezone(t *testing.T) {
	utc, est := Timezone()

	diff := utc.Sub(est)
	if diff < -time.Minute || diff > time.Minute {
		t.Errorf("Times don't match: UTC=%v, EST=%v", utc, est)
	}
}

func TestDateComponents(t *testing.T) {
	year, month, day := DateComponents()

	if year != 2024 {
		t.Errorf("Expected year 2024, got %d", year)
	}
	if month != time.March {
		t.Errorf("Expected March, got %v", month)
	}
	if day != 15 {
		t.Errorf("Expected day 15, got %d", day)
	}
}

func TestTimeFormatting(t *testing.T) {
	formats := TimeFormatting()

	rfc := formats["RFC3339"]
	if !contains(rfc, "2024-01-15") {
		t.Errorf("Expected date in RFC3339: %s", rfc)
	}
}

func TestTimeUntil(t *testing.T) {
	future := time.Date(2099, 1, 1, 0, 0, 0, 0, time.UTC)
	duration, message := TimeUntil(future)

	if duration <= 0 {
		t.Errorf("Expected positive duration, got %v", duration)
	}

	if !contains(message, "remaining") {
		t.Errorf("Expected 'remaining' in message: %s", message)
	}
}

func TestTimeSince(t *testing.T) {
	past := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	duration, message := TimeSince(past)

	if duration <= 0 {
		t.Errorf("Expected positive duration, got %v", duration)
	}

	if !contains(message, "elapsed") {
		t.Errorf("Expected 'elapsed' in message: %s", message)
	}
}

func TestTimeUntilPassed(t *testing.T) {
	past := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	_, message := TimeUntil(past)

	if !contains(message, "passed") {
		t.Errorf("Expected 'passed' in message for past date: %s", message)
	}
}

func TestStartOfDay(t *testing.T) {
	start, err := StartOfDay()
	if err != nil {
		t.Fatalf("Failed to get start of day: %v", err)
	}

	if start.Hour() != 0 || start.Minute() != 0 || start.Second() != 0 {
		t.Errorf("Expected midnight, got %s", start.Format("15:04:05"))
	}
}

func TestEndOfDay(t *testing.T) {
	end, err := EndOfDay()
	if err != nil {
		t.Fatalf("Failed to get end of day: %v", err)
	}

	if end.Hour() != 23 || end.Minute() != 59 || end.Second() != 59 {
		t.Errorf("Expected 23:59:59, got %s", end.Format("15:04:05"))
	}
}

func TestWeekday(t *testing.T) {
	weekday, name := Weekday()

	if weekday != time.Monday {
		t.Errorf("Expected Monday, got %v (%s)", weekday, name)
	}
}

func TestParseTime(t *testing.T) {
	_, err := ParseRFC3339()
	if err != nil {
		t.Fatalf("Failed to parse: %v", err)
	}
}

func TestTimeTruncation(t *testing.T) {
	tm := time.Date(2024, 1, 15, 14, 30, 45, 123456789, time.UTC)

	truncated := tm.Truncate(time.Hour)
	if truncated.Minute() != 0 || truncated.Second() != 0 {
		t.Errorf("Expected hour truncation, got %s", truncated.Format("15:04:05"))
	}
}

func TestTimeRound(t *testing.T) {
	tm := time.Date(2024, 1, 15, 14, 30, 30, 500000000, time.UTC)

	rounded := tm.Round(time.Hour)
	if rounded.Minute() != 0 || rounded.Second() != 0 {
		t.Errorf("Expected hour rounding, got %s", rounded.Format("15:04:05"))
	}
}
