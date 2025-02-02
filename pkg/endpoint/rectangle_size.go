package endpoint

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/huangsam/go-trial/pkg/abstraction"
)

type RectanglePayload struct {
	Area       float64               `json:"area"`
	Perimeter  float64               `json:"perimeter"`
	Dimensions abstraction.Rectangle `json:"dimensions"`
	Size       string                `json:"size"`
}

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
	payload := RectanglePayload{
		Area:       rect.Area(),
		Perimeter:  rect.Perimeter(),
		Dimensions: rect,
		Size:       size.String(),
	}
	return c.JSON(payload)
}
