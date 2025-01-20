package basicintro_test

import (
	"math"
	"testing"

	"github.com/huangsam/go-trial/pkg/basicintro"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Positive Numbers", args: args{num1: 1, num2: 2}, want: 3},
		{name: "Negative Numbers", args: args{num1: -1, num2: -2}, want: -3},
		{name: "Zero and Positive", args: args{num1: 0, num2: 5}, want: 5},
		{name: "Zero and Negative", args: args{num1: 0, num2: -5}, want: -5},
		{name: "Large Numbers", args: args{num1: 100000, num2: 200000}, want: 300000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, basicintro.Add(tt.args.num1, tt.args.num2))
		})
	}
}

func TestSubtract(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Positive Numbers", args: args{num1: 5, num2: 2}, want: 3},
		{name: "Negative Numbers", args: args{num1: -1, num2: -2}, want: 1},
		{name: "Positive - Negative", args: args{num1: 5, num2: -2}, want: 7},
		{name: "Negative - Positive", args: args{num1: -1, num2: 2}, want: -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, basicintro.Subtract(tt.args.num1, tt.args.num2))
		})
	}
}

func TestMultiply(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Positive Numbers", args: args{num1: 2, num2: 3}, want: 6},
		{name: "Negative Numbers", args: args{num1: -2, num2: -3}, want: 6},
		{name: "Positive and Negative", args: args{num1: 2, num2: -3}, want: -6},
		{name: "Zero", args: args{num1: 0, num2: 5}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, basicintro.Multiply(tt.args.num1, tt.args.num2))
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Positive Numbers", args: args{num1: 10, num2: 2}, want: 5},
		{name: "Negative Numbers", args: args{num1: -10, num2: -2}, want: 5},
		{name: "Positive and Negative", args: args{num1: 10, num2: -2}, want: -5},
		{name: "Zero Dividend", args: args{num1: 0, num2: 5}, want: 0},
		{name: "Division by Zero", args: args{num1: 5, num2: 0}, want: math.MaxInt}, // Test division by zero
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, basicintro.Divide(tt.args.num1, tt.args.num2))
		})
	}
}
