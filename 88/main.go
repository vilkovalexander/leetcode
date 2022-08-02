package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 && len(nums1) != 0 {
		for i, _ := range nums1 {
			nums1[i] = nums2[i]
		}
		return
	}
	end := len(nums1) - 1
	m--
	n--

	for end >= 0 {
		if m < 0 {
			nums1[end] = nums2[n]
			end--
			n--
		}
		if n < 0 {
			//fmt.Println(nums1)
			return
		} //end
		if nums1[m] <= nums2[n] {
			nums1[end] = nums2[n]
			end--
			n--
		} else {
			nums1[end], nums1[m] = nums1[m], nums1[end]
			end--
			m--
		}
	}
	//fmt.Println(nums1)
}

func main() {
	merge([]int{2, 0}, 1, []int{1}, 1)
}
