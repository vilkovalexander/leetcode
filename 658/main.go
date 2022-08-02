package main

import (
	"fmt"
	"sort"
)

func lowerBound(nums []int, x int) int {
	var mid int
	l, r := 0, len(nums)-1
	for l < r {
		mid = (l + r) / 2
		if nums[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func upperBound(nums []int, x int) int {
	var mid int
	l, r := 0, len(nums)-1
	for l < r {
		mid = (l + r) / 2
		if nums[mid] > x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l == 0 {
		return l
	}
	return l - 1
}

func findClosestElements(arr []int, k int, x int) []int {
	var lower, upper int
	result := make([]int, k)
	for i := 0; i < k; i++ {
		lower = lowerBound(arr, x)
		upper = upperBound(arr, x)
		tmp1, tmp2 := abs(arr[lower], x), abs(arr[upper], x)
		if tmp1 < tmp2 {
			result[i] = arr[lower]

		} else if tmp1 == tmp2 {
			if arr[lower] <= arr[upper] {
				result[i] = arr[lower]
			} else {
				result[i] = arr[upper]
				lower = upper
			}
		} else {
			result[i] = arr[upper]
			lower = upper
		}
		if lower == 0 && len(arr) > 1 {
			arr = arr[1:]
		} else if lower == 0 && len(arr) == 1 {
			arr = []int{}
		} else if lower == len(arr)-1 {
			arr = arr[:len(arr)-1]
		} else {
			arr = append(arr[:lower], arr[lower+1:]...)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i] <= result[j] {
			return true
		}
		return false
	})
	return result
}

func abs(a, b int) int {
	tmp := a - b
	if tmp < 0 {
		return -tmp
	}
	return tmp
}

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
}
