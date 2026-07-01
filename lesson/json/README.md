# Working with JSON in Go

Go's `encoding/json` package provides robust support for JSON serialization and deserialization.

## Basic Usage

### Marshaling (Go → JSON)
```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

p := Person{Name: "Alice", Age: 30}
data, err := json.Marshal(p)  // or json.MarshalIndent for pretty print
```

### Unmarshaling (JSON → Go)
```go
var p Person
err := json.Unmarshal(jsonData, &p)
```

## Struct Tags

- **`json:"name"`** - Rename the field in JSON
- **`json:"-"`** - Exclude from JSON
- **`omitempty`** - Omit if empty (zero value)

```go
type Config struct {
    Host     string `json:"host,omitempty"`
    Password string `json:"password,omitempty"`
}
```

## Advanced Patterns

### Custom Marshalers
Implement `MarshalJSON()` and `UnmarshalJSON()` methods for custom behavior.

### json.RawMessage
Defers marshaling/unmarshaling of nested fields.

```go
type Message struct {
    Type    string          `json:"type"`
    Payload json.RawMessage `json:"payload"`  // Unmarshals to []byte
}
```

### json.Number
Preserves large numbers that exceed int64/float64 precision.

## Best Practices

1. **Always check errors** - JSON operations can fail on malformed input
2. **Use `omitempty` carefully** - Only when missing field has different meaning than zero value
3. **Prefer structs over map[string]interface{}** for known schemas
4. **Consider using pointers** for optional fields (enables omitempty to work)
5. **JSON numbers become float64** when unmarshaling into interface{}

## Common Pitfalls

- **No private field support** - Only exported fields are marshaled
- **Map keys must be strings** - Non-string map keys aren't supported
- **Order not preserved** - JSON maps don't guarantee iteration order
