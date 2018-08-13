package main

import (
	"fmt"
	"strconv"
)

func main() {
	list := generateTrees(3)
	for _, nodes := range list {

		fmt.Println(nodes)
	}
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	} else if n == 1 {
		return []*TreeNode{&TreeNode{1, nil, nil}}
	}

	var dp func(int, int) []*TreeNode
	memo := make(map[string][]*TreeNode, 0)
	dp = func(start, end int) []*TreeNode {
		if start > end {
			return nil
		}

		key := strconv.Itoa(start) + "_" + strconv.Itoa(end)
		if v, exist := memo[key]; exist {
			return v
		}
		res := make([]*TreeNode, 0)

		if start == end {
			res = append(res, &TreeNode{start, nil, nil})
		} else {
			for i := start; i <= end; i++ {
				left := dp(start, i-1)
				right := dp(i+1, end)
				if left == nil {

					for _, r := range right {
						root := &TreeNode{i, nil, r}
						res = append(res, root)
					}
				} else if right == nil {
					for _, l := range left {
						root := &TreeNode{i, l, nil}
						res = append(res, root)
					}
				} else {
					for _, l := range left {
						for _, r := range right {
							root := &TreeNode{i, l, r}
							res = append(res, root)
						}
					}
				}
			}
		}

		memo[key] = res
		return res
	}

	return dp(1, n)
}

func recursive(start, end int) []*TreeNode {
	if start > end {
		return nil
	} else if start == end {
		return []*TreeNode{&TreeNode{start, nil, nil}}
	}

	res := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		left := recursive(start, i-1)
		right := recursive(i+1, end)
		if left == nil {

			for _, r := range right {
				root := &TreeNode{i, nil, r}
				res = append(res, root)
			}
		} else if right == nil {
			for _, l := range left {
				root := &TreeNode{i, l, nil}
				res = append(res, root)
			}
		} else {
			for _, l := range left {
				for _, r := range right {
					root := &TreeNode{i, l, r}
					res = append(res, root)
				}
			}
		}
	}
	return res
}
