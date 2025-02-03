package endpoint

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
func RectangleSizeHandler(c *gin.Context) {
	width, err := strconv.ParseFloat(c.DefaultQuery("width", "1.0"), 64)
	if errors.Is(err, strconv.ErrSyntax) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	height, err := strconv.ParseFloat(c.DefaultQuery("height", "1.0"), 64)
	if errors.Is(err, strconv.ErrSyntax) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rect := abstraction.Rectangle{Width: width, Height: height}
	size := abstraction.Classify(&rect)
	payload := RectanglePayload{
		Area:       rect.Area(),
		Perimeter:  rect.Perimeter(),
		Dimensions: rect,
		Size:       size.String(),
	}
	c.JSON(http.StatusOK, payload)
}
