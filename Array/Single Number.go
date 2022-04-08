package main

import "fmt"

func main() {

	arr := []int{1, 2, 2, 3, 1}
	result := singleNumber(arr)
	fmt.Println(result)
}

func singleNumber(A []int) int {
	ans := 0

	for i := 0; i < len(A); i++ {
		ans ^= A[i]
	}

	return ans
}
