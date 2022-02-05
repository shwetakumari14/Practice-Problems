package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int{9, 18, 49, 12, 30}
	result := deleteOneElement(arr)
	fmt.Println(result)
}

func deleteOneElement(A []int) int {
	prefix, suffix, ans := make([]int, len(A)), make([]int, len(A)), math.MinInt
	prefix[0] = A[0]
	suffix[len(suffix)-1] = A[len(A)-1]

	for i := 1; i < len(A); i++ {
		prefix[i] = gcd(prefix[i-1], A[i])
	}

	for i := len(A) - 2; i >= 0; i-- {
		suffix[i] = gcd(suffix[i+1], A[i])
	}

	for i := 0; i < len(A); i++ {
		if i == 0 {
			if suffix[i+1] > ans {
				ans = suffix[i+1]
			}
		} else {
			if i == len(A)-1 {
				if prefix[i-1] > ans {
					ans = prefix[i-1]
				}
			} else if gcd(prefix[i-1], suffix[i+1]) > ans {
				ans = gcd(prefix[i-1], suffix[i+1])
			}
		}
	}

	return ans
}

func gcd(A int, B int) int {
	if A < B {
		A, B = B, A
	}

	if B == 0 {
		return A
	}

	return gcd(B, A%B)

}
