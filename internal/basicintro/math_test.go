package basicintro

import "testing"

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
			if got := Add(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
