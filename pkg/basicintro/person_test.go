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
	assert.Equal(t, expectedName, actualName)

	actualAge := person.Age
	assert.Equal(t, expectedAge, actualAge)

	isOlderThanSeventeen := person.IsOlderThan(17)
	assert.False(t, isOlderThanSeventeen)
}

func BenchmarkIsOlderThan(b *testing.B) {
	person := basicintro.Person{"Jerry", 0}
	for i := 0; i < b.N; i++ {
		person.IsOlderThan(18)
	}
}
