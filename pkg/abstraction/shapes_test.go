package abstraction_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"io.huangsam/trial/pkg/abstraction"
)

// TestRectangleArea tests the Area method of the Rectangle struct
func TestRectangleArea(t *testing.T) {
	rect := abstraction.Rectangle{Width: 5, Height: 3}
	expectedArea := 15.0
	actualArea := rect.Area()
	assert.Equal(t, expectedArea, actualArea)
}

// TestRectanglePerimeter tests the Perimeter method of the Rectangle struct
func TestRectanglePerimeter(t *testing.T) {
	rect := abstraction.Rectangle{Width: 5, Height: 3}
	expectedPerimeter := 16.0
	actualPerimeter := rect.Perimeter()
	assert.Equal(t, expectedPerimeter, actualPerimeter)
}

// TestCircleArea tests the Area method of the Circle struct
func TestCircleArea(t *testing.T) {
	circle := abstraction.Circle{Radius: 2}
	expectedArea := math.Pi * 4
	actualArea := circle.Area()
	assert.Equal(t, expectedArea, actualArea)
}

// TestCirclePerimeter tests the Perimeter method of the Circle struct
func TestCirclePerimeter(t *testing.T) {
	circle := abstraction.Circle{Radius: 2}
	expectedPerimeter := 4 * math.Pi
	actualPerimeter := circle.Perimeter()
	assert.Equal(t, expectedPerimeter, actualPerimeter)
}
