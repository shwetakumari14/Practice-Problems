package main

import (
	"fmt"
	"sort"
)

func main() {

	arr, b := []int{1, 2, 3, 4, 5}, 3
	result := aggressiveCows(arr, b)
	fmt.Println(result)
}

func aggressiveCows(arr []int, b int) int {
	sort.Ints(arr)
	left, right, ans := 1, arr[len(arr)-1], 1

	for i := 0; i < len(arr); i++ {
		mid := left + (right-left)/2
		if isPossible(arr, mid, b) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func isPossible(arr []int, x, c int) bool {
	j, count := 0, 1
	for i := 0; i < len(arr); i++ {
		if arr[i]-arr[j] >= x {
			j = i
			count++
		}
	}

	if count >= c {
		return true
	}
	return false
}
