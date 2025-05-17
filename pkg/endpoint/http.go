package endpoint

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/huangsam/go-trial/internal/model"
	"github.com/huangsam/go-trial/pkg/abstraction"
)

// HelloHandler returns a simple greeting as text.
func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello world"))
}

// ErrorHandler returns a generic error as text.
func ErrorHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("Generic error"))
}

// RectangleSizeHandler returns rectangle info as JSON.
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(model.NewShapePayload(rect))
}

// CircleSizeHandler returns circle info as JSON.
func CircleSizeHandler(w http.ResponseWriter, r *http.Request) {
	radius, err := strconv.ParseFloat(r.FormValue("radius"), 64)
	if err != nil {
		radius = 1.0
	}
	circle := abstraction.NewCircle(radius)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(model.NewShapePayload(circle))
}
