package reverse_nodes_in_k_group

import "leetcode/go/solutions/utils"

// 解法一：头插法（迭代） ⭐ 推荐
// Time: O(n), Space: O(1)
// 核心思想：将92题的头插法应用多次，每次反转 k 个节点

// nil -> 1 -> 2 -> 3 -> 4 -> 5 (k=3)
//  ↑     ↑    ↑    ↑    ↑
// prev curr  next tail nextGroup

func reverseKGroup(head *utils.ListNode, k int) *utils.ListNode {
	dummy := &utils.ListNode{Next: head}
	prev := dummy

	for {
		// 1. 检查剩余节点是否够 k 个
		tail := prev
		for range k {
			tail = tail.Next
			if tail == nil {
				return dummy.Next // 不足 k 个时结束反转
			}
		}

		// 2. 保存下一组的起始位置
		nextGroup := tail.Next

		// 3. 反转这 k 个节点（使用头插法，类似92题）
		curr := prev.Next
		for range k - 1 {
			next := curr.Next
			curr.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}

		// 4. 移动到下一组
		// 反转后，原来的第一个节点变成了最后一个节点
		curr.Next = nextGroup
		prev = curr
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
	prev := reverseKGroupByRecursion(dummy, k)

	// 3. 反转当前 k 个节点
	// 利用 Go 的多重赋值，一行完成三个操作：
	// - next := head.Next  (保存下一个节点)
	// - head.Next = prev    (反转指针)
	// - prev = head         (移动 prev)
	// - head = next         (移动 head)
	for ; cnt > 0; cnt-- {
		head.Next, prev, head = prev, head, head.Next
	}

	return prev
}
