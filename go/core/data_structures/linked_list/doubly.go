package linkedlist

import (
	"fmt"
)

// DoublyNode represents a node in a doubly linked list.
type DoublyNode struct {
	Value      int
	Prev, Next *DoublyNode
}

// DoublyList represents a doubly linked list.
type DoublyList struct {
	Head, Tail *DoublyNode
	Len        int
}

// NewDoublyList creates and returns a new DoublyList initialized with values from an array.
func NewDoublyList(array []int) *DoublyList {
	list := &DoublyList{}

	if len(array) == 0 {
		return list
	}

	headNode := &DoublyNode{Value: array[0], Prev: nil}
	list.Head = headNode
	list.Tail = headNode
	list.Len = 1

	current := headNode
	for i := 1; i < len(array); i++ {
		newNode := &DoublyNode{Value: array[i], Prev: current}
		current.Next = newNode
		current = current.Next
		list.Tail = current
		list.Len++
	}

	return list
}

// DoublyAppend adds a new node with the given value to the end of the list.
func (list *DoublyList) DoublyAppend(value int) {
	// Create node and update the length
	newNode := &DoublyNode{Value: value}

	// Handle edge cases
	if list.Len == 0 {
		list.Head = newNode
		list.Tail = newNode
	} else {
		// Update pointer
		list.Tail.Next = newNode
		newNode.Prev = list.Tail
		list.Tail = newNode
	}

	list.Len++
}

func (list *DoublyList) DoublyPrepend(value int) {
	// Create Node and update the length
	newNode := &DoublyNode{Value: value}

	// Handle edge cases
	if list.Len == 0 {
		list.Head = newNode
		list.Tail = newNode
	} else {
		// Update pointer
		newNode.Next = list.Head
		list.Head.Prev = newNode
		list.Head = newNode
	}

	list.Len++
}

func (list *DoublyList) DoublyInsert(value, index int) {
	// Index out-of-bounds check
	if index < 0 || index > list.Len {
		panic(fmt.Sprintf("index out of range: index %d, len %d", index, list.Len))
	}

	switch index {
	case 0: // Insert at the head
		list.DoublyPrepend(value)
	case list.Len: // Insert at the tail; 注意可插入位置有len+1，尾部是len，而非len-1
		list.DoublyAppend(value)
	default:
		// Find the previous node of the insertion position
		current := list.DoublyGet(index - 1)

		// Update pointer
		newNode := &DoublyNode{Value: value}
		newNode.Prev = current
		newNode.Next = current.Next
		current.Next.Prev = newNode
		current.Next = newNode

		// Update the length
		list.Len++
	}
}

func (list *DoublyList) DoublyDelete(value int) {
	current := list.Head
	for current != nil {
		if current.Value != value {
			current = current.Next
			continue
		}

		next := current.Next
		switch current {
		case list.Head: // delete head
			list.Head = next
			if list.Head != nil {
				list.Head.Prev = nil
			} else {
				// 如果删除后链表为空
				list.Tail = nil
			}
		case list.Tail: // delete tail
			list.Tail = current.Prev
			list.Tail.Next = nil
		default: // delete middle
			current.Prev.Next = current.Next
			current.Next.Prev = current.Prev
		}

		list.Len--
		current = next // move to next node
	}
}

func (list *DoublyList) DoublyDeleteAt(index int) {
	if index < 0 || index >= list.Len {
		return
	}

	if list.Len == 1 {
		list.Head = nil
		list.Tail = nil
		list.Len--
		return
	}

	switch index {
	case 0: // delete head
		list.Head = list.Head.Next
		list.Head.Prev = nil
	case list.Len - 1: // delete tail
		list.Tail = list.Tail.Prev
		list.Tail.Next = nil
	default:
		// find delete position
		current := list.DoublyGet(index)

		// update pointer
		current.Prev.Next = current.Next
		current.Next.Prev = current.Prev
	}
	list.Len--
}

func (list *DoublyList) DoublyFind(value int) *DoublyNode {
	current := list.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}

	return nil
}

func (list *DoublyList) DoublyGet(index int) (current *DoublyNode) {
	// index out of bound check
	if index < 0 || index >= list.Len {
		return nil
	}

	// Decide traversal direction based on the index
	if index < list.Len/2 {
		current = list.Head
		for range index {
			current = current.Next
		}
	} else {
		current = list.Tail
		for i := list.Len - 1; i > index; i-- {
			current = current.Prev
		}
	}

	return current
}
