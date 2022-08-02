package main

import "fmt"

func trap(height []int) int {
	var max int
	stack := make([]int, 0)
	count := 0
	for _, h := range height {
		if len(stack) == 0 {
			stack = append(stack, h)
			continue
		}
		if stack[0] > h {
			stack = append(stack, h)
		} else if len(stack) == 1 {
			stack[0] = h
		} else {
			for len(stack) != 1 {
				count += stack[0] - stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			stack[0] = h

		}
	}
	if len(stack) > 2 {
		if stack[len(stack)-1] <= stack[len(stack)-2] {
			max = stack[len(stack)-2]
			stack = stack[:len(stack)-2]
		} else {
			max = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		for len(stack) != 1 {
			if stack[len(stack)-1] <= max {
				count += max - stack[len(stack)-1]
			} else {
				max = stack[len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		}
	}

	return count
}

func main() {
	fmt.Println(trap([]int{0, 1, 2, 0, 3, 0, 1, 2, 0, 0, 4, 2, 1, 2, 5, 0, 1, 2, 0, 2}))
}
