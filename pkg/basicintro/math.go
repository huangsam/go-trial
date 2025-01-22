package basicintro

import "errors"

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

// Divide returns the quotient of two integers and an error if division by zero occurs.
func Divide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("division by zero")
	}
	return num1 / num2, nil
}
