package main

import "fmt"

func main() {

	result := nCrForBigValues(5, 2, 13)
	fmt.Println(result)
}

func nCr(A int, B int, C int) int {
	ans := 0

	ans = fact(A) / (fact(B) * fact(A-B))

	return ans % C
}

func fact(n int) int {
	if n <= 1 {
		return n
	}

	return (n) * fact(n-1)
}

func nCrForBigValues(A int, B int, C int) int {
	var dp [][]int

	for i := 0; i < A+1; i++ {
		var temp []int
		for j := 0; j < B+1; j++ {
			temp = append(temp, 0)
		}
		dp = append(dp, temp)
	}

	for i := 0; i <= A; i++ {
		for j := 0; j <= min(i, B); j++ {
			if j == i || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = (dp[i-1][j-1]%C + dp[i-1][j]%C) % C
			}
		}
	}

	return dp[A][B] % C
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
