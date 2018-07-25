package main

import "fmt"

func main() {
	fmt.Println(rotateRight(nil, 0))
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	s := head
	l := 1
	curr := head
	for curr.Next != nil {
		curr = curr.Next
		l++
	}
	k %= l
	if k == 0 {
		return head
	}
	for i := 0; i < l-k-1; i++ {
		head = head.Next
	}
	res := head.Next
	curr = head.Next
	head.Next = nil
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = s
	return res
}
