package json

import (
	"encoding/json"
	"fmt"
)

// CustomMarshaler demonstrates implementing json.Marshaler interface.
func CustomMarshaler() (string, error) {
	type Password string

	type User struct {
		Username string   `json:"username"`
		Password Password `json:"password"`
	}

	u := User{
		Username: "alice",
		Password: Password("secret123"),
	}

	data, err := json.Marshal(u)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// StreamingMarshal demonstrates streaming JSON marshaling.
func StreamingMarshal() error {
	type Item struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	items := []Item{
		{Name: "A", Value: 1},
		{Name: "B", Value: 2},
		{Name: "C", Value: 3},
	}

	// Marshal to a writer
	var result string
	data, err := json.Marshal(items)
	if err != nil {
		return err
	}
	result = string(data)

	fmt.Println(result)
	return nil
}

// RawMessage demonstrates using json.RawMessage for deferred marshaling.
func RawMessage() (string, error) {
	type Message struct {
		Type    string          `json:"type"`
		Payload json.RawMessage `json:"payload"`
	}

	msg := Message{
		Type:    "user",
		Payload: []byte(`{"name": "Bob", "age": 30}`),
	}

	data, err := json.MarshalIndent(msg, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// UnmarshalUnknownStructure demonstrates unmarshaling into interface{}.
func UnmarshalUnknownStructure() (map[string]interface{}, error) {
	jsonData := `{"name": "Test", "count": 42, "active": true}`

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NumberParsing demonstrates handling json.Number for large integers.
func NumberParsing() error {
	jsonData := `{"id": 9007199254740993}` // larger than math.MaxInt64

	type Data struct {
		ID json.Number `json:"id"`
	}

	var d Data
	err := json.Unmarshal([]byte(jsonData), &d)
	if err != nil {
		return err
	}

	fmt.Println("ID as string:", d.ID.String())
	return nil
}
