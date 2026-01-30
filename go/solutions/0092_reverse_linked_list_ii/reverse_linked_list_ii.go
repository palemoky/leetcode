package reverse_linked_list_ii

import "leetcode/go/solutions/utils"

// 解法一：头插法 ⭐ 推荐
// Time: O(n), Space: O(1)
// 原地操作，一次遍历完成，适合局部反转
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
		next := cur.Next      // 保存指针
		cur.Next = next.Next  // 修改指针
		next.Next = prev.Next //
		prev.Next = next
	}

	return dummy.Next
}

// 解法二：迭代法
// Time: O(n), Space: O(1)
// 思路更直观：断开-反转-连接，适合学习理解
func reverseBetweenByIteration(head *utils.ListNode, left int, right int) *utils.ListNode {
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
	// 使用迭代法反转链表（类似 206 题）
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

// 两种方法对比：
//
// 迭代法：像翻书，一页一页地翻过来
// 1→2→3  =>  1←2 3  =>  1←2←3
// 适合：全部反转（206. 反转链表）
//
// 头插法：像抽扑克牌，依次把牌插到头部（prev后面）
// prev→1→2→3→4
//      ↑
//     cur
// prev→2→1→3→4 (把2插到prev后) => prev→3→2→1→4 (把3插到prev后)
// 适合：局部反转（92. 反转链表 II）
