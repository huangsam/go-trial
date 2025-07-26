package abstraction

import "cmp"

// MapReducer provides generic Map and Reduce operations for ordered types.
type MapReducer[T cmp.Ordered] struct{}

// Map applies a function to each input element, resulting in a new slice.
func (mp MapReducer[T]) Map(mapper func(T) T, input ...T) []T {
	result := make([]T, len(input))
	for i := range input {
		result[i] = mapper(input[i])
	}
	return result
}

// Reduce applies a function cumulatively to the input elements, resulting in a single value.
func (mp MapReducer[T]) Reduce(reducer func(T, T) T, input ...T) T {
	var result T
	if len(input) == 0 {
		return result
	}
	result = input[0]     // first input
	leftover := input[1:] // remaining inputs
	for i := range leftover {
		result = reducer(result, leftover[i])
	}
	return result
}
