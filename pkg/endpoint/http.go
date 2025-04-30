package endpoint

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/huangsam/go-trial/pkg/abstraction"
)

// HelloHandler returns a simple greeting as text.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello world"))
}

// ErrorHandler returns a generic error as text.
func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("Generic error"))
}

// ShapePayload is the JSON response for any handler that deals with shapes.
type ShapePayload struct {
	Area       float64           `json:"area"`
	Perimeter  float64           `json:"perimeter"`
	Dimensions abstraction.Shape `json:"dimensions"`
	Size       string            `json:"size"`
}

// RectangleSizeHandler returns rectangle info as JSON.
//
// Accepts 'width' and 'height' query parameters (defaulting to 1.0).
// Returns an error if width or height are not valid numbers.
func RectangleSizeHandler(w http.ResponseWriter, r *http.Request) {
	width, err := strconv.ParseFloat(r.URL.Query().Get("width"), 64)
	if err != nil {
		width = 1.0
	}
	height, err := strconv.ParseFloat(r.URL.Query().Get("height"), 64)
	if err != nil {
		height = 1.0
	}
	rect := abstraction.NewRectangle(width, height)
	size := abstraction.Classify(rect)
	payload := ShapePayload{
		Area:       rect.Area(),
		Perimeter:  rect.Perimeter(),
		Dimensions: rect,
		Size:       size.String(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(payload)
}
