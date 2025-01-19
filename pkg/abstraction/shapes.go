package abstraction

import (
	"math"
)

// Shape defines an interface for any geometric shape.
type Shape interface {
	// Area returns the area of the shape.
	Area() float64
	// Perimeter returns the perimeter of the shape.
	Perimeter() float64
}

// Rectangle represents a rectangle with width and height.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates and returns the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates and returns the perimeter of the rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle represents a circle with a given radius.
type Circle struct {
	Radius float64
}

// Area calculates and returns the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates and returns the perimeter (circumference) of the circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
