package reverse_linked_list

import "leetcode/go/solutions/utils"

// Time: O(n), Space: O(1)
func reverseList(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}

	var prev *utils.ListNode // 注意不能写 &SinglyNode{}，这样的初始值是0和nil，只有*SinglyNode的值才是nil
	current := head
	for current != nil {
		next := current.Next // 1. 反转节点前先保存指针信息，避免丢失
		current.Next = prev  // 2. 反转节点
		prev = current       // 3. 移动 prev
		current = next       // 4. 移动 current
	}

	return prev
}
