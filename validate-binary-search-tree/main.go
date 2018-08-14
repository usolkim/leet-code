package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValidBST(&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}))
	fmt.Println(isValidBST(&TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}}}))
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	MIN = -1 << 31
	MAX = 1<<31 - 1
)

func isValidBST(root *TreeNode) bool {
	return recursive(root, MIN, MAX)
}

func recursive(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	} else if min != MIN && node.Val < min || max != MAX && node.Val > max {
		return false
	}
	if node.Left != nil {
		if !recursive(node.Left, min, node.Val-1) {
			return false
		}
	}
	if node.Right != nil {
		if !recursive(node.Right, node.Val+1, max) {
			return false
		}
	}
	return true
}
