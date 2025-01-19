package abstraction_test

import (
	"math"
	"testing"

	"io.huangsam/trial/pkg/abstraction"
)

// TestRectangleArea tests the Area method of the Rectangle struct
func TestRectangleArea(t *testing.T) {
	rect := abstraction.Rectangle{Width: 5, Height: 3}
	expectedArea := 15.0
	actualArea := rect.Area()

	if actualArea != expectedArea {
		t.Errorf("Expected area of rectangle to be %.2f, but got %.2f", expectedArea, actualArea)
	}
}

// TestRectanglePerimeter tests the Perimeter method of the Rectangle struct
func TestRectanglePerimeter(t *testing.T) {
	rect := abstraction.Rectangle{Width: 5, Height: 3}
	expectedPerimeter := 16.0
	actualPerimeter := rect.Perimeter()

	if actualPerimeter != expectedPerimeter {
		t.Errorf("Expected perimeter of rectangle to be %.2f, but got %.2f", expectedPerimeter, actualPerimeter)
	}
}

// TestCircleArea tests the Area method of the Circle struct
func TestCircleArea(t *testing.T) {
	circle := abstraction.Circle{Radius: 2}
	expectedArea := math.Pi * 4
	actualArea := circle.Area()

	if actualArea != expectedArea {
		t.Errorf("Expected area of circle to be %.2f, but got %.2f", expectedArea, actualArea)
	}
}

// TestCirclePerimeter tests the Perimeter method of the Circle struct
func TestCirclePerimeter(t *testing.T) {
	circle := abstraction.Circle{Radius: 2}
	expectedPerimeter := 4 * math.Pi
	actualPerimeter := circle.Perimeter()

	if actualPerimeter != expectedPerimeter {
		t.Errorf("Expected perimeter of circle to be %.2f, but got %.2f", expectedPerimeter, actualPerimeter)
	}
}
