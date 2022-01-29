package main

import (
	"fmt"
	"strings"
)

func main() {

	A := "ABCGAG"
	result := subStringCount(A)
	fmt.Println(result)

}

func subStringCount(A string) int {
	aCount, ans, str := 0, 0, strings.Split(A, "")

	for i := 0; i < len(str); i++ {
		if str[i] == "A" {
			aCount++
		}
		if str[i] == "G" {
			ans += aCount
		}
	}

	return ans
}
