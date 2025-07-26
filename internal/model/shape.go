package model

import "github.com/huangsam/go-trial/lesson/abstraction"

// ShapePayload is the JSON response for any handler that deals with shapes.
//
// The dimensions field is an interface to allow for different shapes.
// The size field indicates the shape classification (e.g., small, medium, large).
type ShapePayload struct {
	Area       float64           `json:"area"`
	Perimeter  float64           `json:"perimeter"`
	Dimensions abstraction.Shape `json:"dimensions"`
	Size       string            `json:"size"`
}

// NewShapePayload creates a new ShapePayload from a Shape.
//
// The endpoint showcases the power of interfaces in Go.
// It allows us to handle different shapes (e.g., Rectangle, Circle) uniformly.
func NewShapePayload(shape abstraction.Shape) ShapePayload {
	return ShapePayload{
		Area:       shape.Area(),
		Perimeter:  shape.Perimeter(),
		Dimensions: shape,
		Size:       abstraction.Classify(shape).String(),
	}
}
