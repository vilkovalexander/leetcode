package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 1 {
		return []string{fmt.Sprint(strconv.Itoa(nums[0]))}
	}
	result := make([]string, 0)
	j := 0

	for i, _ := range nums {
		if i+1 < len(nums) && nums[i+1]-nums[i] > 1 {
			if i-j == 0 {
				result = append(result, fmt.Sprint(nums[j]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[j], nums[i]))
			}
			j = i + 1
		} else if i+1 >= len(nums) {
			if i-j == 0 {
				result = append(result, fmt.Sprint(nums[j]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[j], nums[i]))
			}
		}
	}
	return result
}

func main() {
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
