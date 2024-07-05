package main

import "fmt"

func main() {
	PascalTriangle()
}

func PascalTriangle() {
	var N int
	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&N)

	fmt.Printf("Pascal Row at %d: %v\n", N, nthRow(N))
}

func nthRow(N int) []int {
	res := make([]int, N+1)
	res[0] = 1

	for i := 1; i <= N; i++ {
		res[i] = res[i-1] * (N - i + 1) / i
	}

	return res
}
