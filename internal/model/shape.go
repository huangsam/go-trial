package model

import "github.com/huangsam/go-trial/pkg/abstraction"

// ShapePayload is the JSON response for any handler that deals with shapes.
type ShapePayload struct {
	Area       float64           `json:"area"`
	Perimeter  float64           `json:"perimeter"`
	Dimensions abstraction.Shape `json:"dimensions"`
	Size       string            `json:"size"`
}

// NewShapePayload creates a new ShapePayload from a Shape.
func NewShapePayload(shape abstraction.Shape) ShapePayload {
	return ShapePayload{
		Area:       shape.Area(),
		Perimeter:  shape.Perimeter(),
		Dimensions: shape,
		Size:       abstraction.Classify(shape).String(),
	}
}
