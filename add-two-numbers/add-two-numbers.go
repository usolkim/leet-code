package main

import (
	"fmt"
)

func main() {
	l11 := ListNode{}
	l12 := ListNode{}
	l13 := ListNode{}

	l11.Val = 2
	l11.Next = &l12
	l12.Val = 4
	l12.Next = &l13
	l13.Val = 3
	l13.Next = nil

	l21 := ListNode{}
	l22 := ListNode{}
	l23 := ListNode{}

	l21.Val = 5
	l21.Next = &l22
	l22.Val = 6
	l22.Next = &l23
	l23.Val = 7
	l23.Next = nil

	res := addTwoNumbers(&l11, &l21)

	arr := []int{}

	for next := res; next != nil; next = next.Next {
		nn := *next
		arr = append(arr, nn.Val)
	}

	fmt.Println(arr)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		l1 = &ListNode{}
	}

	if l2 == nil {
		l2 = &ListNode{}
	}

	l3 := ListNode{}

	l3.Val = l1.Val + l2.Val

	if l1.Val+l2.Val > 9 {
		l3.Val -= 10
		if l1.Next == nil {
			l1n := ListNode{}
			l1n.Val = 1
			l1.Next = &l1n
		} else {
			l1n := *l1.Next
			l1n.Val = l1n.Val + 1
			l1.Next = &l1n
		}
	}

	l3.Next = addTwoNumbers(l1.Next, l2.Next)

	return &l3
}
