package main

import (
	"sort"
)

func missingNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] < nums[j] {
			return true
		} else {
			return false
		}
	})
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 {
			return nums[i] - 1
		}
	}
	if nums[0] != 0 {
		return 0
	}
	return nums[len(nums)-1] + 1
}
