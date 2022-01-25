package main

import (
	"fmt"
)

func main() {

	arr, totalOranges := []int{1, 3, 2, 4}, 5
	result := orangeIndex(arr, totalOranges)
	fmt.Println(result)
}

func orangeIndex(arr []int, totalOranges int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum >= totalOranges {
			return i
		}
	}

	return -1
}
