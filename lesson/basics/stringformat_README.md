# String Formatting in Go

Go provides multiple ways to format and concatenate strings. The right choice depends on your use case.

## Methods Overview

| Method | Best For | Performance |
|--------|----------|-------------|
| `fmt.Sprintf` | One-off formatting, small strings | Good for single calls |
| `strings.Builder` | Building strings incrementally | Best overall |
| `bytes.Buffer` | Similar to Builder, but works with bytes | Slightly faster than Builder |
| `strings.Join` | Joining slice elements | Fast and clean |
| `+ operator` | Simple concatenation (2-3 parts) | Fine for small cases |

## Examples

### fmt.Sprintf
```go
result := fmt.Sprintf("Name: %s, Age: %d", name, age)
```

### strings.Builder (Recommended for loops)
```go
var builder strings.Builder
for _, word := range words {
    if builder.Len() > 0 {
        builder.WriteString(", ")
    }
    builder.WriteString(word)
}
result := builder.String()
```

### bytes.Buffer
```go
var buf bytes.Buffer
buf.WriteString("Hello")
buf.WriteString(" ")
buf.WriteString("World")
result := buf.String()
```

### strings.Join (Best for slices)
```go
result := strings.Join([]string{"a", "b", "c"}, ", ")
// Result: "a, b, c"
```

## When to Use Each

1. **fmt.Sprintf** - Use when you need formatted output with specific patterns
2. **strings.Builder** - Use for building large strings in loops or incrementally
3. **bytes.Buffer** - Alternative to Builder when you already work with bytes
4. **strings.Join** - Always prefer this for joining slice elements

## Key Methods

- `builder.WriteString(s)` - Append string
- `builder.String()` - Get final string
- `builder.Len()` - Current length
- `builder.Reset()` - Clear builder

## Performance Tip

In benchmarks, `strings.Builder` and `bytes.Buffer` outperform string concatenation with `+` by 10x or more when building many strings in a loop.
