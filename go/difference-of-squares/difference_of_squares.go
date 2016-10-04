package diffsquares

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func SquareOfSums(n int) int {
	return (n * n * (n*n + 2*n + 1)) / 4
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
