package intersection_of_two_linked_lists

import "leetcode/go/solutions/utils"

// Solution 1: 哈希表扫描
// Time: O(n), Space: O(n)
func getIntersectionNodeHashMap(headA, headB *utils.ListNode) *utils.ListNode {
	seen := map[*utils.ListNode]struct{}{}
	for headA != nil {
		seen[headA] = struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := seen[headB]; ok {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

// Solution 2: 双指针
// Time: O(n), Space: O(1)
func getIntersectionNodeTwoPointer(headA, headB *utils.ListNode) *utils.ListNode {
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return pA // 没有相交时，遍历结束的链表最终指向 nil
}
