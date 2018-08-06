package main

import (
	"fmt"
)

func main() {
	fmt.Println(inorderTraversal(&TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}))
	fmt.Println(inorderTraversal(&TreeNode{1, nil, nil}))
	fmt.Println(inorderTraversal(nil))
	fmt.Println(inorderTraversal(&TreeNode{1, &TreeNode{2, &TreeNode{4, &TreeNode{8, nil, nil}, &TreeNode{9, nil, nil}}, &TreeNode{5, &TreeNode{10, nil, nil}, &TreeNode{11, nil, nil}}}, &TreeNode{3, &TreeNode{6, &TreeNode{12, nil, nil}, &TreeNode{13, nil, nil}}, &TreeNode{7, &TreeNode{14, nil, nil}, &TreeNode{15, nil, nil}}}}))
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	curr := root
	var parent *TreeNode

	for curr != nil {
		if curr.Left == nil {
			res = append(res, curr.Val)
			curr = curr.Right
		} else {
			parent = curr.Left
			for parent.Right != nil {
				parent = parent.Right
			}
			parent.Right = curr
			temp := curr
			curr = curr.Left
			temp.Left = nil
		}
	}
	return res
}
