package main

import (
	"fmt"
	"math"
)

func main() {

	result := getXOR(5)
	fmt.Println(result)
}

func getXOR(A int) int {

	noOfBits := int(math.Log2(float64(A))) + 1
	X, Y := 0, 0

	for i := 0; i < noOfBits; i++ {
		if (A & (1 << i)) != 0 {
			continue
		}
		X += 1 << i
	}
	Y = 1 << noOfBits

	return X ^ Y
}
