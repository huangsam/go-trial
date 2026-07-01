package concurrency

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// HandlerWithCancellation demonstrates a HTTP handler that respects context cancellation.
func HandlerWithCancellation() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Simulate processing
	resultChan := make(chan string, 1)

	go func() {
		time.Sleep(500 * time.Millisecond) // Slow operation
		resultChan <- "processing complete"
	}()

	select {
	case result := <-resultChan:
		fmt.Printf("Result: %s\n", result)
	case <-ctx.Done():
		fmt.Println("Request cancelled")
	}
}

// CancelableRequest demonstrates making an HTTP request with cancellation.
func CancelableRequest() error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://example.com", nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		// Expected to get a timeout or cancellation error
		fmt.Printf("Request failed as expected: %v\n", err)
	}

	return err
}

// ContextInMiddleware demonstrates using context for middleware-like patterns.
func ContextInMiddleware() (string, string) {
	ctx := context.WithValue(context.Background(), requestIDKey{}, "req-789")

	// Simulate middleware processing
	ctx = context.WithValue(ctx, timestampKey{}, time.Now().Unix())

	requestID := ctx.Value(requestIDKey{}).(string)
	timestamp := ctx.Value(timestampKey{}).(int64)

	return requestID, fmt.Sprintf("%d", timestamp)
}

// GracefulShutdown demonstrates using context for graceful server shutdown.
func GracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Simulate running operations
	done := make(chan struct{})
	go func() {
		defer close(done)
		time.Sleep(2 * time.Second) // Simulate ongoing work
	}()

	// Wait for operation or timeout
	select {
	case <-done:
		fmt.Println("Operations completed gracefully")
	case <-ctx.Done():
		fmt.Printf("Shutdown timed out: %v\n", ctx.Err())
	}
}
