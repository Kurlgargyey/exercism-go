// Package diffsquares provides functionality for finding the difference between the square of the sum and the sum of the squares of the first N natural numbers.
package diffsquares

// SquareOfSum calculates the square of the sum of the first n natural numbers.
func SquareOfSum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares calculates the sum of the squares of the first n natural numbers.
func SumOfSquares(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// Difference calculates the difference between the square of the sum and the sum of the squares of the first n natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
