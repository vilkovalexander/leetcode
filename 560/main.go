package main

func subarraySum(nums []int, k int) int {
	sum := 0
	res := 0
	prefixMap := map[int]int{0: 1}
	for _, num := range nums {
		sum += num
		diff := sum - k
		if val, ok := prefixMap[diff]; ok {
			res += val
		}
		prefixMap[sum]++
	}
	return res
}
