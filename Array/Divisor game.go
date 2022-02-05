package main

import "fmt"

func main() {

	a, b, c := 12, 3, 2
	result := solve(a, b, c)
	fmt.Println(result)
}

func solve(A int, B int, C int) int {
	lcm, count := (B*C)/gcd(B, C), 0

	i := lcm
	for i <= A {
		if i%lcm == 0 {
			count++
		}
		i += lcm
	}

	return count

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
