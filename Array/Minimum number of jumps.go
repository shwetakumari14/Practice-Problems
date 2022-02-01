package main

import "fmt"

func main() {

	A := []int{3, 6, 6, 10, 6, 1, 3, 1, 10, 1, 1, 10, 1, 7, 7, 2, 3, 1, 2, 4, 5, 8, 7, 2, 6, 8, 6, 7, 5, 4, 10, 4, 8, 10, 8}
	result := minJump(A)
	fmt.Println(result)
}

func minJump(A []int) int {
	if len(A) <= 1 {
		return 0
	}
	if A[0] == 0 {
		return -1
	}
	jumps, maxRechable, steps := 1, A[0], A[0]

	for i := 1; i < len(A); i++ {
		if i == len(A)-1 {
			return jumps
		}
		maxRechable = max(maxRechable, A[i]+i)
		steps--
		if steps == 0 {
			jumps++
			if i >= maxRechable {
				return -1
			}
			steps = maxRechable - i
		}
	}

	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
