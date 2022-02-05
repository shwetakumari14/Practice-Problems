package main

import "fmt"

func main() {

	arr := []int{7, 2, 5}
	result := makeGCDOne(arr)
	fmt.Println(result)
}

func makeGCDOne(A []int) int {
	gcd := A[0]

	for i := 1; i < len(A); i++ {
		gcd = findGcd(gcd, A[i])
	}

	if gcd == 1 {
		return 0
	}

	return gcd
}

func findGcd(A int, B int) int {
	if A < B {
		A, B = B, A
	}

	if B == 0 {
		return A
	}

	return findGcd(B, A%B)

}
