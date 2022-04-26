package main

import "fmt"

func main() {

	result := consecutiveSum(9)
	fmt.Println(result)
}

func consecutiveSum(n int) int {
	val, ans, i := 2*n, 0, 1

	for i*i <= val {
		if val%i == 0 {
			temp := val / i
			if (temp-i-1)%2 == 0 && (temp-i-1)/2 >= 0 {
				ans++
			}
		}
		i++
	}

	return ans
}
