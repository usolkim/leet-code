package main

import (
	"fmt"
)

func main() {
	head := reverseBetween(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 1, 5)
	fmt.Print("\"")
	for {
		fmt.Print(head.Val)
		if head.Next != nil {
			head = head.Next
			fmt.Print("->")
		} else {
			break
		}
	}
	fmt.Println("\"")
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m < n {

		var dfs func(*ListNode, int) (*ListNode, *ListNode, *ListNode)
		dfs = func(curr *ListNode, c int) (*ListNode, *ListNode, *ListNode) {
			if c == n {
				return curr, curr, curr.Next
			}
			last, next, nnext := dfs(curr.Next, c+1)
			next.Next = curr
			return last, curr, nnext
		}

		temp := head
		for i := 0; i < m-2; i++ {
			temp = temp.Next
		}

		var t *ListNode
		if m == 1 {
			t = temp
		} else {
			t = temp.Next
		}

		last, next, nnext := dfs(t, m)
		if m == 1 {
			head = last
		} else {
			temp.Next = last
		}
		next.Next = nnext

	}

	return head
}
