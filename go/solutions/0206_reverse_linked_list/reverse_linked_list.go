package reverse_linked_list

import "leetcode/go/solutions/utils"

// Time: O(n), Space: O(1)
func reverseList(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}

	var prev *utils.ListNode // 注意不能写 &SinglyNode{}，这样的初始值是 0 和 nil，只有 *SinglyNode 的值才是 nil
	current := head
	for current != nil {
		next := current.Next // 1. 反转节点前先保存指针信息，避免丢失
		current.Next = prev  // 2. 反转节点
		prev = current       // 3. 移动 prev
		current = next       // 4. 移动 current
	}

	// current 始终比 prev 快一步，所以当 current 在 nil 节点时，prev 正好处在最后一个有效节点
	return prev
}
