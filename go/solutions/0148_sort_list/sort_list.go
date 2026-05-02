package sort_list

import "leetcode/go/solutions/utils"

// Solution 1: 归并排序
// Time: O(nlogn), Space: O(nlogn)
func sortList(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := split(head)
	left := sortList(head)
	right := sortList(mid)

	return mergeTwoLists(left, right)
}

func split(head *utils.ListNode) *utils.ListNode {
	prev, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		prev, slow, fast = slow, slow.Next, fast.Next.Next
	}

	prev.Next = nil // 注意需要切断链表

	return slow
}

func mergeTwoLists(l1, l2 *utils.ListNode) *utils.ListNode {
	dummy := &utils.ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}

	return dummy.Next
}
