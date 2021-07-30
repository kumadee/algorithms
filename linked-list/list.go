package linkedlist

import (
    "fmt"
)

// Element is a member of linked list
type Element struct {
    value int
    next *Element
}

// PrintLinkedList prints a linked list
func (e Element) PrintLinkedList() {
    for {
        if (e.next == nil) {
            fmt.Printf("END\n")
            return
        }
        fmt.Printf("|%d|-->", e.value)
    }
}

func main() {

}
