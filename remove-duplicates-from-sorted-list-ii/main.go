package main

import (
	"fmt"
)

func main() {
	print(deleteDuplicates(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}))
	print(deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}}}}))
	print(deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, nil}}}}}}))
	print(deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, nil}}}}}}))
	print(deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}}))
	print(deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{3, nil}}}}}}}}}}))
	print(deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{3, &ListNode{3, nil}}}}}}}}}}))
	print(deleteDuplicates(&ListNode{1, nil}))
	print(deleteDuplicates(&ListNode{1, &ListNode{2, nil}}))
	print(deleteDuplicates(nil))
}

func print(head *ListNode) {
	fmt.Print("\"")
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println("\"")
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := &ListNode{}
	curr := res
	if head.Val != head.Next.Val {
		curr.Next = head
		curr = curr.Next
	}
	for head.Next.Next != nil {
		if head.Next.Next != nil && head.Val != head.Next.Val && head.Next.Val != head.Next.Next.Val {
			curr.Next = head.Next
			curr = curr.Next
		}
		head = head.Next
	}
	if head.Val != head.Next.Val {
		curr.Next = head.Next
		curr = curr.Next
	}
	curr.Next = nil
	return res.Next
}
