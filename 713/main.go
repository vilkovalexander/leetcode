package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	j := 0
	count := 0
	nowProduct := 1
	for i, _ := range nums {
		nowProduct *= nums[i]
		for nowProduct >= k {
			nowProduct /= nums[j]
			j++
		}
		count += i - j + 1
	}

	return count
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 100, 6}, 100))
}
