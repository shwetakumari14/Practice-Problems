package main

import "fmt"

var mod int = 1000003

func main() {

	result := findRank("dca")
	fmt.Println(result)
}

func findRank(A string) int {
	ans, str := 0, []rune(A)

	for i := 0; i < len(str)-1; i++ {
		count := 0
		for j := i + 1; j < len(str); j++ {
			if str[j] < str[i] {
				count++
			}
		}
		ans += (count * fact(len(str)-i-1)) % mod
	}

	return (ans + 1) % mod
}

func fact(n int) int {
	if n <= 1 {
		return 1
	}

	return (n * fact(n-1)) % mod
}
