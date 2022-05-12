package main

import "fmt"

func main() {

	arr := [][]int{{1, 3, 5}, {2, 6, 9}, {3, 6, 9}}
	result := matrixMedian(arr)
	fmt.Println(result)

}

func matrixMedian(arr [][]int) int {
	low, high := 0, 1000000000
	medianPos := len(arr) * len(arr[0]) / 2
	ans := -1

	for low <= high {
		mid := low + (high-low)/2
		count := 0

		for i := 0; i < len(arr); i++ {
			count += countEleLessThanMedian(arr[i], mid)
		}

		if count > medianPos {
			high = mid - 1
		} else {
			ans = mid
			low = mid + 1
		}
	}

	return ans
}

func countEleLessThanMedian(arr []int, val int) int {
	low, high := 0, len(arr)-1
	ans := -1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] < val {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans + 1
}
