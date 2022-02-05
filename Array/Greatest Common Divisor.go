package main

import "fmt"

func main() {

	A, B := 4, 6
	result := gcd(A, B)
	fmt.Println(result)
}

func gcd(A int, B int) int {
	var x, y int
	if A > B {
		x = A
		y = B
	} else {
		x = B
		y = A
	}

	if y == 0 {
		return x
	}

	return gcd(y, x%y)

}
