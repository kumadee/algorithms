package linkedlist

import (
	"fmt"
)

// Element is a member of linked list
type Element struct {
	Value string
	Next  *Element
}

// PrintLinkedList prints a linked list
func (head *Element) PrintLinkedList() {
	for {
		fmt.Printf("|%s|-->", head.Value)
		if head.Next == nil {
			fmt.Printf("END\n")
			return
		}
		head = head.Next
	}
}
