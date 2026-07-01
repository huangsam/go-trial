# Context: Managing Cancellation, Timeouts, and Values

In Go, `context.Context` is a fundamental type for managing request-scoped values, cancellation signals, and deadlines across API boundaries and goroutines.

## The Four Main Functions

1. **`context.Background()`** - Root context, typically used in main functions
2. **`context.WithCancel(parent)`** - Creates a cancellable context
3. **`context.WithTimeout(parent, duration)`** - Creates a context with automatic cancellation after timeout
4. **`context.WithValue(parent, key, val)`** - Adds key-value pairs to the context

## Key Principles

### 1. Always Pass Context as First Parameter
```go
func DoWork(ctx context.Context, data string) error {
    // Use ctx for cancellation checks
}
```

### 2. Never Store Context in Structs
Context is meant to be passed through function calls, not stored:
```go
// Anti-pattern!
type Service struct {
    ctx context.Context  // DON'T do this
}

// Good pattern - pass as parameter
func (s *Service) DoWork(ctx context.Context) error { ... }
```

### 3. Don't Forget to Cancel
Use `defer cancel()` immediately after creating a cancellable context:
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel() // Always call cancel to release resources
```

## Common Patterns

### Cancellation
Use when you need to stop an operation based on external signals:
```go
ctx, cancel := context.WithCancel(parentCtx)
defer cancel()

// Later, somewhere else in your code:
cancel() // Signals all goroutines using ctx to stop
```

### Timeouts
Use for operations with deadline constraints:
```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

res, err := doWork(ctx)
if err == context.DeadlineExceeded {
    // Handle timeout
}
```

### Request-Scoped Values
Use sparingly for request IDs, auth tokens, etc.:
```go
ctx = context.WithValue(ctx, "request_id", req.ID)
// Access with: ctx.Value("request_id")
```

## Best Practices

1. **Don't pass nil context** - Use `context.Background()` instead
2. **Use standard keys for well-known values** - Don't create new key types unnecessarily
3. **Check context.Done() in long-running operations** - Don't ignore cancellation signals
4. **Cancel when done** - Always call the cancel function to release resources
5. **Values should be immutable** - Don't modify values stored in context

## Related Resources

- [Go Context Blog Post](https://go.dev/blog/context)
- [Go Context Package Docs](https://pkg.go.dev/context)
