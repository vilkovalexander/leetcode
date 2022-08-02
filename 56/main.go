package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return [][]int{intervals[0]}
	}
	result := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else {
			return false
		}
	})
	i := 0
	j := 1
	for j < len(intervals) {
		if intervals[i][1] >= intervals[j][0] {
			intervals[i][1] = max(intervals[j][1], intervals[i][1])
			j += 1
			if j == len(intervals) {
				result = append(result, intervals[i])
			}
		} else {
			result = append(result, intervals[i])
			i = j
			j++
			if j == len(intervals) {
				result = append(result, intervals[i])
			}
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
}
