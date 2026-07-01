package json

import (
	"encoding/json"
)

// BasicMarshal demonstrates basic JSON marshaling.
func BasicMarshal() (string, error) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := Person{Name: "Alice", Age: 30}
	data, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// BasicUnmarshal demonstrates basic JSON unmarshaling.
func BasicUnmarshal() (string, int, error) {
	jsonData := `{"name": "Bob", "age": 25}`

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var p Person
	err := json.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		return "", 0, err
	}

	return p.Name, p.Age, nil
}

// NestedStruct demonstrates marshaling nested structures.
func NestedStruct() (string, error) {
	type Address struct {
		City    string `json:"city"`
		Country string `json:"country"`
	}

	type Person struct {
		Name    string  `json:"name"`
		Age     int     `json:"age"`
		Address Address `json:"address"`
	}

	p := Person{
		Name:    "Charlie",
		Age:     35,
		Address: Address{City: "New York", Country: "USA"},
	}

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// SliceOfStructs demonstrates marshaling slices.
func SliceOfStructs() (string, error) {
	type Product struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	}

	products := []Product{
		{Name: "Apple", Price: 100},
		{Name: "Banana", Price: 80},
		{Name: "Cherry", Price: 200},
	}

	data, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Map demonstrates marshaling maps.
func Map() (string, error) {
	data := map[string]int{
		"apple":  100,
		"banana": 80,
		"cherry": 200,
	}

	result, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
