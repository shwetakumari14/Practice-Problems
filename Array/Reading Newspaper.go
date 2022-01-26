package main

import "fmt"

func main() {

	A, B := 100, []int{15, 20, 20, 15, 10, 30, 45}
	result := lastDay(A, B)
	fmt.Println(result)
}

func lastDay(A int, B []int) int {
	sum, i := 0, 0
	for sum < A {
		sum += B[i]
		i++
		i = i % 7

	}

	if i == 0 {
		return 7
	} else {
		return i
	}
}
