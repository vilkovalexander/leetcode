package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	pointer := dummy
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			pointer.Next = list2
			list2 = list2.Next
			pointer = pointer.Next
		} else {
			pointer.Next = list1
			list1 = list1.Next
			pointer = pointer.Next
		}
	}
	if list1 == nil {
		pointer.Next = list2
	}
	if list2 == nil {
		pointer.Next = list1
	}
	return dummy.Next
}

func main() {
	res := mergeTwoLists(&ListNode{1, &ListNode{3, &ListNode{5, nil}}}, &ListNode{2, &ListNode{4, nil}})
	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
}