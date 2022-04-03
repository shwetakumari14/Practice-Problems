package main

import "fmt"

func main() {

	result := rectArea(0, 0, 4, 4, 2, 2, 6, 6)
	fmt.Println(result)
}

func rectArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	x := min(C, G) - max(A, E)
	y := min(D, H) - max(B, F)
	area := x * y
	if x < 0 || y < 0 {
		area = 0
	}
	return area
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
