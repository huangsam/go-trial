package basicintro_test

import (
	"testing"

	"github.com/huangsam/go-trial/pkg/basicintro"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		{name: "PositiveNumbers", args: args{num1: 1, num2: 2}, want: 3},
		{name: "NegativeNumbers", args: args{num1: -1, num2: -2}, want: -3},
		{name: "ZeroAndPositive", args: args{num1: 0, num2: 5}, want: 5},
		{name: "ZeroAndNegative", args: args{num1: 0, num2: -5}, want: -5},
		{name: "LargeNumbers", args: args{num1: 100000, num2: 200000}, want: 300000},
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
		{name: "PositiveNumbers", args: args{num1: 5, num2: 2}, want: 3},
		{name: "NegativeNumbers", args: args{num1: -1, num2: -2}, want: 1},
		{name: "PositiveAndNegative", args: args{num1: 5, num2: -2}, want: 7},
		{name: "NegativeAndPositive", args: args{num1: -1, num2: 2}, want: -3},
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
		{name: "PositiveNumbers", args: args{num1: 2, num2: 3}, want: 6},
		{name: "NegativeNumbers", args: args{num1: -2, num2: -3}, want: 6},
		{name: "PositiveAndNegative", args: args{num1: 2, num2: -3}, want: -6},
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
		name    string
		args    args
		want    int
		wantErr error
	}{
		{name: "PositiveNumbers", args: args{num1: 10, num2: 2}, want: 5, wantErr: nil},
		{name: "NegativeNumbers", args: args{num1: -10, num2: -2}, want: 5, wantErr: nil},
		{name: "PositiveAndNegative", args: args{num1: 10, num2: -2}, want: -5, wantErr: nil},
		{name: "ZeroDividend", args: args{num1: 0, num2: 5}, want: 0, wantErr: nil},
		{name: "DivisionByZero", args: args{num1: 5, num2: 0}, want: 0, wantErr: basicintro.ErrDivision},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := basicintro.Divide(tt.args.num1, tt.args.num2)
			if tt.wantErr != nil {
				require.EqualError(t, err, tt.wantErr.Error())
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
