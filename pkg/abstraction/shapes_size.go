package abstraction

// Classify determines the size category of a shape based on its area.
func Classify[T Shape](shape T) string {
	area := shape.Area()
	switch {
	case area <= 10:
		return "Small"
	case area <= 100:
		return "Medium"
	default:
		return "Large"
	}
}
