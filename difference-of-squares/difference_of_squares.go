package diffsquares

// SquareOfSum takes integer parameter and sums from 1->n, then returns the
// square
func SquareOfSum(n int) int {
	result := n * (n + 1) / 2

	return result * result
}

// SumOfSquares takes interger parameter and sums the square of 1->n
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference does the SumOfSquares - SquareOfSum for n
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
