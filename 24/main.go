package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, head}
	prev, curr := dummy, dummy.Next
	for curr != nil && curr.Next != nil {
		second := curr.Next
		third := second.Next

		second.Next = curr
		curr.Next = third
		prev.Next = second
		prev = prev.Next.Next
		curr = curr.Next
	}
	return dummy.Next
}

func main() {
	res := swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}})
	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
}
