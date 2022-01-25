package main

import (
	"fmt"
	"math"
)

//Find the maximum value such that |arr[i]-arr[j]| + |i-j| is maximum in an array

func main() {

	arr := []int{1, 2, 3, 1, 5}
	result := maxDifferenceOptimized(arr)
	fmt.Println(result)
}

//Brute force approach with O(n^2) complexity
func maxDifference(arr []int) int {

	max := 0

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ { //We are companring only i-j pairs and ignoring j-i pairs since value of |i-j| or |j-i| will be same
			if i != j {
				if max < int(math.Abs(float64(i-j)))+int(math.Abs(float64(arr[i]-arr[j]))) {
					max = int(math.Abs(float64(i-j))) + int(math.Abs(float64(arr[i]-arr[j])))
				}
			}
		}
	}

	return max
}

//Optimized approach with O(n) complexity - Let is enforce that i > j
//Case1: arr[i]-arr[j] + i - j --> arr[i] + i - (arr[j] + j)
//Case2: arr[j]-arr[i] + i - j --> arr[j] - j - (arr[i] - i)
func maxDifferenceOptimized(arr []int) int {

	max1, min1, max2, min2 := -100000000, 1000000000, -1000000000, 1000000000

	for i := 0; i < len(arr); i++ {
		if max1 < arr[i]+i {
			max1 = arr[i] + i
		}
		if min1 > arr[i]+i {
			min1 = arr[i] + i
		}

		if max2 < arr[i]-i {
			max2 = arr[i] - i
		}
		if min2 > arr[i]-i {
			min2 = arr[i] - i
		}
	}

	return max((max1 - min1), (max2 - min2))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
