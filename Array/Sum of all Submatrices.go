package main

import "fmt"

func main() {

	A := [][]int{{1, 1}, {1, 1}}
	result := sumOfSubMatrix(A)
	fmt.Println(result)
}

func sumOfSubMatrix(A [][]int) int {
	sum := 0
	n, m := len(A), len(A[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sum += (i + 1) * (n - i) * (j + 1) * (m - j) * A[i][j]
		}
	}

	return sum
}
