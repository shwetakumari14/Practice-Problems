package main

import "fmt"

func main() {

	arr, b := []int{4, 5, 6, 7, 0, 1, 2, 3}, 2
	result := findElementInRotatedArray(arr, b)
	fmt.Println(result)
}

func findElementInRotatedArray(arr []int, b int) int {
	low, high := 0, len(arr)-1

	for low < high {
		mid := low + (high-low)/2
		if arr[mid] >= arr[low] {
			if arr[mid] < b || b < arr[low] {
				low = mid + 1
			} else {
				high = mid
			}
		} else {
			if arr[mid] >= b || b >= arr[low] {
				low = mid + 1
			} else {
				high = mid
			}
		}
	}

	if arr[low] == b {
		return low
	}

	return -1
}
