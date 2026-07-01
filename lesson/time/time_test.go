package time

import (
	"testing"
	"time"
)

func TestCurrentTime(t *testing.T) {
	result := CurrentTime()

	parsed, err := time.Parse(time.RFC3339, result)
	if err != nil {
		t.Errorf("Invalid RFC3339 format: %s", result)
	}

	now := time.Now()
	diff := now.Sub(parsed)
	if diff < -time.Second || diff > time.Second {
		t.Errorf("Time difference too large: %v", diff)
	}
}

func TestParseRFC3339(t *testing.T) {
	parsed, err := ParseRFC3339()
	if err != nil {
		t.Fatalf("Failed to parse RFC3339: %v", err)
	}

	expectedYear := 2024
	if parsed.Year() != expectedYear {
		t.Errorf("Expected year %d, got %d", expectedYear, parsed.Year())
	}
}

func TestParseCustomFormat(t *testing.T) {
	dateStr := "2024-01-15"
	parsed, err := ParseCustomFormat(dateStr)
	if err != nil {
		t.Fatalf("Failed to parse custom format: %v", err)
	}

	if parsed.Year() != 2024 || parsed.Month() != time.January || parsed.Day() != 15 {
		t.Errorf("Unexpected date: %s", parsed.Format("2006-01-02"))
	}
}

func TestDurationParsing(t *testing.T) {
	d, err := DurationParsing()
	if err != nil {
		t.Fatalf("Failed to parse duration: %v", err)
	}

	expectedSeconds := 2*3600 + 30*60 + 15
	actualSeconds := int(d.Seconds())
	if actualSeconds != expectedSeconds {
		t.Errorf("Expected %d seconds, got %d", expectedSeconds, actualSeconds)
	}
}

func TestTimeComparison(t *testing.T) {
	t1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	result1 := TimeComparison(t1, t2)
	if result1 != "t1 is before t2" {
		t.Errorf("Expected 't1 is before t2', got '%s'", result1)
	}

	result2 := TimeComparison(t2, t1)
	if result2 != "t1 is after t2" {
		t.Errorf("Expected 't1 is after t2', got '%s'", result2)
	}

	result3 := TimeComparison(t1, t3)
	if result3 != "t1 equals t2" {
		t.Errorf("Expected 't1 equals t2', got '%s'", result3)
	}
}

func TestTimeArithmetic(t *testing.T) {
	_, duration := TimeArithmetic()

	expectedDuration := 5 * time.Minute
	if duration < expectedDuration-time.Second || duration > expectedDuration+time.Second {
		t.Errorf("Expected ~%v duration, got %v", expectedDuration, duration)
	}
}

func TestSleepDemonstration(t *testing.T) {
	result := SleepDemonstration()

	if !contains(result, "Slept for") {
		t.Errorf("Expected 'Slept for' in result: %s", result)
	}
}

func TestUnixTimestamps(t *testing.T) {
	seconds, milli := UnixTimestamps()

	expectedMilli := seconds * 1000
	if milli < expectedMilli || milli > expectedMilli+1000 {
		t.Errorf("Milliseconds not consistent with seconds: %d vs %d", milli, seconds)
	}

	if seconds <= 0 {
		t.Errorf("Expected positive Unix timestamp, got %d", seconds)
	}
}

func TestTimeAdd(t *testing.T) {
	// Use a fixed date to avoid DST issues
	base := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)
	added := base.Add(24 * time.Hour)

	if added.Day() != 16 {
		t.Errorf("Expected day to change from 15 to 16, got %d", added.Day())
	}
}

func TestTimeSubtract(t *testing.T) {
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)

	diff := now.Sub(yesterday)
	expectedDiff := 24 * time.Hour

	if diff < expectedDiff-time.Second || diff > expectedDiff+time.Second {
		t.Errorf("Expected ~%v difference, got %v", expectedDiff, diff)
	}
}

func TestTimeParsing(t *testing.T) {
	_, err := ParseRFC3339()
	if err != nil {
		t.Fatalf("ParseRFC3339 failed: %v", err)
	}
}
