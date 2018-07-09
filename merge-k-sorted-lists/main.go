package main

import (
	"fmt"
)

func main() {
	a3 := ListNode{Val: 5, Next: nil}
	a2 := ListNode{Val: 4, Next: &a3}
	a1 := ListNode{Val: 1, Next: &a2}

	// b3 := ListNode{Val: 4, Next: nil}
	// b2 := ListNode{Val: 3, Next: &b3}
	// b1 := ListNode{Val: 1, Next: &b2}

	// c2 := ListNode{Val: 6, Next: nil}
	// c1 := ListNode{Val: 2, Next: &c2}

	d4 := ListNode{Val: 26, Next: nil}
	d3 := ListNode{Val: 22, Next: &d4}
	d2 := ListNode{Val: 16, Next: &d3}
	d1 := ListNode{Val: 2, Next: &d2}

	lists := make([]*ListNode, 0)
	lists = append(lists, &a1)
	// lists = append(lists, &b1)
	// lists = append(lists, &c1)
	lists = append(lists, &d1)

	for l := mergeKLists(lists); l != nil; l = l.Next {
		fmt.Print(l.Val, ",")
	}
	fmt.Println()
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	} else if l == 1 {
		return lists[0]
	}
	return mergeTwoLists(mergeKLists(lists[:l/2]), mergeKLists(lists[l/2:]))
}

func mergeTwoLists(left *ListNode, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.Val < right.Val {
		left.Next = mergeTwoLists(left.Next, right)
		return left
	} else {
		right.Next = mergeTwoLists(left, right.Next)
		return right
	}
}
