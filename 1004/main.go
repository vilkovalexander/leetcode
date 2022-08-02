package main

import "fmt"

func longestOnes(nums []int, k int) int {
	var result, j int
	for i, val := range nums {
		if val == 0 {
			k--
		}
		for k < 0 {
			if nums[j] == 0 {
				k++
			}
			j++
		}
		result = max(result, i-j+1)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
}
