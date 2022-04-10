package main

import (
	"fmt"
	"math"
)

func main() {

	result := divideInt(81, 3)
	fmt.Println(result)
}

func divideInt(A int, B int) int {
	sign, q := 1, 0

	if A < 0 {
		sign = -sign
	}
	if B < 0 {
		sign = -sign
	}

	A = int(math.Abs(float64(A)))
	B = int(math.Abs(float64(B)))

	for i := 31; i >= 0; i-- {
		if (B << i) <= A {
			A -= B << i
			q += 1 << i
		}
	}

	if sign < 0 {
		q = -q
	}

	if q > math.MaxInt {
		q = math.MaxInt
	}

	return q
}
