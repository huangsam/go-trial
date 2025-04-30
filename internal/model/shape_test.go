package model_test

import (
	"testing"

	"github.com/huangsam/go-trial/internal/model"
	"github.com/stretchr/testify/assert"
)

type mockShape struct {
	area      float64
	perimeter float64
}

func (m mockShape) Area() float64 {
	return m.area
}

func (m mockShape) Perimeter() float64 {
	return m.perimeter
}

func TestNewShapePayload(t *testing.T) {
	mock := mockShape{
		area:      25.0,
		perimeter: 20.0,
	}

	payload := model.NewShapePayload(mock)

	assert.InEpsilon(t, 25.0, payload.Area, 0.01)
	assert.InEpsilon(t, 20.0, payload.Perimeter, 0.01)
	assert.NotEmpty(t, payload.Size) // Ensure Size is classified
}
