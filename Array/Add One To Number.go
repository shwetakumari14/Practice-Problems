package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	A := []int{0, 0, 4, 4, 6, 0, 9, 6, 5, 1}
	result := plusOne1(A)
	fmt.Println(result)
}

func plusOne(A []int) []int {
	tempNum := ""
	for i := 0; i < len(A); i++ {
		tempNum += strconv.Itoa(A[i])
	}

	num, _ := strconv.Atoi(tempNum)
	num += 1

	tempNum = strconv.Itoa(num)
	str := strings.Split(tempNum, "")

	var ans []int

	for _, val := range str {
		intNum, _ := strconv.Atoi(val)
		ans = append(ans, intNum)
	}

	return ans
}

func plusOne1(A []int) []int {
	var temp []int
	var ans []int

	carry := 1

	for i := len(A) - 1; i >= 0; i-- {
		sum := A[i] + carry
		carry = sum / 10
		temp = append(temp, sum%10)
	}
	if carry > 0 {
		temp = append(temp, carry)
	}

	idx := len(temp) - 1
	for i := len(A); i >= 0; i-- {
		if temp[idx] == 0 {
			idx--
		}
	}

	for i := idx; i >= 0; i-- {
		ans = append(ans, temp[i])
	}

	return ans
}
