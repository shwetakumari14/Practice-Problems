package main

import "fmt"

var mod int = 1e9 + 7

func main() {

	arr := []int{2, 3, 5}
	result := countBits(arr)
	fmt.Println(result)
}

func countBits(arr []int) int {

	ans := 0

	for i := 0; i < 31; i++ {
		count := 0
		for j := 0; j < len(arr); j++ {
			if (arr[j] & (1 << i)) != 0 {
				count++
			}
		}
		ans += (2 * count * (len(arr) - count)) % mod
		ans %= mod
	}

	return ans
}
