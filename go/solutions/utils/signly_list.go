package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	dummy := &ListNode{}
	current := dummy
	for _, val := range vals {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return dummy.Next
}

// NewCycleList 从一个整数切片创建链表，并能指定环的入口
// vals: 链表节点的值
// cyclePos: 尾节点指向的节点的索引。如果为 -1，则不创建环。
func NewCycleList(vals []int, cyclePos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	// 1. 创建所有节点，并用一个切片存起来，方便后续通过索引访问
	nodes := make([]*ListNode, len(vals))
	for i, val := range vals {
		nodes[i] = &ListNode{Val: val}
	}

	// 2. 将节点线性连接起来
	for i := 0; i < len(vals)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	// 3. 如果需要，创建环
	if cyclePos != -1 && cyclePos < len(vals) {
		tail := nodes[len(vals)-1]
		tail.Next = nodes[cyclePos] // 尾节点指向环的入口
	}

	return nodes[0] // 返回头节点
}

// ToSlice 将链表转换回整数切片，方便断言比较
func ToSlice(head *ListNode) []int {
	s := []int{}
	for current := head; current != nil; current = current.Next {
		s = append(s, current.Val)
	}
	return s
}
