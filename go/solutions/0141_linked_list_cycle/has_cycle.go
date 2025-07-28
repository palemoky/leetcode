package has_cycle

import "leetcode/go/solutions/utils"

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

func hasCycleTwoPoints(head *utils.ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}
