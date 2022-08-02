package main

import "fmt"

func search(nums []int, target int) int {
	var mid int
	l, r := 0, len(nums)-1
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[0] {
			if target < nums[mid] && target < nums[0] || target > nums[mid] {
				l = mid + 1
			} else if target < nums[mid] && target >= nums[0] {
				r = mid - 1
			}
		} else {
			if target < nums[mid] {
				r = mid - 1
			} else if target > nums[len(nums)-1] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}

func main() {
	c := search([]int{3, 1}, 1)
	fmt.Println(c)
}
