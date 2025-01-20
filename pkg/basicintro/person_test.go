package basicintro_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/basicintro"
	"github.com/stretchr/testify/assert"
)

func TestGetInfo(t *testing.T) {
	var expectedName string = "John"
	var expectedAge int = 12
	person := basicintro.Person{expectedName, expectedAge}
	assert.Equal(t, expectedName, person.Name)
	assert.Equal(t, expectedAge, person.Age)
	assert.False(t, person.IsOlderThan(17))
}

func BenchmarkIsOlderThan(b *testing.B) {
	person := basicintro.Person{"Jerry", 0}
	for i := 0; i < b.N; i++ {
		person.IsOlderThan(18)
	}
}
