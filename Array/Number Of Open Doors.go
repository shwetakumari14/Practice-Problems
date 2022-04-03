package main

import (
	"fmt"
	"math"
)

func main() {

	result := openDoors(10)
	fmt.Println(result)
}

func openDoorsForAllValues(A int) int {
	ans := int(math.Sqrt(float64(A)))

	return ans
}

func openDoors(A int) int {
	count := 0

	factors := make([]int, A+1)
	for i := 1; i <= A; i++ {
		factors[i] = 1
	}

	for i := 1; i <= A; i++ {
		for j := 2 * i; j <= A; j += i {
			factors[j]++
		}
	}

	for i := 1; i < len(factors); i++ {
		if factors[i]%2 == 1 {
			count++
		}
	}

	return count
}
