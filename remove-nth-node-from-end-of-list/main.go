package main

import (
	"fmt"
)

func main() {
	// n5 := ListNode{
	// 	Val:  5,
	// 	Next: nil,
	// }
	// n4 := ListNode{
	// 	Val:  4,
	// 	Next: &n5,
	// }
	// n3 := ListNode{
	// 	Val:  3,
	// 	Next: &n4,
	// }
	// n2 := ListNode{
	// 	Val:  2,
	// 	Next: &n3,
	// }
	// n1 := ListNode{
	// 	Val:  1,
	// 	Next: &n2,
	// }

	n1 := ListNode{
		Val:  1,
		Next: nil,
	}

	n := *removeNthFromEnd(&n1, 1)

	for {
		fmt.Print(n.Val)
		if n.Next == nil {
			break
		}
		n = *n.Next
	}

}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr := head
	bef := head
	t := 0
	for curr.Next != nil {
		t++
		if t > n {
			bef = bef.Next
		}
		curr = curr.Next
	}
	if n > t {
		return head.Next
	}
	bef.Next = bef.Next.Next
	return head
}
