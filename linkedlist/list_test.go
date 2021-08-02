package linkedlist

import (
	"testing"
)

func TestPrintLinkedList(t *testing.T) {
	head := new(Element)
	head.Value = "HEAD"
	last := head
	elements := []string{"Element1", "Element2"}
	for _, v := range elements {
		last.Next = &Element{Value: v, Next: nil}
		last = last.Next
	}
	head.PrintLinkedList()
	t.Logf("Pass")
}
