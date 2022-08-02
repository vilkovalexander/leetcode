package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	histogram := make([][]int, 0)
	for j, vector := range matrix {
		tmpVector := make([]int, len(matrix[0]))
		if len(histogram) == 0 {
			for i, val := range vector {
				if val == '1' {
					tmpVector[i] = 1
				} else {
					tmpVector[i] = 0
				}

			}
			histogram = append(histogram, tmpVector)
			continue
		}
		for i, val := range vector {
			if val == '0' {
				tmpVector[i] = 0
			} else {
				tmpVector[i] = histogram[j-1][i] + 1
			}
		}
		histogram = append(histogram, tmpVector)
	}
	maxi := 0
	for _, hist := range histogram {
		maxi = max(maxi, largestRectangleArea(hist))
	}
	return maxi
}

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
	max := maximalRectangle([][]byte{{'1', '0', '1', '0', '0'}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}})
	fmt.Println(max)
}
