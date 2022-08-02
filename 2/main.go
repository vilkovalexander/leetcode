package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{0, nil}
	curr := result
	var x, y, div, sum int
	for l1 != nil || l2 != nil || div != 0 {
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		} else {
			x = 0
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		} else {
			y = 0
		}
		sum = x + y + div
		div = sum / 10
		newNode := &ListNode{sum % 10, nil}
		curr.Next = newNode
		curr = newNode
	}
	return result.Next
}
