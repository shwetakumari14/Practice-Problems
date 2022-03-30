package main

import (
	"fmt"
)

func main() {

	result := oddFibpnacciCount(43, 45)
	fmt.Println(result)
}

func oddFibpnacciCount(a, b int) int {
	numbers := b - a + 1

	multipleOf3 := (b / 3) - ((a - 1) / 3)

	return numbers - multipleOf3
}

func oddFibpnacciCount1(a, b int) int {
	count := 0
	var ans []int

	for i := a; i <= b; i++ {
		ans = append(ans, fibonacci(i))
	}

	for i := 0; i < len(ans); i++ {
		if ans[i]%2 == 1 {
			count++
		}
	}

	return count
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
