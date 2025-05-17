package abstraction

import "math"

// Shape defines an interface for any geometric shape.
type Shape interface {
	// Area returns the area of the shape.
	Area() float64

	// Perimeter returns the perimeter of the shape.
	Perimeter() float64
}

var _ Shape = &Rectangle{} // Rectangle implements Shape

// Rectangle represents a rectangle with width and height.
type Rectangle struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// NewRectangle creates a new Rectangle instance with the given width and height.
func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{Width: width, Height: height}
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
	Radius float64 `json:"radius"`
}

var _ Shape = &Circle{} // Circle implements Shape

// NewCircle creates a new Circle instance with the given radius.
func NewCircle(radius float64) *Circle {
	return &Circle{Radius: radius}
}

// Area calculates and returns the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates and returns the perimeter (circumference) of the circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// ShapeSize represents the size category of a shape.
type ShapeSize int

const (
	// SizeSmall represents a small shape.
	SizeSmall ShapeSize = iota

	// SizeMedium represents a medium shape.
	SizeMedium

	// SizeLarge represents a large shape.
	SizeLarge
)

// ShapeSize implements the Stringer interface for ShapeSize.
func (s ShapeSize) String() string {
	switch s {
	case SizeSmall:
		return "Small"
	case SizeMedium:
		return "Medium"
	case SizeLarge:
		return "Large"
	default:
		return "Unknown"
	}
}

// Classify determines the size category of a shape based on its area.
//
// This function takes a shape that implements the Shape interface and
// classifies it into one of three size categories: Small, Medium, or Large.
// The classification is based on the following rules:
//
//   - Small: Area <= 10
//   - Medium: 10 < Area <= 100
//   - Large: Area > 100
//
// Assume that shape area is greater than or equal to 0.
func Classify(shape Shape) ShapeSize {
	area := shape.Area()
	switch {
	case area <= 10:
		return SizeSmall
	case area <= 100:
		return SizeMedium
	default:
		return SizeLarge
	}
}
