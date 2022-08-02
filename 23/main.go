package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return &ListNode{0, nil}
	}
	if len(lists) == 1 {
		return lists[0]
	}
	return MergeLists(lists[:len(lists)/2], lists[len(lists)/2:])
}

func MergeLists(firstSlice, secondSLice []*ListNode) *ListNode {
	var first, second *ListNode
	if len(firstSlice) > 1 {
		first = MergeLists(firstSlice[:len(firstSlice)/2], firstSlice[len(firstSlice)/2:])
	} else {
		first = firstSlice[0]
	}
	if len(secondSLice) > 1 {
		second = MergeLists(secondSLice[:len(secondSLice)/2], secondSLice[len(secondSLice)/2:])
	} else {
		second = secondSLice[0]
	}
	return mergeTwoLists(first, second)
}

// [[1,4,5],[1,2,3,4,6]]
func mergeTwoLists(first, second *ListNode) *ListNode {
	merged := &ListNode{0, nil}
	tmp := merged
	for {
		if first == nil {
			tmp.Next = second
			merged = merged.Next
			return merged
		}
		if second == nil {
			tmp.Next = first
			merged = merged.Next
			return merged
		}

		if first.Val <= second.Val {
			tmp.Next = &ListNode{first.Val, nil}
			first = first.Next
			tmp = tmp.Next
		} else {
			tmp.Next = &ListNode{second.Val, nil}
			second = second.Next
			tmp = tmp.Next
		}
	}
}

//[[1,4,5],[1,3,4],[2,6]]
// [[1,4,5],[1,2,3,4,6]]
func main() {
	testCase := []*ListNode{
		&ListNode{1, &ListNode{4, &ListNode{5, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
		&ListNode{2, &ListNode{6, nil}},
	}
	res := mergeKLists(testCase)
	fmt.Println(res)

}
