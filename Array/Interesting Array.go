package main

import "fmt"

func main() {

	arr := []int{9, 17}
	result := checkInterestingArray(arr)
	fmt.Println(result)
}

func checkInterestingArray(arr []int) string {
	count := 0

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			count++
		}
	}

	if count%2 == 0 {
		return "Yes"
	}

	return "No"
}
