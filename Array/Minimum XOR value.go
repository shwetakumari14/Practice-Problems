package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	arr := []int{0, 4, 7, 9}
	result := findMinXor(arr)
	fmt.Println(result)
}

func findMinXor(A []int) int {
	ans := math.MaxInt

	sort.Ints(A)

	for i := 0; i < len(A)-1; i++ {
		min := A[i] ^ A[i+1]
		if ans > min {
			ans = min
		}
	}

	return ans
}

func findMinXor2(A []int) int {

	ans := math.MaxInt

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if i != j {
				min := A[i] ^ A[j]
				if ans > min {
					ans = min
				}
			}
		}
	}

	return ans
}
