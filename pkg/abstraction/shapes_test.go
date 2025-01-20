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
	assert.Equal(t, 15.0, rect.Area())
}

// TestRectanglePerimeter tests the Perimeter method of the Rectangle struct
func TestRectanglePerimeter(t *testing.T) {
	rect := abstraction.Rectangle{Width: 5, Height: 3}
	assert.Equal(t, 16.0, rect.Perimeter())
}

// TestCircleArea tests the Area method of the Circle struct
func TestCircleArea(t *testing.T) {
	circle := abstraction.Circle{Radius: 2}
	assert.Equal(t, math.Pi*4, circle.Area())
}

// TestCirclePerimeter tests the Perimeter method of the Circle struct
func TestCirclePerimeter(t *testing.T) {
	circle := abstraction.Circle{Radius: 2}
	assert.Equal(t, 4*math.Pi, circle.Perimeter())
}
