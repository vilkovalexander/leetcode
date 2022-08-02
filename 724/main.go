package main

import "fmt"

func pivotIndex(nums []int) int {
	suffix := make([]int, len(nums))
	prefix := make([]int, len(nums))
	prefix[0] = 0
	suffix[len(suffix)-1] = 0
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] + nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		if suffix[i] == prefix[i] {
			return i
		}
	}
	return -1

}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6, 7}))
}
