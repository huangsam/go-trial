package basicintro_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/basicintro"
	"github.com/stretchr/testify/assert"
)

func TestPersonFields(t *testing.T) {
	var expectedName string = "John"
	var expectedAge int = 12
	person := basicintro.Person{expectedName, expectedAge}
	assert.Equal(t, expectedName, person.Name)
	assert.Equal(t, expectedAge, person.Age)
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
			person := basicintro.Person{"Jerry", tt.age}
			assert.Equal(t, tt.expected, person.IsOlderThan(tt.compare))
		})
	}
}
