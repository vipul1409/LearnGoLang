package bst

import (
	"fmt"
	"math/rand"
)

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

func CreateNode(val int) *TreeNode {
	t := TreeNode{Val: val, Left: nil, Right: nil}
	return &t
}

func Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return CreateNode(val)
	}
	if root.Val < val {
		root.Right = Insert(root.Right, val)
	} else {
		root.Left = Insert(root.Left, val)
	}
	return root
}

func PrintInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PrintInOrder(root.Left)
	fmt.Println(root.Val)
	PrintInOrder(root.Right)
}

func BuildTree(n, k int) *TreeNode {
	var t *TreeNode
	for _, v := range rand.Perm(n) {
		t = Insert(t, (1+v)*k)
	}
	return t
}
