package basicintro

// Fibonacci calculates the nth Fibonacci number.
//
// The Fibonacci sequence is a series of numbers where each number is the sum
// of the two preceding ones, starting from 0 and 1.
func Fibonacci(n int) int {
	if n <= 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	num1 := 2
	num2 := 1
	for i := 2; i < n; i++ {
		tmp := num1 + num2
		num2 = num1
		num1 = tmp
	}
	return num1
}