package bst

import (
	"fmt"
)

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

func (n TreeNode) String() string {
	s := fmt.Sprintf("Node val is ; %v", n.Val)
	return s

}
