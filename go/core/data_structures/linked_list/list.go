package linkedlist

type Node interface {
	GetValue() int
	GetNext() Node
}

type List interface {
	GetHead() Node
}

// Methods to make SinglyNode and SinglyList satisfy the interfaces
func (n *SinglyNode) GetValue() int { return n.Value }
func (n *SinglyNode) GetNext() Node {
	if n.Next == nil {
		return nil
	}
	return n.Next
}
func (l *SinglyList) GetHead() Node {
	if l.Head == nil {
		return nil
	}
	return l.Head
}

// Methods to make DoublyNode and DoublyList satisfy the interfaces
func (n *DoublyNode) GetValue() int { return n.Value }
func (n *DoublyNode) GetNext() Node {
	if n.Next == nil {
		return nil
	}
	return n.Next
}
func (l *DoublyList) GetHead() Node {
	if l.Head == nil {
		return nil
	}
	return l.Head
}

// toSlice is a helper function to convert a DoublyList to a slice for easy comparison.
// It traverses the list from Head to Tail.
func toSlice(list List) []int {
	head := list.GetHead()
	if head == nil {
		return []int{}
	}

	slice := make([]int, 0)
	current := head
	for current != nil {
		slice = append(slice, current.GetValue())
		current = current.GetNext()
	}
	return slice
}

// toSliceReverse is a helper function to test the integrity of Prev pointers.
// It traverses the list from Tail to Head.
func toSliceReverse(list *DoublyList) []int {
	if list == nil || list.Tail == nil {
		return []int{}
	}
	slice := make([]int, 0, list.Len)
	current := list.Tail
	for current != nil {
		slice = append(slice, current.Value)
		current = current.Prev
	}
	return slice
}

// reverseSlice is a simple helper to reverse a slice for comparison.
func reverseSlice(s []int) []int {
	// Create a new slice to avoid modifying the original
	reversed := make([]int, len(s))
	for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
		reversed[i] = s[j]
	}
	return reversed
}
