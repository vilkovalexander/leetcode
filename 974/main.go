package main

import "fmt"

func subarraysDivByK(nums []int, k int) int {
	hashmap := map[int]int{0: 1}
	res := 0
	sum := 0
	for _, num := range nums {
		sum = (sum + num) % k
		if sum < 0 {
			sum = (sum + k) % k
		}
		res += hashmap[sum]
		hashmap[sum]++
	}
	return res
}

func main() {
	c := subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5)
	fmt.Println(c)
}
