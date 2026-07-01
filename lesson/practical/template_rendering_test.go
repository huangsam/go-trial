package practical_test

import (
	"strconv"
	"testing"

	"github.com/huangsam/go-trial/lesson/practical"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// An example car to use for testing purposes
var exampleCar = &practical.Car{
	Make:       "Toyota",
	Model:      "Corolla",
	ModelYear:  2020,
	WheelCount: 4,
	MileCount:  15000,
}

// An example employee to use for testing purposes
var exampleEmployee = &practical.Employee{
	FirstName: "John",
	LastName:  "Doe",
	Age:       20,
	IsManager: false,
	Skills:    []string{"Java", "Python", "JavaScript"},
	Salary:    100000,
}

func TestRenderCarInfo(t *testing.T) {
	info, err := practical.RenderCarInfo(exampleCar)
	require.NoError(t, err)
	assert.NotNil(t, info)
	assert.Contains(t, info, exampleCar.Make)
	assert.Contains(t, info, exampleCar.Model)
	assert.Contains(t, info, "made in "+strconv.Itoa(exampleCar.ModelYear))
}

func TestRenderEmployeeInfo(t *testing.T) {
	info, err := practical.RenderEmployeeInfo(exampleEmployee)
	require.NoError(t, err)
	assert.NotNil(t, info)
	assert.Contains(t, info, exampleEmployee.FirstName)
	assert.Contains(t, info, exampleEmployee.LastName)
	for _, skill := range exampleEmployee.Skills {
		assert.Contains(t, info, skill)
	}
	assert.Contains(t, info, strconv.Itoa(exampleEmployee.Age))
	assert.Contains(t, info, "Role: Individual Contributor")
	assert.Contains(t, info, strconv.Itoa(exampleEmployee.Salary))
}
