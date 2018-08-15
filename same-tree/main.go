package main

import (
	"fmt"
)

func main() {
	fmt.Println(isSameTree(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}))
	fmt.Println(isSameTree(&TreeNode{1, &TreeNode{2, nil, nil}, nil}, &TreeNode{1, nil, &TreeNode{2, nil, nil}}))
	fmt.Println(isSameTree(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{1, nil, nil}}, &TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}))
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		return q == nil
	} else if q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}

	if !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}
