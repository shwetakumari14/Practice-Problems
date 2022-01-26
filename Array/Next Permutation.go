package main

import (
	"fmt"
	"sort"
)

func main() {

	A := []int{1, 5, 8, 4, 7, 3, 5, 6, 1}
	result := nextPermutation(A)
	fmt.Println(result)
}

func nextPermutation(A []int) []int {
	var idx int
	for i := len(A) - 1; i >= 0; i-- {
		if i == 0 {
			sort.Ints(A)
			return A
		}
		if A[i] > A[i-1] {
			idx = i - 1
			break
		}
	}

	for i := len(A) - 1; i >= 0; i-- {
		if A[i] > A[idx] {
			A[i], A[idx] = A[idx], A[i]
			break
		}
	}

	idx++
	n := len(A) - 1
	for idx < n {
		A[idx], A[n] = A[n], A[idx]
		idx++
		n--
	}

	return A

}
