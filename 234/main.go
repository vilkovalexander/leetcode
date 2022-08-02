package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var dummy *ListNode
	slow := head
	fast := head

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
	}
	dummy = nil
	for slow != nil {
		tmp := slow.Next
		slow.Next = dummy
		dummy = slow
		slow = tmp
	}
	for dummy != nil && head != nil {
		if dummy.Val != head.Val {
			return false
		}
		dummy = dummy.Next
		head = head.Next
	}
	return true
}
