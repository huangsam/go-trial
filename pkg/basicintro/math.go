package basicintro

import "math"

// Add returns the sum of two integers.
func Add(num1 int, num2 int) int {
	return num1 + num2
}

// Subtract returns the difference of two integers.
func Subtract(num1 int, num2 int) int {
	return num1 - num2
}

// Multiply returns the product of two integers.
func Multiply(num1 int, num2 int) int {
	return num1 * num2
}

// Divide returns the quotient of two integers.
func Divide(num1 int, num2 int) int {
	if num2 == 0 {
		// Handle division by zero
		// You can return an error, panic, or use a default value
		// Here, we return Maxint as a placeholder
		return math.MaxInt
	}
	return num1 / num2
}
