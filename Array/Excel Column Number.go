package main

import "fmt"

func main() {

	result := titleToNumber("ABCD")
	fmt.Println(result)
}

func titleToNumber(A string) int {
	ans := 0

	for i := 0; i < len(A); i++ {
		ans = ans*26 + int(A[i]) - 64
	}

	return ans
}
