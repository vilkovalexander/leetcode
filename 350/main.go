package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]int, 0)
	map2 := make(map[int]int, 0)
	for _, val := range nums1 {
		map1[val]++
	}
	for _, val := range nums2 {
		map2[val]++
	}
	result := make([]int, 0)
	for key, count := range map1 {
		if val, ok := map2[key]; ok {
			for i := 0; i < min(val, count); i++ {
				result = append(result, key)
			}
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}
