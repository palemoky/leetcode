package intersection_of_two_linked_lists

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

// buildIntersectedLists 构造两条在 sharedVals 处相交的链表。
// 返回 headA、headB 以及相交的起始节点（无相交时为 nil）。
func buildIntersectedLists(aVals, bVals, sharedVals []int) (headA, headB, intersection *utils.ListNode) {
	// 构造共享尾链
	if len(sharedVals) > 0 {
		intersection = utils.NewList(sharedVals)
	}

	// 构造 A 的私有前缀，尾部接上共享链
	if len(aVals) > 0 {
		headA = utils.NewList(aVals)
		cur := headA
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = intersection
	} else {
		headA = intersection
	}

	// 构造 B 的私有前缀，尾部接上共享链
	if len(bVals) > 0 {
		headB = utils.NewList(bVals)
		cur := headB
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = intersection
	} else {
		headB = intersection
	}

	return headA, headB, intersection
}

func TestGetIntersectionNode(t *testing.T) {
	t.Parallel()

	// 构造各测试用例的链表（需在用例外构造，保证指针共享）
	headA1, headB1, inter1 := buildIntersectedLists([]int{4, 1}, []int{5, 6, 1}, []int{8, 4, 5})
	headA2, headB2, inter2 := buildIntersectedLists([]int{1, 9, 1}, []int{3}, []int{2, 4})
	headA3, headB3, _ := buildIntersectedLists([]int{2, 6, 4}, []int{1, 5}, nil)
	headA4, headB4, inter4 := buildIntersectedLists(nil, nil, []int{1})

	testCases := []struct {
		name     string
		headA    *utils.ListNode
		headB    *utils.ListNode
		expected *utils.ListNode
	}{
		{
			name:     "both nil",
			headA:    nil,
			headB:    nil,
			expected: nil,
		},
		{
			name:     "headA nil",
			headA:    nil,
			headB:    utils.NewList([]int{1, 2, 3}),
			expected: nil,
		},
		{
			name:     "headB nil",
			headA:    utils.NewList([]int{1, 2, 3}),
			headB:    nil,
			expected: nil,
		},
		{
			name:     "no intersection",
			headA:    headA3,
			headB:    headB3,
			expected: nil,
		},
		{
			name:     "intersect after different length prefixes (LeetCode example 1)",
			headA:    headA1,
			headB:    headB1,
			expected: inter1,
		},
		{
			name:     "intersect after different length prefixes (LeetCode example 2)",
			headA:    headA2,
			headB:    headB2,
			expected: inter2,
		},
		{
			name:     "both point to same single node",
			headA:    headA4,
			headB:    headB4,
			expected: inter4,
		},
	}

	funcsToTest := map[string]func(*utils.ListNode, *utils.ListNode) *utils.ListNode{
		"HashMap":    getIntersectionNodeHashMap,
		"TwoPointer": getIntersectionNodeTwoPointer,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					actual := fn(tc.headA, tc.headB)
					assert.Equal(t, tc.expected, actual)
				})
			}
		})
	}
}
