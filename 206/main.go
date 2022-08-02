package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var left, right *ListNode
	if head == nil || head.Next == nil {
		return head
	}
	left = nil
	right = head
	for right != nil {
		tmp := right.Next
		right.Next = left
		left = right
		right = tmp
	}
	return left
}
