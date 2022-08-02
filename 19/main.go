package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := head
	for n > 0 && fast != nil {
		fast = fast.Next
		n--
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func main() {
	fmt.Println(removeNthFromEnd(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2))
}
