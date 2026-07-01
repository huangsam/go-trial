package json

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestBasicMarshal(t *testing.T) {
	result, err := BasicMarshal()
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Verify the JSON contains expected fields
	if !strings.Contains(result, `"name"`) {
		t.Errorf("Expected 'name' field in JSON: %s", result)
	}
	if !strings.Contains(result, `"age"`) {
		t.Errorf("Expected 'age' field in JSON: %s", result)
	}
}

func TestBasicUnmarshal(t *testing.T) {
	name, age, err := BasicUnmarshal()
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if name != "Bob" {
		t.Errorf("Expected name 'Bob', got '%s'", name)
	}
	if age != 25 {
		t.Errorf("Expected age 25, got %d", age)
	}
}

func TestNestedStruct(t *testing.T) {
	result, err := NestedStruct()
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Verify indentation (pretty printing)
	if !strings.Contains(result, "\n") {
		t.Error("Expected pretty-printed JSON with newlines")
	}
	if !strings.Contains(result, `"city"`) {
		t.Errorf("Expected 'city' field in nested JSON: %s", result)
	}
}

func TestSliceOfStructs(t *testing.T) {
	result, err := SliceOfStructs()
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	if !strings.HasPrefix(result, "[") {
		t.Errorf("Expected array output, got: %s", result)
	}
}

func TestMapMarshal(t *testing.T) {
	result, err := Map()
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var data map[string]int
	err = json.Unmarshal([]byte(result), &data)
	if err != nil {
		t.Fatalf("Could not unmarshal result: %v", err)
	}

	if data["apple"] != 100 {
		t.Errorf("Expected apple price 100, got %d", data["apple"])
	}
}

func TestCustomStructTags(t *testing.T) {
	type Item struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Hidden string `json:"-"`
	}

	item := Item{
		ID:     1,
		Title:  "Test",
		Hidden: "secret",
	}

	data, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	result := string(data)

	// Hidden field should not be in output
	if strings.Contains(result, "Hidden") || strings.Contains(result, "secret") {
		t.Errorf("Hidden field should not be in JSON: %s", result)
	}
}

func TestOmitEmpty(t *testing.T) {
	type Config struct {
		Host     string `json:"host,omitempty"`
		Port     int    `json:"port,omitempty"`
		Password string `json:"password,omitempty"`
	}

	c := Config{Host: "localhost", Port: 8080}
	data, err := json.Marshal(c)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	result := string(data)

	// password should be omitted since it's empty
	if strings.Contains(result, "password") {
		t.Errorf("Empty field 'password' should be omitted: %s", result)
	}
}

func TestUnmarshalToInterface(t *testing.T) {
	jsonData := `{"name": "Test", "count": 42}`

	var data interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// The result is a map[string]interface{}
	resultMap, ok := data.(map[string]interface{})
	if !ok {
		t.Errorf("Expected map[string]interface{}, got %T", data)
	}

	_ = resultMap["name"]
	_ = resultMap["count"]
}
