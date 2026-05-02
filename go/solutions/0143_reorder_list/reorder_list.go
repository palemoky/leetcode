package reorder_list

import "leetcode/go/solutions/utils"

// Solution 1: 反转后半段
// Time: O(nlogn), Space: O()
func reorderList(head *utils.ListNode) {
	mid := middleNode(head)
	reversedHead := reverseList(mid)
	p1, p2 := head, reversedHead
	for p2.Next != nil {
		p1Next, p2Next := p1.Next, p2.Next // 保存指针避免断链
		p1.Next, p2.Next = p2, p1Next      // 交叉连接
		p1, p2 = p1Next, p2Next            // 移动节点
	}
}

func middleNode(head *utils.ListNode) *utils.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	return slow
}

func reverseList(head *utils.ListNode) *utils.ListNode {
	var prev *utils.ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
