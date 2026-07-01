package context

import (
	"context"
	"testing"
	"time"
)

func TestBackgroundContext(t *testing.T) {
	ctx := BackgroundContext()
	if ctx == nil {
		t.Error("Background() should never return nil")
	}
}

func TestCancellableContext(t *testing.T) {
	result, err := CancellableContext()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedPattern := "stopped at iteration"
	if !contains(result, expectedPattern) {
		t.Errorf("Expected result to contain '%s', got '%s'", expectedPattern, result)
	}
}

func TestTimeoutContext(t *testing.T) {
	result, err := TimeoutContext()

	// Should get a timeout error since we sleep 300ms but timeout is 200ms
	if err == nil {
		t.Errorf("Expected timeout error, got nil")
	}

	expectedErr := context.DeadlineExceeded
	if err != expectedErr {
		t.Errorf("Expected '%v', got '%v'", expectedErr, err)
	}

	if result != "" {
		t.Errorf("Expected empty result on timeout, got '%s'", result)
	}
}

func TestValueContext(t *testing.T) {
	result := ValueContext()

	expectedValues := []string{"request_id=req-123", "user_id=user-456"}
	for _, expected := range expectedValues {
		if !contains(result, expected) {
			t.Errorf("Expected result to contain '%s', got '%s'", expected, result)
		}
	}
}

func TestNestedCancel(t *testing.T) {
	parentResult, _ := NestedCancel()

	// When parent is cancelled, child should also be cancelled
	if !contains(parentResult, "parent cancelled") {
		t.Errorf("Expected 'parent cancelled', got '%s'", parentResult)
	}
}

func TestPropagationGoodPattern(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := goodPattern(ctx)

	if err != context.DeadlineExceeded && err != context.Canceled {
		t.Errorf("Expected deadline or cancellation error, got %v", err)
	}
}

// Helper function to check if string contains substring
func contains(haystack, needle string) bool {
	return len(haystack) >= len(needle) &&
		(haystack == needle ||
			len(haystack) > len(needle) &&
				(haystack[:len(needle)] == needle || haystack[len(haystack)-len(needle):] == needle ||
					containsSubstring(haystack, needle)))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
