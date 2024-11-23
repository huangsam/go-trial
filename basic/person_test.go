package basic

import (
	"testing"
)

func TestGetNameDirectly(t *testing.T) {
	var expected string = "Sally"
	person := Person{expected, 0}
	actual := person.Name
	if actual != expected {
		t.Errorf("Expected name %s but got %s", expected, actual)
	}
}

func TestGetAge(t *testing.T) {
	var expected int32 = 12
	person := Person{"John", expected}
	actual := GetAge(person)
	if actual != expected {
		t.Errorf("Expected age %d but got %d", expected, actual)
	}
}

func TestSetAge(t *testing.T) {
	var expected int32 = 24
	person := Person{"Mary", 0}
	SetAge(&person, expected)
	actual := GetAge(person)
	if actual != expected {
		t.Errorf("Expected age %d but got %d", expected, actual)
	}
}

func BenchmarkGetAge(b *testing.B) {
	person := Person{"Bob", 0}
	for i := 0; i < b.N; i++ {
		GetAge(person)
	}
}
