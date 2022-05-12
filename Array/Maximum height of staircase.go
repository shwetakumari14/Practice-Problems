package main

import "fmt"

func main() {

	result := maxHeighrOfStairCase(10)
	fmt.Println(result)
}

func maxHeighrOfStairCase(n int) int {
	low, high, ans := 0, 1000000000, 0

	for low <= high {
		mid := low + (high-low)/2
		val := mid * (mid + 1) / 2
		if val == n {
			return mid
		}
		if val < n {
			if ans < mid {
				ans = mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}
