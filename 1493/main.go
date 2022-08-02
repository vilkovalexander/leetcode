package main

import "fmt"

func longestSubarray(nums []int) int {
	cur := 0
	last := 0
	longest := 0
	nums = append(nums, 0)
	for i, val := range nums {
		if val == 1 {
			cur++
		} else {
			if last+cur > longest {
				longest = last + cur
			}
			last = 0
			if i+1 < len(nums) && nums[i+1] == 1 {
				last = cur
			}
			cur = 0
		}
	}
	if longest == len(nums)-1 {
		return longest - 1
	}
	return longest
}

func main() {
	fmt.Println(longestSubarray([]int{0, 0, 0, 0}))
}
