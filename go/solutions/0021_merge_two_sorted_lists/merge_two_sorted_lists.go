package merge_two_sorted_lists

import "leetcode/go/solutions/utils"

// 解题思路：创建一个dummy作为新的链表，再将list1和list2中较小的值不断挂载到dummy的链表上。注意处理list1或list2剩余部分
// Time: O(n), Space: O(n)
func mergeSortedTwoLists(l1, l2 *utils.ListNode) *utils.ListNode {
	dummy := &utils.ListNode{}
	current := dummy

	// 注意该遍历需同时操作3个链表
	for l1 != nil && l2 != nil { // 注意是 &&
		if l1.Val < l2.Val { // 将较小的值挂载在新链表上
			current.Next = l1
			l1 = l1.Next // 移动原链表
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next // 移动新链表
	}

	// 将链表剩余部分挂载，同时处理原链表为空
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	// 注意返回的是 dummy.Next
	return dummy.Next
}
