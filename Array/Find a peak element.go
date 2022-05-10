package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}
	result := findPeakElement(arr)
	fmt.Println(result)
}

func findPeakElement(A []int) int {
	if len(A) == 1 {
		return A[0]
	}
	if A[0] >= A[1] {
		return A[0]
	}
	if A[len(A)-1] >= A[len(A)-2] {
		return A[len(A)-1]
	}
	low, high, mid := 0, len(A)-2, -1

	for low <= high {
		mid = low + (high-low)/2
		if A[mid] >= A[mid-1] && A[mid] >= A[mid+1] {
			return A[mid]
		} else if A[mid] > A[mid-1] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
