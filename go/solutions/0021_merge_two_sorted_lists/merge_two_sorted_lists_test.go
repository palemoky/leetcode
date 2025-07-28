package merge_two_sorted_lists

import (
	"leetcode/go/solutions/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSortedTwoLists(t *testing.T) {
	// 定义测试用例表，覆盖各种情况
	testCases := []struct {
		name     string
		l1       *utils.ListNode
		l2       *utils.ListNode
		expected *utils.ListNode
	}{
		{
			name:     "both lists are nil",
			l1:       nil,
			l2:       nil,
			expected: nil,
		},
		{
			name:     "l1 is nil, l2 is not",
			l1:       nil,
			l2:       utils.NewList([]int{1, 2, 3}),
			expected: utils.NewList([]int{1, 2, 3}),
		},
		{
			name:     "l2 is nil, l1 is not",
			l1:       utils.NewList([]int{4, 5, 6}),
			l2:       nil,
			expected: utils.NewList([]int{4, 5, 6}),
		},
		{
			name:     "standard merge with interleaved values",
			l1:       utils.NewList([]int{1, 3, 5}),
			l2:       utils.NewList([]int{2, 4, 6}),
			expected: utils.NewList([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name:     "merge with different lengths (l1 shorter)",
			l1:       utils.NewList([]int{1, 5}),
			l2:       utils.NewList([]int{2, 3, 4}),
			expected: utils.NewList([]int{1, 2, 3, 4, 5}),
		},
		{
			name:     "merge with different lengths (l2 shorter)",
			l1:       utils.NewList([]int{1, 2, 10}),
			l2:       utils.NewList([]int{5, 6}),
			expected: utils.NewList([]int{1, 2, 5, 6, 10}),
		},
		{
			name:     "l1 values are all smaller than l2",
			l1:       utils.NewList([]int{1, 2, 3}),
			l2:       utils.NewList([]int{4, 5, 6}),
			expected: utils.NewList([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name:     "merge with duplicate values",
			l1:       utils.NewList([]int{1, 2, 2, 5}),
			l2:       utils.NewList([]int{1, 3, 6}),
			expected: utils.NewList([]int{1, 1, 2, 2, 3, 5, 6}),
		},
		{
			name:     "merge with single node lists",
			l1:       utils.NewList([]int{10}),
			l2:       utils.NewList([]int{4}),
			expected: utils.NewList([]int{4, 10}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := mergeSortedTwoLists(tc.l1, tc.l2)
			assert.Equal(t, utils.ToSlice(tc.expected), utils.ToSlice(actual))
		})
	}
}
