package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	dummy := &ListNode{}
	current := dummy
	for _, val := range vals {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return dummy.Next
}

// toSlice 将链表转换回整数切片，方便断言比较
func ToSlice(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	var res []int
	current := head
	for current != nil {
		res = append(res, current.Val)
		current = current.Next
	}
	return res
}
