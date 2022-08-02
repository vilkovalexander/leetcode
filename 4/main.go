package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		if len(nums2) == 1 {
			return float64(nums2[0])
		}
		if len(nums2)%2 == 1 {
			return float64(nums2[len(nums2)/2]+nums2[len(nums2)/2+1]) / 2
		} else {
			return float64(nums2[len(nums2)/2+1])
		}
	}
	if len(nums2) == 0 {
		if len(nums1) == 1 {
			return float64(nums1[0])
		}
		if len(nums1)%2 == 1 {
			return float64(nums1[len(nums1)/2]+nums1[len(nums1)/2+1]) / 2
		} else {
			return float64(nums1[len(nums1)/2+1])
		}
	}
	var aLeft, aRight, bLeft, bRight int
	a, b := nums1, nums2
	if len(a) > len(b) {
		b, a = a, b
	}
	total := len(a) + len(b)
	half := total / 2
	left, right := 0, len(a)-1
	for {
		mid := (left + right) / 2
		j := half - mid - 2
		if mid >= 0 {
			aLeft = a[mid]
		} else {
			aLeft = math.MinInt
		}
		if mid+1 < len(a) {
			aRight = a[mid+1]
		} else {
			aRight = math.MaxInt
		}
		if j >= 0 {
			bLeft = b[j]
		} else {
			bLeft = math.MinInt
		}
		if j+1 < len(b) {
			bRight = b[j+1]
		} else {
			bRight = math.MaxInt
		}
		if aLeft <= bRight && bLeft <= aRight {
			if total%2 == 1 {
				return float64(min(aRight, bRight))
			}
			return float64(max(aLeft, bLeft)+min(aRight, bRight)) / 2
		} else if aLeft > bRight {
			if mid == 0 {
				right = mid - 2
				continue
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}
