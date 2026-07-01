package context

import (
	"context"
	"time"
)

// Define custom types for context keys to avoid collisions
type (
	grpcMethodKey  struct{}
	peerAddressKey struct{}
	traceIDKey     struct{}
)

// ClientSideContext demonstrates client-side context usage in gRPC.
func ClientSideContext() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// In a real gRPC client, you'd pass ctx to the call:
	// response, err := client.MyMethod(ctx, request)

	// Simulating with a simple check
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		return "context is valid for gRPC call", nil
	}
}

// ServerSideContext demonstrates server-side context usage in gRPC.
func ServerSideContext() (string, string) {
	ctx := context.WithValue(context.Background(), grpcMethodKey{}, "/mypackage.MyService/MyMethod")
	ctx = context.WithValue(ctx, peerAddressKey{}, "127.0.0.1:54321")

	method := ctx.Value(grpcMethodKey{}).(string)
	peerAddress := ctx.Value(peerAddressKey{}).(string)

	return method, peerAddress
}

// ContextPropagation shows how context flows from client to server.
func ContextPropagation() (string, bool) {
	// Client side - create context with metadata
	ctx := context.WithValue(context.Background(), traceIDKey{}, "abc-123")

	// In real scenario, this would be sent via gRPC metadata
	// Server receives and can access the same values

	serverCtx := ctx
	traceID := serverCtx.Value(traceIDKey{}).(string)

	return traceID, traceID == "abc-123"
}

// DeadlinePropagation demonstrates deadline propagation.
func DeadlinePropagation() (time.Time, time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	deadline, ok := ctx.Deadline()
	if !ok {
		return time.Time{}, 0
	}

	_ = 2 * time.Second // unused, for documentation
	actualDuration := time.Until(deadline)

	return deadline, actualDuration
}
