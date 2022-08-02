package main

import (
	"fmt"
	"sort"
)

func deleteAndEarn(nums []int) int {
	hashmap := make(map[int]int, 0)
	numsUniq := make([]int, 0)

	for _, val := range nums {
		hashmap[val]++
	}
	for key := range hashmap {
		numsUniq = append(numsUniq, key)
	}
	sort.Slice(numsUniq, func(i, j int) bool {
		if numsUniq[i] < numsUniq[j] {
			return true
		}
		return false
	})
	e1, e2 := 0, 0
	for i, val := range numsUniq {
		currEarn := val * hashmap[val]
		if i > 0 && val == numsUniq[i-1]+1 {
			tmp := e2
			e2 = max(e2, currEarn+e1)
			e1 = tmp
		} else {
			tmp := e2
			e2 = currEarn + e2
			e1 = tmp
		}
	}
	return e2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(deleteAndEarn([]int{3, 4, 2}))
}
