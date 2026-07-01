package context

import (
	"context"
	"testing"
	"time"
)

func TestClientSideContext(t *testing.T) {
	result, err := ClientSideContext()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := "context is valid for gRPC call"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestServerSideContext(t *testing.T) {
	method, peerAddress := ServerSideContext()

	expectedMethod := "/mypackage.MyService/MyMethod"
	if method != expectedMethod {
		t.Errorf("Expected method '%s', got '%s'", expectedMethod, method)
	}

	expectedPeer := "127.0.0.1:54321"
	if peerAddress != expectedPeer {
		t.Errorf("Expected peer address '%s', got '%s'", expectedPeer, peerAddress)
	}
}

func TestContextPropagation(t *testing.T) {
	traceID, isValid := ContextPropagation()

	expectedID := "abc-123"
	if traceID != expectedID {
		t.Errorf("Expected trace ID '%s', got '%s'", expectedID, traceID)
	}

	if !isValid {
		t.Error("Context propagation should be valid")
	}
}

func TestDeadlinePropagation(t *testing.T) {
	deadline, duration := DeadlinePropagation()

	// Duration should be approximately 2 seconds (allowing for test execution time)
	minExpected := 1900 * time.Millisecond
	maxExpected := 3000 * time.Millisecond

	if duration < minExpected || duration > maxExpected {
		t.Errorf("Duration %v not in expected range [%v, %v]", duration, minExpected, maxExpected)
	}

	// Deadline should be in the future
	if deadline.IsZero() {
		t.Error("Deadline should not be zero")
	}
}

func TestContextWithValueTypes(t *testing.T) {
	ctx := context.WithValue(context.Background(), stringKeyType{}, "value1")
	ctx = context.WithValue(ctx, intKeyType{}, 42)

	stringVal := ctx.Value(stringKeyType{}).(string)
	intVal := ctx.Value(intKeyType{}).(int)
	if stringVal != "value1" {
		t.Errorf("Expected 'value1', got '%s'", stringVal)
	}

	if intVal != 42 {
		t.Errorf("Expected 42, got %d", intVal)
	}
}

func TestContextKeyCollision(t *testing.T) {
	ctx := context.WithValue(context.Background(), collisionKeyType{}, "first")
	ctx = context.WithValue(ctx, collisionKeyType{}, "second") // Overwrites previous

	val := ctx.Value(collisionKeyType{}).(string)

	if val != "second" {
		t.Errorf("Expected 'second', got '%s'", val)
	}
}
