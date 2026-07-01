package concurrency

import (
	"context"
	"strings"
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
	if !strings.Contains(result, expectedPattern) {
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
		if !strings.Contains(result, expected) {
			t.Errorf("Expected result to contain '%s', got '%s'", expected, result)
		}
	}
}

func TestNestedCancel(t *testing.T) {
	parentResult, _ := NestedCancel()

	// When parent is cancelled, child should also be cancelled
	if !strings.Contains(parentResult, "parent cancelled") {
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
