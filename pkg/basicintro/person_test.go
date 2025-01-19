package basicintro_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"io.huangsam/trial/pkg/basicintro"
)

func TestGetInfo(t *testing.T) {
	var expectedName string = "John"
	var expectedAge int = 12
	person := basicintro.Person{expectedName, expectedAge}

	actualName := person.Name
	assert.Equal(t, expectedName, actualName, "Expected name %s but got %s", expectedName, actualName)

	actualAge := person.Age
	assert.Equal(t, expectedAge, actualAge, "Expected age %d but got %d", expectedAge, actualAge)

	isOlderThanSeventeen := person.IsOlderThan(17)
	assert.False(t, isOlderThanSeventeen, "Expected a young one, but got someone with age %d", actualAge)
}

func BenchmarkIsOlderThan(b *testing.B) {
	person := basicintro.Person{"Jerry", 0}
	for i := 0; i < b.N; i++ {
		person.IsOlderThan(18)
	}
}
