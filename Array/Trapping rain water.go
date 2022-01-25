package main

import "fmt"

func main() {

	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	result := totalTrappedRainWater(arr)
	fmt.Println(result)
}

func totalTrappedRainWater(arr []int) int {
	result := 0
	leftMax, rightMax := make([]int, len(arr)), make([]int, len(arr))
	leftMax[0] = arr[0]
	rightMax[len(rightMax)-1] = arr[len(arr)-1]

	for i := 1; i < len(arr); i++ {
		leftMax[i] = max(leftMax[i-1], arr[i])
	}

	for i := len(arr) - 2; i >= 0; i-- {
		rightMax[i] = max(arr[i], rightMax[i+1])
	}

	for i := 0; i < len(arr); i++ {
		result += min(leftMax[i], rightMax[i]) - arr[i]
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
