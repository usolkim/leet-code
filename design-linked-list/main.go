package main

import (
	"fmt"
)

func main() {
	linkedList := new(MyLinkedList)
	linkedList.Print()
	linkedList.AddAtHead(1)
	linkedList.Print()
	linkedList.AddAtTail(3)
	linkedList.Print()
	linkedList.AddAtIndex(1, 2) // linked list becomes 1->2->3
	linkedList.Print()
	fmt.Println(linkedList.Get(0)) // returns 1
	fmt.Println(linkedList.Get(1)) // returns 2
	fmt.Println(linkedList.Get(2)) // returns 3
	fmt.Println(linkedList.Get(3)) // returns -1
	linkedList.Print()
	linkedList.DeleteAtIndex(1) // now the linked list is 1->3
	linkedList.Print()
	linkedList.DeleteAtIndex(2) // now the linked list is 1->3
	linkedList.Print()
	fmt.Println(linkedList.Get(0)) // returns 1
	fmt.Println(linkedList.Get(1)) // returns 3
	fmt.Println(linkedList.Get(2)) // returns -1
	fmt.Println(linkedList.Get(3)) // returns -1
	linkedList.AddAtIndex(3, 2)    // linked list becomes 1->3
	linkedList.Print()
	linkedList.AddAtIndex(2, 2) // linked list becomes 1->3->2
	linkedList.Print()
}

// MyLinkedList linked list
type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	curr := this

	for i := 0; i <= index; i++ {
		if curr.next == nil {
			return -1
		}
		curr = curr.next
	}
	return curr.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.next = &MyLinkedList{val, this.next}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	curr := this
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &MyLinkedList{val, nil}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	curr := this

	for i := 0; i < index; i++ {
		if curr.next == nil {
			return
		}
		curr = curr.next
	}
	curr.next = &MyLinkedList{val, curr.next}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	curr := this

	for i := 0; i < index; i++ {
		if curr.next == nil {
			return
		}
		curr = curr.next
	}
	if curr.next != nil {
		curr.next = curr.next.next
	}

}

func (this *MyLinkedList) Print() {
	curr := this
	for curr != nil {
		fmt.Print(curr.val)
		curr = curr.next
		if curr != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}
