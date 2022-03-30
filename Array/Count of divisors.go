package main

import (
	"fmt"
)

func main() {

	arr := []int{8, 9, 10}
	result := countDivisors(arr)
	fmt.Println(result)
}

func countDivisors(A []int) []int {

	var divisors [1000005]int
	var ansArr []int

	for i := 1; i <= 1000000; i++ {
		divisors[i] = i
	}

	for i := 2; i <= 1000000; i++ {
		if divisors[i] == i {
			for j := i * i; j <= 1000000; j += i {
				if divisors[j] == j {
					divisors[j] = i
				}
			}
		}
	}

	for i := 0; i < len(A); i++ {
		temp := A[i]
		ans := 1

		for temp != 1 {
			count := 1
			div := divisors[temp]
			for temp != 1 && temp%div == 0 {
				count++
				temp /= div
			}
			ans *= count
		}

		ansArr = append(ansArr, ans)
	}

	return ansArr
}
