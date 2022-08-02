package main

import "fmt"

func twoSum(nums []int, target int) []int {
	diff := make(map[int]int, 0)
	for index, el := range nums {
		if _, ok := diff[el]; ok {
			return []int{diff[el], index}
		}
		diff[target-el] = index
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
