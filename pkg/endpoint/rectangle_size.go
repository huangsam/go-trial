package endpoint

import (
	"net/http"
	"strconv"

	"github.com/huangsam/go-trial/pkg/abstraction"
	"github.com/labstack/echo/v4"
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
func RectangleSizeHandler(c echo.Context) error {
	width, err := strconv.ParseFloat(c.QueryParam("width"), 64)
	if err != nil {
		width = 1.0
	}
	height, err := strconv.ParseFloat(c.QueryParam("height"), 64)
	if err != nil {
		height = 1.0
	}
	rect := abstraction.Rectangle{Width: width, Height: height}
	size := abstraction.Classify(&rect)
	payload := RectanglePayload{
		Area:       rect.Area(),
		Perimeter:  rect.Perimeter(),
		Dimensions: rect,
		Size:       size.String(),
	}
	return c.JSON(http.StatusOK, payload)
}
