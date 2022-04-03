package main

import "fmt"

func main() {

	result := convertToTitle(19010)
	fmt.Println(result)
}

func convertToTitle(n int) string {
	ans := ""

	for n > 0 {
		ans = string((n-1)%26+65) + ans
		n = (n - 1) / 26
	}

	return ans
}
