package concurrency_test

import (
	"testing"

	"io.huangsam/trial/pkg/concurrency"
)

func TestGetAnswersWithWaitGroup(t *testing.T) {
	// Get the results from the function
	results := concurrency.GetAnswersWithWaitGroup()

	// Compare the results with the expected values
	if results[0] != 0 {
		t.Errorf("Expected 0 at index 0 but got %d", results[0])
	}
	if results[99] != 198 {
		t.Errorf("Expected 198 at index 0 but got %d", results[99])
	}
}
