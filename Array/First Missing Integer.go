package main

import (
	"fmt"
)

func main() {

	A := []int{3, 4, -1, 1}
	result := firstMissingPositive(A)
	fmt.Println(result)
}

func firstMissingPositive(A []int) int {

	for i := 0; i < len(A); i++ {
		if A[i] >= 1 && A[i] < len(A) {
			if A[i] != i+1 && A[A[i]-1] != A[i] {

				A[A[i]-1], A[i] = A[i], A[A[i]-1]

				if A[i] >= 1 && A[i] < len(A) {
					if A[i] != i+1 && A[A[i]-1] != A[i] {
						i--
					}
				}
			}
		}
	}

	for i := 0; i < len(A); i++ {
		if i >= 0 && i < len(A) {
			if A[i] != i+1 {
				return i + 1
			}
		}
	}

	return len(A) + 1
}
