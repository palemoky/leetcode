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

// 两两交换链表的节点（迭代解法）
func swapPairsIterative(head *ListNode) *ListNode {
	// 由于交换后首节点会发生变化，因此用dummy节点始终指向首节点
	dummy := &ListNode{Next: head}

	// 由于涉及到节点的操作，因此需要用操作节点的上一个节点
	prev := dummy
	// 该遍历条件已经确保节点长度 >= 2
	for prev.Next != nil && prev.Next.Next != nil {
		// 移动并定义需要交换的节点
		left, right := prev.Next, prev.Next.Next

		// 进行交换操作
		next := right.Next
		right.Next = left
		left.Next = next
		prev.Next = right

		// 移动 prev 为下次交换做准备
		// 由于左右两节点已经发生交换，因此 prev 应移动到 left 为下次交换做准备
		prev = left
	}

	return dummy.Next
}

// 两两交换链表的节点（递归解法）
// 最小重复单元：两节点交换
// 递归终止条件：整个链表不再有可交换的节点
// 返回交换后的头部节点
func swapPairsRecursive(head *ListNode) *ListNode {
	// 整个递归的终止条件
	if head == nil || head.Next == nil {
		return head
	}

	// 定义需交换的节点
	left, right := head, head.Next

	// 递归处理链表的节点交换
	// 注意此处传入的 right.Next 即为第三个节点
	left.Next = swapPairsRecursive(right.Next)
	// 交换当前的两个节点
	right.Next = left

	// 递归单元的返回值
	// right 即为新的头结点
	return right
}

// 部分反转链表

// 重排链表

// 两个链表的第一个公共节点

// 合并 K 个有序链表

// 回文链表

// 链表的两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}
