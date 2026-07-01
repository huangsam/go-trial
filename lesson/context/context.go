package context

import (
	"context"
	"fmt"
	"time"
)

// Define custom types for context keys to avoid collisions
type (
	requestIDKey     struct{}
	userIDKey        struct{}
	timestampKey     struct{}
	stringKeyType    struct{}
	intKeyType       struct{}
	collisionKeyType struct{}
)

// BackgroundContext returns a background context.
// This is the base context that all other contexts derive from.
func BackgroundContext() context.Context {
	return context.Background()
}

// CancellableContext demonstrates context cancellation.
// When you cancel the context, the goroutine stops processing.
func CancellableContext() (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resultChan := make(chan string, 1)

	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				resultChan <- fmt.Sprintf("stopped at iteration %d", i)
				return
			default:
				time.Sleep(50 * time.Millisecond)
			}
		}
		resultChan <- "completed all iterations"
	}()

	// Let it run for a bit, then cancel
	time.Sleep(150 * time.Millisecond)
	cancel()

	return <-resultChan, nil
}

// TimeoutContext demonstrates context with timeout.
// If the operation takes too long, it's cancelled automatically.
func TimeoutContext() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	resultChan := make(chan string, 1)

	go func() {
		time.Sleep(300 * time.Millisecond) // Simulates slow operation
		resultChan <- "operation completed"
	}()

	select {
	case result := <-resultChan:
		return result, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// ValueContext demonstrates passing values through context.
// Use this sparingly - only for request-scoped data like IDs or auth tokens.
func ValueContext() string {
	ctx := context.WithValue(context.Background(), requestIDKey{}, "req-123")
	ctx = context.WithValue(ctx, userIDKey{}, "user-456")

	requestID := ctx.Value(requestIDKey{})
	userID := ctx.Value(userIDKey{})

	return fmt.Sprintf("request_id=%v, user_id=%v", requestID, userID)
}

// NestedCancel demonstrates that canceling parent cancels children.
func NestedCancel() (string, string) {
	parentCtx, parentCancel := context.WithCancel(context.Background())
	defer parentCancel()

	childCtx, childCancel := context.WithCancel(parentCtx)
	defer childCancel()

	go func() {
		time.Sleep(100 * time.Millisecond)
		parentCancel()
	}()

	select {
	case <-parentCtx.Done():
		return "parent cancelled", "child also cancelled"
	case <-childCtx.Done():
		return "child cancelled", ""
	}
}

// PropagationWarning demonstrates why context should not be stored in structs.
// This is an anti-pattern!
type AntiPattern struct {
	// DON'T do this - context should be passed as parameter
}

// Good pattern: pass context as first parameter
func goodPattern(ctx context.Context) error {
	// Use ctx for cancellation/timeout checks
	<-ctx.Done()
	return ctx.Err()
}
