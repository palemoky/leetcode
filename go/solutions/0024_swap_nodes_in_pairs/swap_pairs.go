package swap_pairs

import "leetcode/go/solutions/utils"

// 解法一：迭代解法
func swapPairsIterative(head *utils.ListNode) *utils.ListNode {
	// 由于交换后首节点会发生变化，因此用dummy节点始终指向首节点
	dummy := &utils.ListNode{Next: head}

	// 由于涉及到节点的操作，因此需要用操作节点的上一个节点
	prev := dummy
	// 该遍历条件已经确保节点长度 >= 2
	for prev.Next != nil && prev.Next.Next != nil {
		// 移动并定义需要交换的节点
		left, right := prev.Next, prev.Next.Next

		// 进行交换操作
		next := right.Next
		right.Next = left
		left.Next = next
		prev.Next = right

		// 移动 prev 为下次交换做准备
		// 由于左右两节点已经发生交换，因此 prev 应移动到 left 为下次交换做准备
		prev = left
	}

	return dummy.Next
}

// 解法二：递归解法（更简洁，且无需占用额外空间）
// 最小重复单元：两节点交换
// 递归终止条件：整个链表不再有可交换的节点
// 返回交换后的头部节点
func swapPairsRecursive(head *utils.ListNode) *utils.ListNode {
	// 整个递归的终止条件
	if head == nil || head.Next == nil {
		return head
	}

	// 定义需交换的节点
	left, right := head, head.Next

	// 递归处理链表的节点交换
	// 注意此处传入的 right.Next 即为第三个节点
	left.Next = swapPairsRecursive(right.Next)
	// 交换当前的两个节点
	right.Next = left

	// 递归单元的返回值
	// right 即为新的头结点
	return right
}
