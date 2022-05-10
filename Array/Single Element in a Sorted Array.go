package main

import "fmt"

func main() {

	arr := []int{1, 1, 7}
	result := findSingleElement(arr)
	fmt.Println(result)
}

func findSingleElement(arr []int) int {
	low, high := 0, len(arr)-1

	for low < high-2 {
		mid := low + (high-low)/2
		if arr[mid] == arr[mid+1] {
			mid++
		}
		if (high-mid)%2 == 1 {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if arr[low] == arr[low+1] {
		return arr[high]
	}
	return arr[low]
}
