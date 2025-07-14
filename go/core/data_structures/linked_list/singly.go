package linkedlist

import (
	"fmt"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

func NewLinkedList(arr []int) *ListNode {
	// If the input slice is empty, return nil (empty list)
	if len(arr) == 0 {
		return nil
	}

	// Create the head node with the first element of the slice
	head := &ListNode{Value: arr[0]}
	// Create a movable pointer
	current := head
	// Iterate through the rest of the slice and append nodes to the list
	for _, value := range arr[1:] { // Start from the second element
		current.Next = &ListNode{Value: value} // Create a new node and link it
		current = current.Next                 // Move to the new node
	}

	// Return the head of the constructed linked list
	return head
}

func Append(head *ListNode, value int) *ListNode {
	if head == nil {
		return &ListNode{Value: value}
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &ListNode{Value: value}

	return head
}

func Prepend(head *ListNode, value int) *ListNode {
	return &ListNode{Value: value, Next: head}
}

func Insert(head *ListNode, index, value int) (*ListNode, error) {
	if index < 0 || index > Len(head) {
		return nil, fmt.Errorf("index out of range")
	}

	dummy := &ListNode{Next: head}
	current := dummy
	for range index {
		current = current.Next
	}
	newNode := &ListNode{Value: value, Next: current.Next}
	current.Next = newNode

	return dummy.Next, nil
}

func Delete(head *ListNode, value int) *ListNode {
	dummy := &ListNode{Next: head}

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

func DeleteAt(head *ListNode, index int) (*ListNode, error) {
	if len := Len(head); index < 0 || index >= len {
		return head, fmt.Errorf("index out of range: index=%d, len=%d", index, len)
	}

	dummy := &ListNode{Next: head}
	current := dummy
	for range index {
		current = current.Next
	}
	current.Next = current.Next.Next

	return dummy.Next, nil
}

func Find(head *ListNode, value int) *ListNode {
	current := head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}

	return nil
}

func Get(head *ListNode, index int) *ListNode {
	if index < 0 || index >= Len(head) {
		return nil
	}

	current := head
	for range index {
		current = current.Next
	}

	return current
}

func Len(head *ListNode) (length int) {
	for head != nil {
		head = head.Next
		length++
	}

	return
}

func ToSlice(head *ListNode) []int {
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
