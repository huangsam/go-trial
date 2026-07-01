package basics_test

import (
	"math"
	"testing"

	"github.com/huangsam/go-trial/lesson/basics"
	"github.com/stretchr/testify/assert"
)

func TestRectangleArea(t *testing.T) {
	rect := basics.NewRectangle(5, 3)
	assert.InEpsilon(t, 15.0, rect.Area(), 0.01)
}

func TestRectanglePerimeter(t *testing.T) {
	rect := basics.NewRectangle(5, 3)
	assert.InEpsilon(t, 16.0, rect.Perimeter(), 0.01)
}

func TestCircleArea(t *testing.T) {
	circle := basics.NewCircle(2)
	assert.InEpsilon(t, 4*math.Pi, circle.Area(), 0.01)
}

func TestCirclePerimeter(t *testing.T) {
	circle := basics.NewCircle(2)
	assert.InEpsilon(t, 4*math.Pi, circle.Perimeter(), 0.01)
}

func TestShapeSize_String(t *testing.T) {
	tests := []struct {
		name string
		size basics.ShapeSize
		want string
	}{
		{
			name: "SmallString",
			size: basics.SizeSmall,
			want: "Small",
		},
		{
			name: "MediumString",
			size: basics.SizeMedium,
			want: "Medium",
		},
		{
			name: "LargeString",
			size: basics.SizeLarge,
			want: "Large",
		},
		{
			name: "UnknownString",
			size: basics.ShapeSize(999), // Invalid ShapeSize value
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
		shape basics.Shape
		want  basics.ShapeSize
	}{
		{
			name:  "SmallRectangle",
			shape: basics.NewRectangle(2, 2),
			want:  basics.SizeSmall,
		},
		{
			name:  "MediumRectangle",
			shape: basics.NewRectangle(5, 5),
			want:  basics.SizeMedium,
		},
		{
			name:  "LargeRectangle",
			shape: basics.NewRectangle(11, 11),
			want:  basics.SizeLarge,
		},
		{
			name:  "SmallCircle",
			shape: basics.NewCircle(1),
			want:  basics.SizeSmall,
		},
		{
			name:  "MediumCircle",
			shape: basics.NewCircle(5),
			want:  basics.SizeMedium,
		},
		{
			name:  "LargeCircle",
			shape: basics.NewCircle(6),
			want:  basics.SizeLarge,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := basics.Classify(tt.shape)
			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkClassify(b *testing.B) {
	b.Run("Rectangle", func(b *testing.B) {
		rect := basics.NewRectangle(5, 5)
		for b.Loop() {
			basics.Classify(rect)
		}
	})
	b.Run("Circle", func(b *testing.B) {
		circle := basics.NewCircle(5)
		for b.Loop() {
			basics.Classify(circle)
		}
	})
}
