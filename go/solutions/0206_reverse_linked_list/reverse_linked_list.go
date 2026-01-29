package reverse_linked_list

import "leetcode/go/solutions/utils"

// Time: O(n), Space: O(1)
func reverseList(head *utils.ListNode) *utils.ListNode {
	// 注意不能写 &SinglyNode{}，这样的初始值是 0 和 nil，只有 *SinglyNode 的值才是 nil
	var prev *utils.ListNode
	// 反转链表类似交换变量：先暂存被修改的值，再修改指针
	// 1→2→3 reverse 1←2→3
	// ↑             ↑ ↑
	// head       prev head
	for head != nil {
		next := head.Next // 1. 反转节点前先保存指针信息，避免丢失
		head.Next = prev  // 2. 反转节点
		prev = head       // 3. 移动 prev
		head = next       // 4. 移动 head
	}

	// head 始终比 prev 快一步，所以当 head 在 nil 节点时，prev 正好处在最后一个有效节点
	return prev
}
