package main

import "fmt"

func main() {

	A := 2
	result := generateMatrix(A)
	fmt.Println(result)
}

func generateMatrix(A int) [][]int {
	matrix := make([][]int, A)

	for i := range matrix {
		matrix[i] = make([]int, A)
	}

	stRow, endRow, stCol, endCol := 0, len(matrix)-1, 0, len(matrix[0])-1

	k := 1
	for stRow <= endRow && stCol <= endCol {

		for i := stCol; i <= endCol; i++ {
			matrix[stRow][i] = k
			k++
		}
		stRow++

		for i := stRow; i <= endRow; i++ {
			matrix[i][endCol] = k
			k++
		}
		endCol--

		for i := endCol; i >= stCol; i-- {
			matrix[endRow][i] = k
			k++
		}
		endRow--

		for i := endRow; i >= stRow; i-- {
			matrix[i][stCol] = k
			k++
		}
		stCol++
	}

	return matrix
}
