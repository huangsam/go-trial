package json

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestCustomMarshaler(t *testing.T) {
	result, err := CustomMarshaler()
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	type User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var u User
	err = json.Unmarshal([]byte(result), &u)
	if err != nil {
		t.Fatalf("Could not unmarshal: %v", err)
	}

	if u.Username != "alice" {
		t.Errorf("Expected username 'alice', got '%s'", u.Username)
	}
}

func TestStreamingMarshal(t *testing.T) {
	err := StreamingMarshal()
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
}

func TestRawMessage(t *testing.T) {
	result, err := RawMessage()
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Verify the output structure
	if !strings.Contains(result, `"type"`) {
		t.Errorf("Expected 'type' field: %s", result)
	}
	if !strings.Contains(result, `"payload"`) {
		t.Errorf("Expected 'payload' field: %s", result)
	}
}

func TestUnmarshalUnknownStructure(t *testing.T) {
	data, err := UnmarshalUnknownStructure()
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	name, ok := data["name"].(string)
	if !ok || name != "Test" {
		t.Errorf("Expected name 'Test', got '%v'", data["name"])
	}

	count, ok := data["count"].(float64) // JSON numbers are float64 by default
	if !ok || int(count) != 42 {
		t.Errorf("Expected count 42, got '%v'", data["count"])
	}
}

func TestNumberParsing(t *testing.T) {
	err := NumberParsing()
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
}

func TestUnmarshalArrayIntoInterface(t *testing.T) {
	jsonData := `[1, 2, 3]`

	var data interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// Arrays become []interface{} in Go
	arr, ok := data.([]interface{})
	if !ok {
		t.Errorf("Expected []interface{}, got %T", data)
	}
	if len(arr) != 3 {
		t.Errorf("Expected array of length 3, got %d", len(arr))
	}
}

func TestMarshalIndent(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := Person{Name: "Test", Age: 25}
	data, err := json.MarshalIndent(p, ">>> ", "---")
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	result := string(data)

	// Verify indentation is applied
	if !strings.Contains(result, ">>> ") {
		t.Errorf("Expected custom indent '>>> ': %s", result)
	}
}
