package abstraction_test

import (
	"testing"

	"io.huangsam/trial/pkg/abstraction"
)

func TestClassifyRectangle(t *testing.T) {
	tests := []struct {
		name  string
		shape abstraction.Rectangle
		want  string
	}{
		{
			name:  "Small rectangle",
			shape: abstraction.Rectangle{Width: 2, Height: 2},
			want:  "Small",
		},
		{
			name:  "Medium rectangle",
			shape: abstraction.Rectangle{Width: 5, Height: 5},
			want:  "Medium",
		},
		{
			name:  "Large rectangle",
			shape: abstraction.Rectangle{Width: 11, Height: 11},
			want:  "Large",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abstraction.Classify(tt.shape); got != tt.want {
				t.Errorf("Classify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassifyCircle(t *testing.T) {
	tests := []struct {
		name  string
		shape abstraction.Circle
		want  string
	}{
		{
			name:  "Small circle",
			shape: abstraction.Circle{Radius: 1},
			want:  "Small",
		},
		{
			name:  "Medium circle",
			shape: abstraction.Circle{Radius: 5},
			want:  "Medium",
		},
		{
			name:  "Large circle",
			shape: abstraction.Circle{Radius: 6},
			want:  "Large",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abstraction.Classify(tt.shape); got != tt.want {
				t.Errorf("Classify() = %v, want %v", got, tt.want)
			}
		})
	}
}
