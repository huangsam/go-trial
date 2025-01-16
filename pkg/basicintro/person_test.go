package basicintro_test

import (
	"testing"

	"io.huangsam/trial/pkg/basicintro"
)

func TestGetInfo(t *testing.T) {
	var expectedName string = "John"
	var expectedAge int32 = 12
	person := basicintro.Person{expectedName, expectedAge}

	actualName := person.Name
	if actualName != expectedName {
		t.Errorf("Expected name %s but got %s", expectedName, actualName)
	}

	actualAge := person.Age
	if actualAge != expectedAge {
		t.Errorf("Expected age %d but got %d", expectedAge, actualAge)
	}

	isOlderThanSeventeen := person.IsOlderThan(17)
	if isOlderThanSeventeen {
		t.Errorf("Expected a young one, but got someone with age %d", actualAge)
	}
}

func BenchmarkIsOlderThan(b *testing.B) {
	person := basicintro.Person{"Jerry", 0}
	for i := 0; i < b.N; i++ {
		person.IsOlderThan(18)
	}
}
