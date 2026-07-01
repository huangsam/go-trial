# Time and Date Handling in Go

Go's `time` package provides comprehensive functionality for working with dates, times, and durations.

## The Reference Time

Go uses a specific reference time for formatting: **Mon Jan 2 15:04:05 MST 2006**

To create a format string, rearrange this date to match your desired output:
- `2006` = year
- `01` = month (zero-padded)
- `02` = day (zero-padded)
- `15` = hour (24-hour)
- `04` = minute
- `05` = second

## Common Patterns

### Parsing
```go
// RFC3339 format
t, _ := time.Parse(time.RFC3339, "2024-01-15T10:30:00Z")

// Custom format
t, _ := time.Parse("2006-01-02", "2024-01-15")
```

### Formatting
```go
now := time.Now()
fmt.Println(now.Format(time.RFC3339))  // 2024-01-15T10:30:45Z07:00
fmt.Println(now.Format("02/01/2006"))  // 15/01/2024
```

### Duration Parsing
```go
d, _ := time.ParseDuration("2h30m")  // 2 hours, 30 minutes
result := now.Add(d)                 // Add duration to time
```

## Key Functions

| Function | Description |
|----------|-------------|
| `time.Now()` | Get current time |
| `time.Parse(layout, value)` | Parse string to Time |
| `Time.Format(layout)` | Format Time as string |
| `Time.Add(d Duration)` | Add duration to time |
| `Time.Sub(t Time) Duration` | Subtract times |
| `time.Since(t Time) Duration` | Time since t |
| `time.Sleep(d Duration)` | Sleep for duration |

## Timezones

```go
// Load timezone
loc, _ := time.LoadLocation("America/New_York")

// Convert to timezone
t := time.Now().In(loc)
```

## Best Practices

1. **Always specify timezone** - Use `UTC` explicitly when appropriate
2. **Store as RFC3339** - Standard format for JSON/APIs
3. **Use Duration for calculations** - Not integers (handles DST, leap seconds)
4. **Compare with Before/After** - Not `<` or `>`
