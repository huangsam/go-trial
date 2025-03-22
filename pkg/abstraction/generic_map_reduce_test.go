package abstraction_test

import (
	"cmp"
	"testing"

	"github.com/huangsam/go-trial/pkg/abstraction"
	"github.com/stretchr/testify/assert"
)

func TestMapReducer_Map(t *testing.T) {
	type args[T cmp.Ordered] struct {
		mapper func(T) T
		input  []T
	}
	type testCase[T cmp.Ordered] struct {
		name string
		mp   abstraction.MapReducer[T]
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "double ints",
			mp:   abstraction.MapReducer[int]{},
			args: args[int]{
				mapper: func(x int) int { return x * 2 },
				input:  []int{1, 2, 3, 4},
			},
			want: []int{2, 4, 6, 8},
		},
		{
			name: "empty input",
			mp:   abstraction.MapReducer[int]{},
			args: args[int]{
				mapper: func(x int) int { return x * 2 },
				input:  []int{},
			},
			want: []int{},
		},
	}
	testsString := []testCase[string]{
		{
			name: "add suffix",
			mp:   abstraction.MapReducer[string]{},
			args: args[string]{
				mapper: func(x string) string { return x + "_suffix" },
				input:  []string{"a", "b", "c"},
			},
			want: []string{"a_suffix", "b_suffix", "c_suffix"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mp.Map(tt.args.mapper, tt.args.input...)
			assert.Equal(t, tt.want, got, "MapReducer.Map() = %v, want %v", got, tt.want)
		})
	}
	for _, tt := range testsString {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mp.Map(tt.args.mapper, tt.args.input...)
			assert.Equal(t, tt.want, got, "MapReducer.Map() = %v, want %v", got, tt.want)
		})
	}
}

func TestMapReducer_Reduce(t *testing.T) {
	type args[T cmp.Ordered] struct {
		reducer func(T, T) T
		input   []T
	}
	type testCase[T cmp.Ordered] struct {
		name string
		mp   abstraction.MapReducer[T]
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "sum ints",
			mp:   abstraction.MapReducer[int]{},
			args: args[int]{
				reducer: func(x, y int) int { return x + y },
				input:   []int{1, 2, 3, 4},
			},
			want: 10,
		},
		{
			name: "empty input",
			mp:   abstraction.MapReducer[int]{},
			args: args[int]{
				reducer: func(x, y int) int { return x + y },
				input:   []int{},
			},
			want: 0,
		},
		{
			name: "one input",
			mp:   abstraction.MapReducer[int]{},
			args: args[int]{
				reducer: func(x, y int) int { return x + y },
				input:   []int{5},
			},
			want: 5,
		},
	}
	testsString := []testCase[string]{
		{
			name: "concat strings",
			mp:   abstraction.MapReducer[string]{},
			args: args[string]{
				reducer: func(x, y string) string { return x + y },
				input:   []string{"a", "b", "c"},
			},
			want: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mp.Reduce(tt.args.reducer, tt.args.input...)
			assert.Equal(t, tt.want, got, "MapReducer.Reduce() = %v, want %v", got, tt.want)
		})
	}
	for _, tt := range testsString {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mp.Reduce(tt.args.reducer, tt.args.input...)
			assert.Equal(t, tt.want, got, "MapReducer.Reduce() = %v, want %v", got, tt.want)
		})
	}
}
