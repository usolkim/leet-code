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
	res := head
	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return res
}
