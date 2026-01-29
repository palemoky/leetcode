package reverse_linked_list_ii

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func TestReverseBetween(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		head     []int
		left     int
		right    int
		expected []int
	}{
		{
			name:     "Example 1: 反转中间部分",
			head:     []int{1, 2, 3, 4, 5},
			left:     2,
			right:    4,
			expected: []int{1, 4, 3, 2, 5},
		},
		{
			name:     "Example 2: 单节点",
			head:     []int{5},
			left:     1,
			right:    1,
			expected: []int{5},
		},
		{
			name:     "反转整个链表",
			head:     []int{1, 2, 3, 4, 5},
			left:     1,
			right:    5,
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "反转前两个节点",
			head:     []int{1, 2, 3},
			left:     1,
			right:    2,
			expected: []int{2, 1, 3},
		},
		{
			name:     "反转后两个节点",
			head:     []int{1, 2, 3},
			left:     2,
			right:    3,
			expected: []int{1, 3, 2},
		},
		{
			name:     "两个节点的链表",
			head:     []int{3, 5},
			left:     1,
			right:    2,
			expected: []int{5, 3},
		},
		{
			name:     "反转相邻两个节点",
			head:     []int{1, 2, 3, 4, 5},
			left:     3,
			right:    4,
			expected: []int{1, 2, 4, 3, 5},
		},
	}

	funcsToTest := map[string]func(*utils.ListNode, int, int) *utils.ListNode{
		"reverseBetweenByThreeSteps": reverseBetweenByThreeSteps,
		"reverseBetweenByHeadInsert": reverseBetweenByHeadInsert,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					head := utils.NewList(tc.head)
					result := fn(head, tc.left, tc.right)
					resultArray := utils.ToSlice(result)
					assert.Equal(t, tc.expected, resultArray)
				})
			}
		})
	}
}
