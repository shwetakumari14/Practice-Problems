package main

import (
	"fmt"
	"math"
)

func main() {

	arr, b := []int{12, 34, 67, 90}, 2
	result := allocateBooks(arr, b)
	fmt.Println(result)
}

func allocateBooks(arr []int, b int) int {
	n, sum := len(arr), 0
	if n < b {
		return -1
	}

	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	start, end, ans := 0, sum, math.MaxInt

	for i := 0; i < n; i++ {
		mid := start + (end-start)/2
		if isAllocationPossible(arr, b, mid) {
			ans = min(ans, mid)
			end = mid - 1
		} else {
			start = mid + 1
		}

	}

	return ans
}

func isAllocationPossible(arr []int, b, curMin int) bool {
	studentReq, curSum := 1, 0

	for i := 0; i < len(arr); i++ {
		if arr[i] > curMin {
			return false
		}
		if curSum+arr[i] > curMin {
			studentReq++
			curSum = arr[i]
			if studentReq > b {
				return false
			}
		} else {
			curSum += arr[i]
		}
	}

	return true
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
