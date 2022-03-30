package main

import "fmt"

var mod int = 1e9 + 7

func main() {

	result := rangeSum(0, 3)
	fmt.Println(result)
}

func rangeSum(A, B int) int {
	fA := fibonacciSum(A + 1)
	fB := fibonacciSum(B + 2)

	res := fB - fA

	if res < 0 {
		res += mod
	}

	return res % mod
}

func fibonacciSum(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	var matrix [2][2]int
	matrix[0][0] = 1
	matrix[0][1] = 1
	matrix[1][0] = 1

	var matPower [2][2]int = power(matrix, n-1)

	return matPower[0][0]
}

func power(mat [2][2]int, n int) [2][2]int {
	if n == 0 {
		var matrix [2][2]int
		matrix[0][0] = 1
		matrix[0][1] = 0
		matrix[1][0] = 0
		matrix[1][0] = 1
	}

	if n == 1 {
		return mat
	}

	var half [2][2]int = power(mat, n/2)
	half = matrixMultiple(half, half)
	if n%2 == 1 {
		half = matrixMultiple(mat, half)
	}

	return half
}

func matrixMultiple(m1, m2 [2][2]int) [2][2]int {
	var res [2][2]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				res[i][j] += (m1[i][k] * m2[k][j]) % mod
				res[i][j] = res[i][j] % mod
			}
		}
	}

	return res
}
