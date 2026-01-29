package reverse_linked_list_ii

import "leetcode/go/solutions/utils"

// 解法一：三步法（断开-反转-连接）
// Time: O(n), Space: O(1)
// 思路更直观：先断开区间，反转区间，再连接回去
func reverseBetweenByThreeSteps(head *utils.ListNode, left int, right int) *utils.ListNode {
	dummy := &utils.ListNode{Next: head}

	// 步骤1: 定位关键节点
	// prev: 反转区间的前驱节点（第 left-1 个）
	// leftNode: 反转区间的第一个节点（第 left 个）
	// rightNode: 反转区间的最后一个节点（第 right 个）
	// succ: 反转区间的后继节点（第 right+1 个）

	prev := dummy
	for range left - 1 {
		prev = prev.Next
	}

	leftNode := prev.Next
	rightNode := leftNode
	for range right - left {
		rightNode = rightNode.Next
	}

	succ := rightNode.Next

	// 步骤2: 断开区间
	prev.Next = nil
	rightNode.Next = nil

	// 步骤3: 反转区间 [leftNode, rightNode]
	reverseList := func(head *utils.ListNode) *utils.ListNode {
		var prev *utils.ListNode
		for head != nil {
			next := head.Next
			head.Next = prev
			prev = head
			head = next
		}
		return prev
	}

	// 步骤4: 重新连接
	// prev -> (反转后的区间) -> succ
	// 反转后：rightNode 变成新头，leftNode 变成新尾
	prev.Next = reverseList(leftNode)
	leftNode.Next = succ

	return dummy.Next
}

// 解法二：穿针引线法（头插法）
// Time: O(n), Space: O(1)
func reverseBetweenByHeadInsert(head *utils.ListNode, left int, right int) *utils.ListNode {
	// 使用 dummy 节点简化边界处理（当 left == 1 时）
	dummy := &utils.ListNode{Next: head}

	// 1. 找到反转区间的前驱节点（第 left-1 个节点）
	prev := dummy
	for range left - 1 {
		prev = prev.Next
	}

	// 2. 反转区间 [left, right]
	// prev -> 1 -> 2 -> 3 -> 4 -> 5
	//         ↑         ↑
	//       left      right
	cur := prev.Next
	for range right - left {
		// 头插法：每次将 cur.Next 移到 prev 后面
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}
