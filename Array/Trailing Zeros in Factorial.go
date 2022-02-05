package main

import "fmt"

func main() {

	A := 4617
	result := trailingZeroes(A)
	fmt.Println(result)
}

func trailingZeroes(A int) int {
	ans, i, div := 0, 5, A/5

	for div > 0 {
		ans += div
		i *= 5
		div = A / i
	}

	return ans
}
