package basicintro_test

import (
	"math"
	"testing"

	"io.huangsam/trial/pkg/basicintro"
)

func TestAdd(t *testing.T) {
	type args struct {
		num1 int32
		num2 int32
	}

	tests := []struct {
		name string
		args args
		want int32
	}{
		{name: "Positive Numbers", args: args{num1: 1, num2: 2}, want: 3},
		{name: "Negative Numbers", args: args{num1: -1, num2: -2}, want: -3},
		{name: "Zero and Positive", args: args{num1: 0, num2: 5}, want: 5},
		{name: "Zero and Negative", args: args{num1: 0, num2: -5}, want: -5},
		{name: "Large Numbers", args: args{num1: 100000, num2: 200000}, want: 300000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basicintro.Add(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	type args struct {
		num1 int32
		num2 int32
	}

	tests := []struct {
		name string
		args args
		want int32
	}{
		{name: "Positive Numbers", args: args{num1: 5, num2: 2}, want: 3},
		{name: "Negative Numbers", args: args{num1: -1, num2: -2}, want: 1},
		{name: "Positive - Negative", args: args{num1: 5, num2: -2}, want: 7},
		{name: "Negative - Positive", args: args{num1: -1, num2: 2}, want: -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basicintro.Subtract(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	type args struct {
		num1 int32
		num2 int32
	}

	tests := []struct {
		name string
		args args
		want int32
	}{
		{name: "Positive Numbers", args: args{num1: 2, num2: 3}, want: 6},
		{name: "Negative Numbers", args: args{num1: -2, num2: -3}, want: 6},
		{name: "Positive and Negative", args: args{num1: 2, num2: -3}, want: -6},
		{name: "Zero", args: args{num1: 0, num2: 5}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basicintro.Multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		num1 int32
		num2 int32
	}

	tests := []struct {
		name string
		args args
		want int32
	}{
		{name: "Positive Numbers", args: args{num1: 10, num2: 2}, want: 5},
		{name: "Negative Numbers", args: args{num1: -10, num2: -2}, want: 5},
		{name: "Positive and Negative", args: args{num1: 10, num2: -2}, want: -5},
		{name: "Zero Dividend", args: args{num1: 0, num2: 5}, want: 0},
		{name: "Division by Zero", args: args{num1: 5, num2: 0}, want: math.MaxInt32}, // Test division by zero
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basicintro.Divide(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		basicintro.Add(1, 2)
	}
}

func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		basicintro.Subtract(10, 5)
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		basicintro.Multiply(10, 5)
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		basicintro.Divide(10, 2) // Avoid division by zero in benchmark
	}
}
