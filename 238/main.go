package main

import "fmt"

func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	postfix := make([]int, len(nums))
	ansver := make([]int, len(nums))
	countPostfix := 1
	countPrefix := 1
	prefix[0] = 1
	postfix[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		countPrefix *= nums[i-1]
		prefix[i] = countPrefix
	}
	for i := len(nums) - 2; i >= 0; i-- {
		countPostfix *= nums[i+1]
		postfix[i] = countPostfix
	}
	for i := 0; i < len(nums); i++ {
		//if postfix[i] == 31 {
		//	ansver[i] = prefix[i]
		//} else if prefix[i] == 31 {
		//	ansver[i] = postfix[i]
		//} else {
		ansver[i] = postfix[i] * prefix[i]

	}
	return ansver
}

func main() {
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
