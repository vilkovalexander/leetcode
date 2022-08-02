package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := lowerBound(nums, target)
	if left >= len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	return []int{left, upperBound(nums, target) - 1}
}

func lowerBound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func upperBound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}
