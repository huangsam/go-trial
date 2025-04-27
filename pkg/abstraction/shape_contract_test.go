package abstraction_test

import (
	"math"
	"testing"

	"github.com/huangsam/go-trial/pkg/abstraction"
	"github.com/stretchr/testify/assert"
)

func TestRectangleArea(t *testing.T) {
	rect := abstraction.NewRectangle(5, 3)
	assert.Equal(t, 15.0, rect.Area())
}

func TestRectanglePerimeter(t *testing.T) {
	rect := abstraction.NewRectangle(5, 3)
	assert.Equal(t, 16.0, rect.Perimeter())
}

func TestCircleArea(t *testing.T) {
	circle := abstraction.NewCircle(2)
	assert.Equal(t, 4*math.Pi, circle.Area())
}

func TestCirclePerimeter(t *testing.T) {
	circle := abstraction.NewCircle(2)
	assert.Equal(t, 4*math.Pi, circle.Perimeter())
}

func TestShapeSize_String(t *testing.T) {
	tests := []struct {
		name string
		size abstraction.ShapeSize
		want string
	}{
		{
			name: "SmallString",
			size: abstraction.SizeSmall,
			want: "Small",
		},
		{
			name: "MediumString",
			size: abstraction.SizeMedium,
			want: "Medium",
		},
		{
			name: "LargeString",
			size: abstraction.SizeLarge,
			want: "Large",
		},
		{
			name: "UnknownString",
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

func TestClassify(t *testing.T) {
	tests := []struct {
		name  string
		shape abstraction.Shape
		want  abstraction.ShapeSize
	}{
		{
			name:  "SmallRectangle",
			shape: abstraction.NewRectangle(2, 2),
			want:  abstraction.SizeSmall,
		},
		{
			name:  "MediumRectangle",
			shape: abstraction.NewRectangle(5, 5),
			want:  abstraction.SizeMedium,
		},
		{
			name:  "LargeRectangle",
			shape: abstraction.NewRectangle(11, 11),
			want:  abstraction.SizeLarge,
		},
		{
			name:  "SmallCircle",
			shape: abstraction.NewCircle(1),
			want:  abstraction.SizeSmall,
		},
		{
			name:  "MediumCircle",
			shape: abstraction.NewCircle(5),
			want:  abstraction.SizeMedium,
		},
		{
			name:  "LargeCircle",
			shape: abstraction.NewCircle(6),
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

func BenchmarkClassify(b *testing.B) {
	b.Run("Rectangle", func(b *testing.B) {
		rect := abstraction.NewRectangle(5, 5)
		for b.Loop() {
			abstraction.Classify(rect)
		}
	})
	b.Run("Circle", func(b *testing.B) {
		circle := abstraction.NewCircle(5)
		for b.Loop() {
			abstraction.Classify(circle)
		}
	})
}
