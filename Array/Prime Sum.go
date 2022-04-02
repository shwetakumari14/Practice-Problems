package main

import "fmt"

func main() {

	result := primesum(184)
	fmt.Println(result)
}

func primesum(A int) []int {
	ans := make([]int, 2)
	prime := make([]bool, A+1)
	for i := 0; i <= A; i++ {
		prime[i] = true
	}

	prime[0] = false
	prime[1] = false

	for i := 2; i <= A; i++ {

		if !prime[i] {
			continue
		}

		if i > A/i {
			break
		}

		for j := i * i; j <= A; j += i {
			prime[j] = false
		}
	}

	for i := 2; i <= A; i++ {
		if prime[i] && prime[A-i] {
			ans[0] = i
			ans[1] = A - i
			return ans
		}
	}

	return ans

}
