package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBST(root *TreeNode, prev *int) bool {
	if root == nil {
		return true
	}
	if isBST(root.Left, prev) == false {
		return false
	}
	if root.Val <= *prev {
		return false
	}
	*prev = root.Val
	return isBST(root.Right, prev)

}

func isValidBST(root *TreeNode) bool {
	var prev = math.MinInt
	return isBST(root, &prev)
}

func main() {
	fmt.Println(isValidBST(&TreeNode{5,
		&TreeNode{1, nil, nil},
		&TreeNode{5,
			&TreeNode{6, nil, nil}, &TreeNode{8, nil, nil}}}))
}
