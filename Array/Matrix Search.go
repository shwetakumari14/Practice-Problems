package main

import "fmt"

func main() {

	arr, b := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 3
	result := findElementInSortedMatrix(arr, b)
	fmt.Println(result)
}

func findElementInSortedMatrix(arr [][]int, b int) int {
	i, j := 0, len(arr[0])-1

	for i <= len(arr)-1 && j >= 0 {
		if arr[i][j] == b {
			return 1
		}

		if arr[i][j] > b {
			j--
		} else {
			i++
		}
	}

	return 0
}
