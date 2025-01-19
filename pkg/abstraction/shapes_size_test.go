package abstraction_test

import (
	"testing"

	"io.huangsam/trial/pkg/abstraction"
)

func TestClassifyRectangle(t *testing.T) {
	tests := []struct {
		name  string
		shape abstraction.Rectangle
		want  abstraction.ShapeSize
	}{
		{
			name:  "Small rectangle",
			shape: abstraction.Rectangle{Width: 2, Height: 2},
			want:  abstraction.SizeSmall,
		},
		{
			name:  "Medium rectangle",
			shape: abstraction.Rectangle{Width: 5, Height: 5},
			want:  abstraction.SizeMedium,
		},
		{
			name:  "Large rectangle",
			shape: abstraction.Rectangle{Width: 11, Height: 11},
			want:  abstraction.SizeLarge,
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
		want  abstraction.ShapeSize
	}{
		{
			name:  "Small circle",
			shape: abstraction.Circle{Radius: 1},
			want:  abstraction.SizeSmall,
		},
		{
			name:  "Medium circle",
			shape: abstraction.Circle{Radius: 5},
			want:  abstraction.SizeMedium,
		},
		{
			name:  "Large circle",
			shape: abstraction.Circle{Radius: 6},
			want:  abstraction.SizeLarge,
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

func BenchmarkClassifyRectangle(b *testing.B) {
	rect := abstraction.Rectangle{Width: 5, Height: 5}
	for i := 0; i < b.N; i++ {
		abstraction.Classify(rect)
	}
}

func BenchmarkClassifyCircle(b *testing.B) {
	circle := abstraction.Circle{Radius: 5}
	for i := 0; i < b.N; i++ {
		abstraction.Classify(circle)
	}
}
