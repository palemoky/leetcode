package practice

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表
func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev, next *ListNode
	current := head
	for current != nil {
		// save next node
		next = current.Next
		// reverse
		current.Next = prev
		// move prev
		prev = current
		// move current
		current = next
	}

	return prev
}

// 通过数组的连续性查找中间节点
func findMiddleArray(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	nodes := []*ListNode{}
	current := head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	return nodes[len(nodes)/2]
}

// 通过快慢指针查找中间节点
func findMiddleTwoPointers(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	// 只需校验快指针的有效性即可
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	return slow
}
