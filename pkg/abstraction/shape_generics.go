package abstraction

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
func Classify[T Shape](shape T) ShapeSize {
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
