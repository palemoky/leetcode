package find_linked_list_cycle_entry

import "leetcode/go/solutions/utils"

// 找到环的入口节点
// Time: O(n), Space: O(n)
func findCycleEnteryHashMap(head *utils.ListNode) *utils.ListNode {
	scanned := map[*utils.ListNode]struct{}{}
	for head != nil {
		if _, ok := scanned[head]; ok {
			return head
		}
		scanned[head] = struct{}{}
		head = head.Next
	}

	return nil
}

// 数学分析规律，再以双指针求解
// Time: O(n), Space: O(1)
func findCycleEnteryMathTwoPointers(head *utils.ListNode) *utils.ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		// 先判断是否有环
		if fast == slow {
			headPtr, meetPtr := head, fast
			// 再根据数学推导出的从head到入口点的步数与相遇点到入口点的步数相同获取入口点位置
			// 此时新建两个指针，分别从head和相遇点以相同的速度移动，当二者相遇时，即为入口点
			for headPtr != meetPtr {
				headPtr, meetPtr = headPtr.Next, meetPtr.Next
			}
			return headPtr // 返回meetPtr也可以
		}
	}

	return nil
}
