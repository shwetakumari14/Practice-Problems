package main

import "fmt"

func main() {

	A, B := []int{1, 12, 10, 3, 14, 10, 5}, 8
	result := minSwaps(A, B)
	fmt.Println(result)
}

func minSwaps(A []int, B int) int {
	minCount, ans := 0, 0

	for i := 0; i < len(A); i++ {
		if A[i] <= B {
			minCount++
		}
	}
	if minCount <= 1 {
		return 0
	} else {
		leftCount, rightCount, count := 0, 0, 0
		for rightCount < minCount {
			if A[rightCount] > B {
				count++
			}
			rightCount++
		}
		ans = count
		for rightCount < len(A) {
			if A[rightCount] > B {
				count++
			}
			if A[leftCount] > B {
				count--
			}
			if count < ans {
				ans = count
			}
			rightCount++
			leftCount++
		}
	}

	return ans
}
