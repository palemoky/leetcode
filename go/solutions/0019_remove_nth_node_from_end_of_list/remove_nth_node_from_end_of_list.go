package remove_nth_node_from_end_of_list

import "leetcode/go/solutions/utils"

// 删除链表的倒数第N个节点
func removeNthFromEndTwoPointers(head *utils.ListNode, n int) *utils.ListNode {
	// 1. 可能删除任意位置的节点，因此需将虚拟节点指向头结点
	dummy := &utils.ListNode{Next: head}
	fast, slow := dummy, dummy

	// 2. 让快指针移动n位
	// 边界条件：~~空链表、n > len(list)、~~一个节点的链表、删除头结点、删除尾节点
	// 题目已经限制节点数量为[1, 30]，n的范围[1, len(list)]
	for range n {
		fast = fast.Next
	}

	// 3. 同时移动快慢指针，直至快指针到达节点尾部
	for fast.Next != nil {
		fast, slow = fast.Next, slow.Next
	}

	// 4. 删除slow所在的节点
	// A->B->C->D to A->B->D
	slow.Next = slow.Next.Next

	// 返回头结点
	return dummy.Next
}
