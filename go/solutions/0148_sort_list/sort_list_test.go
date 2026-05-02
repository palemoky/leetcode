package sort_list

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func TestSortList(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "nil",
			input:    nil,
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "unsorted",
			input:    []int{4, 2, 1, 3},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "duplicates",
			input:    []int{3, 1, 2, 1, 3},
			expected: []int{1, 1, 2, 3, 3},
		},
		{
			name:     "all equal",
			input:    []int{2, 2, 2},
			expected: []int{2, 2, 2},
		},
		{
			name:     "two elements",
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
	}

	funcsToTest := map[string]func(*utils.ListNode) *utils.ListNode{
		"MergeSort": sortList,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					actual := fn(utils.NewList(tc.input))
					assert.Equal(t, tc.expected, utils.ToSlice(actual))
				})
			}
		})
	}
}
