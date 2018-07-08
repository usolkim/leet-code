package main

import (
	"fmt"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a3 := ListNode{Val: 16, Next: nil}
	a2 := ListNode{Val: 13, Next: &a3}
	a1 := ListNode{Val: 2, Next: &a2}

	b3 := ListNode{Val: 7, Next: nil}
	b2 := ListNode{Val: 4, Next: &b3}
	b1 := ListNode{Val: 1, Next: &b2}

	node := *mergeTwoLists(&a1, &b1)

	for node.Next != nil {
		fmt.Print(node.Val)
		node = *node.Next
	}

	fmt.Println(node.Val)

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	node := ListNode{}
	c := &node

	for {
		if l1 == nil {
			c.Next = l2
			break
		}
		if l2 == nil {
			c.Next = l1
			break
		}
		if l1.Val <= l2.Val {
			c.Next = l1
			l1 = l1.Next
		} else {
			c.Next = l2
			l2 = l2.Next
		}
		c = c.Next
	}

	return node.Next
}
