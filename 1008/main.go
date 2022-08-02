package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{preorder[0], nil, nil}
	for _, val := range preorder[1:] {
		root = insert(root, val)
	}
	return root
}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if root.Val > val {
		root.Left = insert(root.Left, val)
	}
	if root.Val < val {
		root.Right = insert(root.Right, val)
	}
	return root
}

func main() {
	res := bstFromPreorder([]int{8, 5, 1, 7, 10, 12})
	fmt.Println(res)
}
