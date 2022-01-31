package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	A := "010"
	result := flip(A)
	fmt.Println(result)
}

func flip(A string) []int {
	str := strings.Split(A, "")

	tempArray, ansArray := make([]int, len(str)), make([]int, 2)
	ansArray[0] = math.MaxInt
	ansArray[1] = math.MaxInt

	for i := 0; i < len(str); i++ {
		if str[i] == "0" {
			tempArray[i] = 1
		} else {
			tempArray[i] = -1
		}
	}

	p, max_ending_here, max_ending_till_now := 0, 0, 0

	for i := 0; i < len(tempArray); i++ {
		if max_ending_here+tempArray[i] < 0 {
			p = i + 1
			max_ending_here = 0
		} else {
			max_ending_here += tempArray[i]
		}

		if max_ending_here > max_ending_till_now {
			max_ending_till_now = max_ending_here
			ansArray[0] = p
			ansArray[1] = i
		}
	}

	if ansArray[0] == math.MaxInt {
		return []int{}
	} else {
		ansArray[0] += 1
		ansArray[1] += 1
	}

	return ansArray
}
