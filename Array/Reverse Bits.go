package main

import (
	"fmt"
)

func main() {

	var num uint32 = 2
	result := reverseBits(num)
	fmt.Println(result)
}

func reverseBits(num uint32) uint32 {
	var ans = uint32(0)
	var power = uint32(31)
	for num != 0 {
		ans += (num & 1) << power
		num = num >> 1
		power -= 1
	}
	return ans
}
