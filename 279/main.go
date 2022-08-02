package main

import "fmt"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = n
	}
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			square := j * j
			if square > i {
				break
			}
			dp[i] = min(dp[i], dp[i-square]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(numSquares(16))
}
