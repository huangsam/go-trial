package basicintro_test

import (
	"testing"

	"github.com/huangsam/go-trial/lesson/basicintro"
	"github.com/stretchr/testify/assert"
)

var sampleAddress = basicintro.Address{Street: "1234 Magic Avenue", Zip: "12345-6789"}

func TestPersonFields(t *testing.T) {
	expectedName := "John"
	expectedAge := 12
	expectedStreet := "1234 Magic Avenue"
	expectedZip := "12345-6789"
	person := basicintro.Person{expectedName, expectedAge, sampleAddress}
	assert.Equal(t, expectedName, person.Name)     // Normal access
	assert.Equal(t, expectedAge, person.Age)       // Normal access
	assert.Equal(t, expectedStreet, person.Street) // Embedded access
	assert.Equal(t, expectedZip, person.Zip)       // Embedded access
}

func TestPerson_IsOlderThan(t *testing.T) {
	tests := []struct {
		name     string
		age      int
		compare  int
		expected bool
	}{
		{"Less", 17, 18, false},
		{"Equal", 18, 18, false},
		{"Greater", 19, 18, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			person := basicintro.Person{"Jerry", tt.age, sampleAddress}
			assert.Equal(t, tt.expected, person.IsOlderThan(tt.compare))
		})
	}
}
