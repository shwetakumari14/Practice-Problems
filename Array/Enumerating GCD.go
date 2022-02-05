package main

import "fmt"

func main() {

	A, B := "2", "9"
	result := solve(A, B)
	fmt.Println(result)
}

func solve(A string, B string) string {
	if A == B {
		return A
	}

	return "1"
}
