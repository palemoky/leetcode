package palindrome_linked_list

import "leetcode/go/solutions/utils"

// Time: O(n), Space: O(n)
func isPalindromeArray(head *utils.ListNode) bool {
	nodes := []int{}
	for head != nil {
		nodes = append(nodes, head.Val)
		head = head.Next
	}

	left, right := 0, len(nodes)-1
	for left < right {
		if nodes[left] != nodes[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// 优化解法：数组方式需要额外的线性空间，通过查找链表的中间节点和反转后半部分即可判断是否为回文链表
// Time: O(n), Space: O(1)
func isPalindromeTwoPointersAndReverse(head *utils.ListNode) bool {
	// 通过快慢指针获取中间节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	// 反转后半部分链表
	secondHalfHead := reverseList(slow)

	// 比较前后两半链表
	firstHalf, secondHalf := head, secondHalfHead
	isPalindrome := true
	for secondHalf != nil { // 注意此处必须使用 secondHalf 以确保只遍历链表后半部分
		if firstHalf.Val != secondHalf.Val {
			isPalindrome = false
			break
		}

		firstHalf, secondHalf = firstHalf.Next, secondHalf.Next
	}

	// （可选）将链表复原
	reverseList(secondHalfHead)

	return isPalindrome
}

func reverseList(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}

	var prev *utils.ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
