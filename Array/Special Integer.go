package main

import "fmt"

func main() {

	arr, b := []int{1, 2, 3, 4, 5}, 10
	result := maxValueOfSubArray(arr, b)
	fmt.Println(result)
}

func maxValueOfSubArray(arr []int, b int) int {
	ans, sum, start := len(arr), 0, 0

	for end := 0; end < len(arr); end++ {
		sum += arr[end]
		for sum > b {
			sum -= arr[start]
			start++

			ans = min(ans, end-start+1)

			if sum == 0 {
				break
			}
		}
		if sum == 0 {
			ans = -1
			break
		}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
