package abstraction_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/abstraction"
	"github.com/stretchr/testify/assert"
)

func TestShapeSize_String(t *testing.T) {
	tests := []struct {
		name string
		size abstraction.ShapeSize
		want string
	}{
		{
			name: "Small string",
			size: abstraction.SizeSmall,
			want: "Small",
		},
		{
			name: "Medium string",
			size: abstraction.SizeMedium,
			want: "Medium",
		},
		{
			name: "Large string",
			size: abstraction.SizeLarge,
			want: "Large",
		},
		{
			name: "Unknown string",
			size: abstraction.ShapeSize(999), // Invalid ShapeSize value
			want: "Unknown",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.size.String())
		})
	}
}

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
			got := abstraction.Classify(tt.shape)
			assert.Equal(t, tt.want, got)
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
			got := abstraction.Classify(tt.shape)
			assert.Equal(t, tt.want, got)
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
