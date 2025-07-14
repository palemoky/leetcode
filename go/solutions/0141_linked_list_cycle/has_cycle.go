package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycleHashMap(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}

	return false
}

func hasCycleTwoPoints(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}
