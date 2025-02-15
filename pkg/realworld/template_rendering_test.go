package realworld_test

import (
	"strconv"
	"testing"

	"github.com/huangsam/go-trial/pkg/realworld"
	"github.com/stretchr/testify/assert"
)

// An example car to use for testing purposes
var exampleCar = &realworld.Car{
	Make:       "Toyota",
	Model:      "Corolla",
	ModelYear:  2020,
	WheelCount: 4,
	MileCount:  15000,
}

func TestGetCarSimpleInfo(t *testing.T) {
	info, err := realworld.GetCarSimpleInfo(exampleCar)
	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Contains(t, info, exampleCar.Make)
	assert.Contains(t, info, exampleCar.Model)
	assert.Contains(t, info, "made in "+strconv.Itoa(exampleCar.ModelYear))
}
