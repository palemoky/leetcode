package linkedlist

import (
	"fmt"
)

type SinglyListNode struct {
	Value int
	Next  *SinglyListNode
}

func NewLinkedList(arr []int) *SinglyListNode {
	// If the input slice is empty, return nil (empty list)
	if len(arr) == 0 {
		return nil
	}

	// Create the head node with the first element of the slice
	head := &SinglyListNode{Value: arr[0]}
	// Create a movable pointer
	current := head
	// Iterate through the rest of the slice and append nodes to the list
	for _, value := range arr[1:] { // Start from the second element
		current.Next = &SinglyListNode{Value: value} // Create a new node and link it
		current = current.Next                       // Move to the new node
	}

	// Return the head of the constructed linked list
	return head
}

func Append(head *SinglyListNode, value int) *SinglyListNode {
	if head == nil {
		return &SinglyListNode{Value: value}
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &SinglyListNode{Value: value}

	return head
}

func Prepend(head *SinglyListNode, value int) *SinglyListNode {
	return &SinglyListNode{Value: value, Next: head}
}

func Insert(head *SinglyListNode, index, value int) (*SinglyListNode, error) {
	if index < 0 || index > Len(head) {
		return nil, fmt.Errorf("index out of range")
	}

	dummy := &SinglyListNode{Next: head}
	current := dummy
	for range index {
		current = current.Next
	}
	newNode := &SinglyListNode{Value: value, Next: current.Next}
	current.Next = newNode

	return dummy.Next, nil
}

func Delete(head *SinglyListNode, value int) *SinglyListNode {
	dummy := &SinglyListNode{Next: head}

	current := dummy
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}

	return dummy.Next
}

func DeleteAt(head *SinglyListNode, index int) (*SinglyListNode, error) {
	if len := Len(head); index < 0 || index >= len {
		return head, fmt.Errorf("index out of range: index=%d, len=%d", index, len)
	}

	dummy := &SinglyListNode{Next: head}
	current := dummy
	for range index {
		current = current.Next
	}
	current.Next = current.Next.Next

	return dummy.Next, nil
}

func Find(head *SinglyListNode, value int) *SinglyListNode {
	current := head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}

	return nil
}

func Get(head *SinglyListNode, index int) *SinglyListNode {
	if index < 0 || index >= Len(head) {
		return nil
	}

	current := head
	for range index {
		current = current.Next
	}

	return current
}

func Len(head *SinglyListNode) (length int) {
	for head != nil {
		head = head.Next
		length++
	}

	return
}

func ToSlice(head *SinglyListNode) []int {
	s := []int{}
	if head == nil {
		return s
	}

	for head != nil {
		s = append(s, head.Value)
		head = head.Next
	}

	return s
}
