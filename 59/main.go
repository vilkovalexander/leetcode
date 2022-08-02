package main

import "fmt"

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}
	c1, r1, c2, r2 := -1, -1, n, n
	count := 1
	for count <= n*n {
		c1++
		c2--
		r1++
		r2--
		for c := c1; c <= c2; c++ {
			result[r1][c] = count
			count++
		}
		for r := r1 + 1; r <= r2; r++ {
			result[r][c2] = count
			count++
		}
		for c := c2 - 1; c >= c1; c-- {
			result[r2][c] = count
			count++
		}
		for r := r2 - 1; r >= r1+1; r-- {
			result[r][c1] = count
			count++
		}
	}
	return result
}

func main() {
	for _, val := range generateMatrix(4) {
		fmt.Println(val)
	}
}
