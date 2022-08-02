package main

import "fmt"

func sortedSquares(nums []int) []int {
	minus := make([]int, 0)
	plus := make([]int, 0)
	for _, val := range nums {
		if val < 0 {
			minus = append(minus, val*val)
		} else {
			plus = append(plus, val*val)
		}
	}
	i := 0
	j := len(minus) - 1
	m := 0
	for ; m < len(nums); m++ {
		if i >= len(plus) {
			for k := m; k < len(nums); k++ {
				nums[k] = minus[j]
				j--
			}
			break
		}
		if j < 0 {
			for k := m; k < len(nums); k++ {
				nums[k] = plus[i]
				i++
			}
			break
		}
		if minus[j] <= plus[i] {
			nums[m] = minus[j]
			j--
		} else {
			nums[m] = plus[i]
			i++
		}
	}
	return nums
}

func main() {
	fmt.Println(sortedSquares([]int{-5, -3, -2, -1}))
}
