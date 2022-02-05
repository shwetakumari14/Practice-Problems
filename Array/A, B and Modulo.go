package main

import "fmt"

func main() {

	a, b := 5, 10
	result := findModulo(a, b)
	fmt.Println(result)
}

func findModulo(A int, B int) int {
	if A < B {
		A, B = B, A
	}

	return A - B
}
