package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	seria   string
	deseria *TreeNode
}

func Constructor() Codec {
	return Codec{"", nil}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}

	res := make([]int, 0)

	dfsSeria(root, &res)
	result := "["
	for _, val := range res[:len(res)-1] {
		result += strconv.Itoa(val) + ","
	}
	result += strconv.Itoa(res[len(res)-1]) + "]"
	return result
}

func dfsSeria(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	if root.Left != nil {
		dfsSeria(root.Left, res)
	}
	if root.Right != nil {
		dfsSeria(root.Right, res)
	}
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	preorder := make([]int, 0)
	for _, val := range strings.Split(data[1:len(data)-1], ",") {
		n, _ := strconv.Atoi(val)
		preorder = append(preorder, n)
	}
	return bstFromPreorder(preorder)
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
	codec := Constructor()
	root := codec.deserialize("[8,5,1,7,10,12]")
	fmt.Println(codec.serialize(root))
}
