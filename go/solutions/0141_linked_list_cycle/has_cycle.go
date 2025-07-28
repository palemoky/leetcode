package has_cycle

import "leetcode/go/solutions/utils"

// 解法一：哈希表
// 优点：简单直观；缺点：使用额外空间
// Time: O(n), Space: O(n)
func hasCycleHashMap(head *utils.ListNode) bool {
	m := make(map[*utils.ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}

	return false
}

// 解法二：快慢指针
// 优点：无需使用额外空间；缺点：实现略微复杂
// 该解法的核心在于环形链表会导致快慢指针相遇
// Time: O(n), Space: O(1)
func hasCycleTwoPoints(head *utils.ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		// 快指针以两倍速移动，一定会与慢指针相遇，为什么呢？
		// 把两个指针的移动看作相对运动，那么就是快指针在每次一个节点的速度靠近慢指针，所以快慢指针必然相遇
		// 如果快指针每次移动3个节点，则相对运动下，就可能跳过慢指针而不相遇
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}
