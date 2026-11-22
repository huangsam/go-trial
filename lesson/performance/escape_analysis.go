package performance

// EscapeRecord is a small value used to explore compiler escape diagnostics.
type EscapeRecord struct {
	Value int
}

// LocalRecordValue uses a pointer that stays within the function.
func LocalRecordValue(value int) int {
	record := &EscapeRecord{Value: value}
	return record.Value
}

// NewEscapeRecord returns a pointer that can outlive the function call.
func NewEscapeRecord(value int) *EscapeRecord {
	return &EscapeRecord{Value: value}
}
