package abstraction_test

import (
	"math"
	"testing"

	"github.com/huangsam/go-trial/pkg/abstraction"
	"github.com/stretchr/testify/assert"
)

func TestRectangleArea(t *testing.T) {
	rect := abstraction.Rectangle{Width: 5, Height: 3}
	assert.Equal(t, 15.0, rect.Area())
}

func TestRectanglePerimeter(t *testing.T) {
	rect := abstraction.Rectangle{Width: 5, Height: 3}
	assert.Equal(t, 16.0, rect.Perimeter())
}

func TestCircleArea(t *testing.T) {
	circle := abstraction.Circle{Radius: 2}
	assert.Equal(t, math.Pi*4, circle.Area())
}

func TestCirclePerimeter(t *testing.T) {
	circle := abstraction.Circle{Radius: 2}
	assert.Equal(t, 4*math.Pi, circle.Perimeter())
}
