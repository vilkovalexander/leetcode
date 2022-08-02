package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var e1, e2, tmp int

	for _, val := range nums {
		tmp = e2
		e2 = max(e1+val, e2)
		e1 = tmp
	}
	return e2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
