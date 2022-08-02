package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	dummy := &ListNode{0, head}
	fast, slow, begin := dummy, dummy, dummy
	n := k
	for n > 0 && begin != nil {
		begin = begin.Next
		fast = fast.Next
		n--
	}
	if begin == nil {
		return dummy.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Val, begin.Val = begin.Val, slow.Val
	return dummy.Next

}

func main() {
	res := swapNodes(&ListNode{1, &ListNode{2, nil}}, 2)
	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
}
