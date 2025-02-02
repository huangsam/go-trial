package endpoint

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/huangsam/go-trial/pkg/abstraction"
)

// RectangleSizeHandler returns rectangle info as JSON.
//
// Accepts 'width' and 'height' query parameters (defaulting to 1.0).
// Returns an error if width or height are not valid numbers.
func RectangleSizeHandler(c *fiber.Ctx) error {
	width, err := strconv.ParseFloat(c.Query("width", "1.0"), 64)
	if errors.Is(err, strconv.ErrSyntax) {
		return c.JSON(map[string]error{"error": err})
	}
	height, err := strconv.ParseFloat(c.Query("height", "1.0"), 64)
	if errors.Is(err, strconv.ErrSyntax) {
		return c.JSON(map[string]error{"error": err})
	}
	rect := abstraction.Rectangle{Width: width, Height: height}
	size := abstraction.Classify(&rect)
	payload := map[string]any{
		"area":      rect.Area(),
		"perimeter": rect.Perimeter(),
		"shape":     rect,
		"size":      size.String(),
	}
	return c.JSON(payload)
}
