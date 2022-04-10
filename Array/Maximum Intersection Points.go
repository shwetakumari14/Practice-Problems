package main

import "fmt"

func main() {

	result := getIntersection(2, 2)
	fmt.Println(result)
}

func getIntersection(A int, B int) int {

	x := A * (A - 1) / 2
	y := B*(B-1) + 2*A*B

	return x + y
}
