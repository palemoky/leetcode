package practice

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表
// Time: O(n), Space: O(1)
func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev, next *ListNode // 注意不能写 &SinglyNode{}，这样的初始值是0和nil，只有*SinglyNode的值才是nil
	current := head
	for current != nil {
		next = current.Next // 1. 反转节点前先保存指针信息，避免丢失
		current.Next = prev // 2. 反转节点
		prev = current      // 3. 移动 prev
		current = next      // 4. 移动 current
	}

	return prev
}

// 通过数组的连续性查找中间节点
// Time: O(n), Space: O(n)
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
// Time: O(n), Space: O(1)
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

// 合并有序链表
// Time: O(n), Space: O(n)
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

// 是否有环
func hasCycleHashMap(head *ListNode) bool {
	scanned := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := scanned[head]; ok {
			return true
		}
		scanned[head] = struct{}{}
		head = head.Next
	}

	return false
}

// 该解法的核心在于环形链表会导致快慢指针相遇
func hasCycleTwoPointers(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		// 快指针以两倍速移动，一定会与慢指针相遇，为什么呢？
		// 把两个指针的移动看作相对运动，那么就是快指针在每次一个节点的速度靠近慢指针，所以快慢指针必然相遇
		// 如果快指针每次移动3个节点，则相对运动下，就可能跳过慢指针而不相遇
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}

// 找到环的入口节点
// Time: O(n), Space: O(n)
func findCycleEnteryHashMap(head *ListNode) *ListNode {
	scanned := map[*ListNode]struct{}{}
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
func findCycleEnteryMathTwoPointers(head *ListNode) *ListNode {
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

// 删除链表的倒数第N个节点
func removeNthFromEndTwoPointers(head *ListNode, n int) *ListNode {
	// 1. 可能删除任意位置的节点，因此需将虚拟节点指向头结点
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy

	// 2. 让快指针移动n位
	// 边界条件：~~空链表、n > len(list)、~~一个节点的链表、删除头结点、删除尾节点
	// 题目已经限制节点数量为[1, 30]，n的范围[1, len(list)]
	for range n {
		fast = fast.Next
	}

	// 3. 同时移动快慢指针，直至快指针到达节点尾部
	for fast.Next != nil {
		fast, slow = fast.Next, slow.Next
	}

	// 4. 删除slow所在的节点
	// A->B->C->D to A->B->D
	slow.Next = slow.Next.Next

	// 返回头结点
	return dummy.Next
}
