package time

import (
	"strings"
)

func contains(s, substr string) bool {
	return len(s) >= len(substr) && strings.Contains(s, substr)
}
