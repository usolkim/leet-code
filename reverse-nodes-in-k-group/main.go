package main

import "fmt"

func main() {

	a8 := ListNode{Val: 8, Next: nil}
	a7 := ListNode{Val: 7, Next: &a8}
	a6 := ListNode{Val: 6, Next: &a7}
	a5 := ListNode{Val: 5, Next: &a6}
	a4 := ListNode{Val: 4, Next: &a5}
	a3 := ListNode{Val: 3, Next: &a4}
	a2 := ListNode{Val: 2, Next: &a3}
	a1 := ListNode{Val: 1, Next: &a2}

	for l := reverseKGroup(&a1, 6); l != nil; l = l.Next {
		fmt.Print(l.Val, ",")
	}
	fmt.Println()

	a8 = ListNode{Val: 8, Next: nil}
	a7 = ListNode{Val: 7, Next: &a8}
	a6 = ListNode{Val: 6, Next: &a7}
	a5 = ListNode{Val: 5, Next: &a6}
	a4 = ListNode{Val: 4, Next: &a5}
	a3 = ListNode{Val: 3, Next: &a4}
	a2 = ListNode{Val: 2, Next: &a3}
	a1 = ListNode{Val: 1, Next: &a2}

	for l := reverseKGroup(&a1, 9); l != nil; l = l.Next {
		fmt.Print(l.Val, ",")
	}
	fmt.Println()
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}
	if head == nil {
		return nil
	}
	last := jump(head, k-1)
	if last == nil {
		return head
	}
	tmp := last.Next

	dbNode := jump(head, k*2-1)
	if dbNode == nil {
		dbNode = last.Next
	}

	reverse(head, k, 1)
	head.Next = dbNode

	reverseKGroup(tmp, k)

	return last
}

func reverse(node *ListNode, k int, n int) {
	if k == n {
		return
	}
	n++
	reverse(node.Next, k, n)
	node.Next.Next = node
}

func jump(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	for i := 1; i <= k; i++ {
		if head.Next == nil {
			return nil
		}
		head = head.Next
	}
	return head
}
