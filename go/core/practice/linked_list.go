package practice

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表
func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev, next *ListNode // 注意不能写 &SinglyNode{}，这样的初始值是0和nil，只有*SinglyNode的值才是nil
	current := head
	for current != nil {
		next = current.Next // 1. save next node
		current.Next = prev // 2. reverse
		prev = current      // 3. move prev
		current = next      // 4. move current
	}

	return prev
}

// 通过数组的连续性查找中间节点
func findMiddleArray(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	nodes := []*ListNode{}
	current := head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	return nodes[len(nodes)/2]
}

// 通过快慢指针查找中间节点
func findMiddleTwoPointers(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	// 只需校验快指针的有效性即可
	for fast != nil && fast.Next != nil { // 因为fast要跨两步，所以要确定两步都非空
		slow, fast = slow.Next, fast.Next.Next
	}

	return slow
}

func getListLen(head *ListNode) int {
	len := 0
	for head != nil {
		len++
		head = head.Next
	}

	return len
}

// 合并有序链表
func mergeSortedTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	// 注意该遍历需同时操作3个链表
	for l1 != nil && l2 != nil {
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

	return dummy.Next
}
