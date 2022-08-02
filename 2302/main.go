package main

func countSubarrays(nums []int, k int64) int64 {
	var now, res int64

	left := 0
	for right, _ := range nums {
		now += int64(nums[right])

		for now*int64(right-left+1) >= k {
			now -= int64(nums[left])
			left++
		}
		res += int64(right - left + 1)
	}
	return res
}
