package main

import "fmt"

func main() {

	result := magicNum(10)
	fmt.Println(result)
}

func magicNum(A int) int {
	ans, x := 0, 1

	for A > 0 {
		x *= 5
		if A%2 == 1 {
			ans += x
		}
		A /= 2
	}

	return ans
}
