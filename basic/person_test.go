package basic

import (
	"testing"
)

func TestGetInfo(t *testing.T) {
	var expectedName string = "John"
	var expectedAge int32 = 12
	person := Person{expectedName, expectedAge}

	actualName := person.Name
	if actualName != expectedName {
		t.Errorf("Expected name %s but got %s", expectedName, actualName)
	}

	actualAge := GetAge(person)
	if actualAge != expectedAge {
		t.Errorf("Expected age %d but got %d", expectedAge, actualAge)
	}

	isOlderThanSeventeen := IsOlderThan(person, 17)
	if isOlderThanSeventeen {
		t.Errorf("Expected a young one, but got someone with age %d", GetAge(person))
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

func BenchmarkIsOlderThan(b *testing.B) {
	person := Person{"Jerry", 0}
	for i := 0; i < b.N; i++ {
		IsOlderThan(person, 18)
	}
}
