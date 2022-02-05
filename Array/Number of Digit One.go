package main

import "fmt"

func main() {

	A := 100
	result := countOne(A)
	fmt.Println(result)
}

func countOne(A int) int {
	ans := 0

	for i := 1; i <= A; i *= 10 {
		ptr := i * 10
		ans += (A/ptr)*i + min(max(A%ptr-i+1, 0), i)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
