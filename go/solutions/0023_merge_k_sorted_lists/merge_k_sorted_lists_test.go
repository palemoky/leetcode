package merge_k_sorted_lists

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name:     "example 1",
			input:    [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:     "empty lists slice",
			input:    [][]int{},
			expected: []int{},
		},
		{
			name:     "all lists are nil",
			input:    [][]int{nil, nil, nil},
			expected: []int{},
		},
		{
			name:     "single list",
			input:    [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "with duplicates and negatives",
			input:    [][]int{{-3, -1, 2}, {-2, -2, 2}, {0, 3}},
			expected: []int{-3, -2, -2, -1, 0, 2, 2, 3},
		},
	}

	funcsToTest := map[string]func([]*utils.ListNode) *utils.ListNode{
		"divide-and-conquer": mergeKListsByDivideAndConquer,
		"min-heap":           mergeKListsByMinHeap,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					result := fn(buildLists(tc.input))
					assert.Equal(t, tc.expected, utils.ToSlice(result))
				})
			}
		})
	}
}

func buildLists(vals [][]int) []*utils.ListNode {
	lists := make([]*utils.ListNode, 0, len(vals))
	for _, listVals := range vals {
		lists = append(lists, utils.NewList(listVals))
	}

	return lists
}
