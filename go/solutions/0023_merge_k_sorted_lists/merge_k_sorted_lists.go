package merge_k_sorted_lists

import (
	"container/heap"

	"leetcode/go/solutions/utils"
)

// Solution 1: 分治解法
// Time: O(Nlogk), Space: O(logk)
func mergeKListsByDivideAndConquer(lists []*utils.ListNode) *utils.ListNode {
	if len(lists) == 0 {
		return nil
	}

	return mergeRange(lists, 0, len(lists)-1)
}

func mergeRange(lists []*utils.ListNode, left, right int) *utils.ListNode {
	if left == right {
		return lists[left]
	}

	mid := left + (right-left)/2
	l1 := mergeRange(lists, left, mid)
	l2 := mergeRange(lists, mid+1, right)

	return mergeTwoLists(l1, l2)
}

func mergeTwoLists(l1, l2 *utils.ListNode) *utils.ListNode {
	dummy := &utils.ListNode{}
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}

	return dummy.Next
}

// Solution 2: 优先队列（最小堆）解法
// Time: O(Nlogk), Space: O(k)
func mergeKListsByMinHeap(lists []*utils.ListNode) *utils.ListNode {
	h := &listNodeHeap{}
	heap.Init(h)

	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}

	dummy := &utils.ListNode{}
	cur := dummy

	for h.Len() > 0 {
		node := heap.Pop(h).(*utils.ListNode)
		cur.Next = node
		cur = cur.Next

		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return dummy.Next
}

type listNodeHeap []*utils.ListNode

func (h listNodeHeap) Len() int { return len(h) }

func (h listNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }

func (h listNodeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *listNodeHeap) Push(x any) {
	*h = append(*h, x.(*utils.ListNode))
}

func (h *listNodeHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}
