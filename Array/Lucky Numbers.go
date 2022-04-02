package main

import "fmt"

func main() {

	result := luckyNumber2(12)
	fmt.Println(result)
}

func luckyNumber(A int) int {
	count := 0

	for i := 2; i <= A; i++ {
		if isLucky(i) {
			count++
		}
	}

	return count
}

func isLucky(A int) bool {
	count := 0
	for i := 2; i <= A; i++ {
		if isPrime(i) {
			if A%i == 0 {
				count++
			}
		}
	}

	if count == 2 {
		return true
	}

	return false
}

func isPrime(A int) bool {
	for i := 2; i*i <= A; i++ {
		if A%i == 0 {
			return false
		}
	}
	return true
}

func luckyNumber2(A int) int {
	count := 0
	factors := make([]int, A+1)

	for i := 0; i <= A; i++ {
		factors[i] = 0
	}

	for i := 2; i <= A; i++ {
		if factors[i] == 0 {
			for j := 2 * i; j <= A; j += i {
				factors[j] += 1
			}
		}
	}

	for i := 6; i <= A; i++ {
		if factors[i] == 2 {
			count++
		}
	}

	return count
}
