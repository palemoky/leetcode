package reverse_nodes_in_k_group

import "leetcode/go/solutions/utils"

// 解法一：头插法（迭代） ⭐ 推荐
// Time: O(n), Space: O(1)
// 核心思想：将92题的头插法应用多次，每次反转 k 个节点

// nil -> 1 -> 2 -> 3 -> 4 -> 5 (k=3)
//  ↑     ↑    ↑    ↑    ↑
// pre   cur  next tail nextGroup

func reverseKGroup(head *utils.ListNode, k int) *utils.ListNode {
	dummy := &utils.ListNode{Next: head}
	prev := dummy

	for {
		// 1. 检查剩余节点是否够 k 个
		tail := prev
		for range k {
			tail = tail.Next
			if tail == nil {
				return dummy.Next // 不足 k 个，直接返回
			}
		}

		// 2. 保存下一组的起始位置
		nextGroup := tail.Next

		// 3. 反转这 k 个节点（使用头插法，类似92题）
		cur := prev.Next
		for range k - 1 {
			next := cur.Next
			cur.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}

		// 4. 移动到下一组
		// 反转后，原来的第一个节点变成了最后一个节点
		cur.Next = nextGroup
		prev = cur
	}
}

// 解法二：递归
// Time: O(n), Space: O(n/k) - 递归栈深度
// 核心思想：先递归处理后续部分，再反转当前 k 个节点
func reverseKGroupByRecursion(head *utils.ListNode, k int) *utils.ListNode {
	// 1. 检查是否有 k 个节点
	dummy, cnt := head, 0
	for ; cnt < k; cnt++ {
		if dummy == nil {
			return head // 不足 k 个，直接返回
		}
		dummy = dummy.Next
	}

	// 2. 递归处理后续部分
	pre := reverseKGroupByRecursion(dummy, k)

	// 3. 反转当前 k 个节点
	// 利用 Go 的多重赋值，一行完成三个操作：
	// - next := head.Next  (保存下一个节点)
	// - head.Next = pre    (反转指针)
	// - pre = head         (移动 pre)
	// - head = next        (移动 head)
	for ; cnt > 0; cnt-- {
		head.Next, pre, head = pre, head, head.Next
	}

	return pre
}
