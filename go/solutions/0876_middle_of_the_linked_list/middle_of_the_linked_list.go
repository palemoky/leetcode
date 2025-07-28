package middle_of_the_linked_list

import "leetcode/go/solutions/utils"

// 通过数组的连续性查找中间节点
// Time: O(n), Space: O(n)
func findMiddleArray(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}

	nodes := []*utils.ListNode{}
	current := head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	return nodes[len(nodes)/2]
}

// 通过快慢指针查找中间节点
// Time: O(n), Space: O(1)
func findMiddleTwoPointers(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	// 只需校验快指针的有效性即可
	for fast != nil && fast.Next != nil { // 因为fast要跨两步，所以要确定两步都非空
		slow, fast = slow.Next, fast.Next.Next
	}

	return slow
}
