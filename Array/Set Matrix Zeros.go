package main

import "fmt"

func main() {

	arr := [][]int{{1, 0, 1}, {1, 1, 1}, {1, 0, 1}}
	setZeroes(arr)
	fmt.Println(arr)
}

func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if row[i] == true || col[j] == true {
				matrix[i][j] = 0
			}
		}
	}
}
