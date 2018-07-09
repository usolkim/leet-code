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

	for l := swapPairs(&a1); l != nil; l = l.Next {
		fmt.Print(l.Val, ",")
	}
	fmt.Println()
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var tmp *ListNode
	if head.Next.Next != nil {
		tmp = swapPairs(head.Next.Next)
	}

	ret := head.Next

	head.Next.Next = head
	head.Next = tmp

	return ret
}

// func jump(head *ListNode, n int) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	curr := head
// 	for i := 0; i < n; i++ {
// 		if curr.Next == nil {
// 			return curr
// 		}
// 		curr = curr.Next
// 	}
// }
