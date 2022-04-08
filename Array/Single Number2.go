package main

import "fmt"

func main() {

	arr := []int{1, 2, 4, 3, 3, 2, 2, 3, 1, 1}
	result := singleNumberII(arr)
	fmt.Println(result)
}

func singleNumberII(A []int) int {
	n := len(A)
	ones := 0
	twos := 0
	threes := 0
	for i := 0; i < n; i++ {
		twos |= (ones & A[i])
		ones ^= A[i]
		threes = ones & twos
		ones &= ^threes
		twos &= ^threes
	}
	return ones
}
