package main

import (
	"fmt"
)

func main() {
	fmt.Println(isSymmetric(&TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}}))
	fmt.Println(isSymmetric(&TreeNode{1, &TreeNode{2, nil, &TreeNode{3, nil, nil}}, &TreeNode{2, nil, &TreeNode{3, nil, nil}}}))
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := make([]*TreeNode, 0)
	r := make([]*TreeNode, 0)

	l = append(l, root.Left)
	r = append(r, root.Right)

	res := true

	for i := 0; i < len(l); i++ {
		p := l[i]
		q := r[i]
		if p == nil && q == nil {
			continue
		}
		if (p == nil && q != nil) || (p != nil && q == nil) || (p.Val != q.Val) {
			res = false
			break
		}

		l = append(l, p.Left)
		l = append(l, p.Right)
		r = append(r, q.Right)
		r = append(r, q.Left)
	}

	return res
	// return recursive(root.Left, root.Right)
}

// func recursive(p *TreeNode, q *TreeNode) bool {
// 	if p == nil {
// 		return q == nil
// 	} else if q == nil {
// 		return false
// 	}
// 	return p.Val == q.Val && recursive(p.Left, q.Right) && recursive(p.Right, q.Left)
// }
