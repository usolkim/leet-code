package main

import (
	"fmt"
)

func main() {
	print(partition(&ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, &ListNode{5, nil}}}}}}}, 3))
	print(partition(&ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, &ListNode{5, nil}}}}}}}, 5))
	print(partition(&ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, &ListNode{5, nil}}}}}}}, 1))
	print(partition(&ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, &ListNode{5, nil}}}}}}}, 2))
	print(partition(&ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, &ListNode{5, nil}}}}}}}, 0))
	print(partition(nil, 1))
	print(partition(&ListNode{1, nil}, 1))
	print(partition(&ListNode{1, &ListNode{2, nil}}, 2))
	print(partition(&ListNode{2, &ListNode{1, &ListNode{1, nil}}}, 2))
	print(partition(&ListNode{1, &ListNode{3, &ListNode{2, nil}}}, 3))
	print(partition(&ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, nil}}}}}}}, 5))

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

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var f *ListNode
	s := head
	for s != nil && s.Val < x {
		f = s
		s = s.Next
	}

	if s != nil {
		t := s

		for s.Next != nil {
			if s.Next.Val < x {
				if f != nil {
					f.Next = s.Next
					f = f.Next
				} else {
					f = s.Next
					head = f
				}
				s.Next = s.Next.Next
				f.Next = t
			} else {
				s = s.Next
			}
		}
	}

	return head
}
