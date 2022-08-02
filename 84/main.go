package main

import (
	"fmt"
)

type pair struct {
	index int
	h     int
}

func largestRectangleArea(heights []int) int {
	maxi := 0
	stack := make([]pair, 0)
	for i, height := range heights {
		start := i
		for len(stack) != 0 && stack[len(stack)-1].h > height {
			ind, heig := stack[len(stack)-1].index, stack[len(stack)-1].h
			maxi = max(maxi, heig*(i-ind))
			stack = stack[:len(stack)-1]
			start = ind
		}
		stack = append(stack, pair{start, height})
	}
	for _, val := range stack {
		maxi = max(maxi, val.h*(len(heights)-val.index))
	}

	return maxi
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
