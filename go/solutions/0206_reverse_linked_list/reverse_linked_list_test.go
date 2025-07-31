package reverse_linked_list

import (
	"leetcode/go/solutions/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name     string
		input    *utils.ListNode
		expected *utils.ListNode
	}

	testCases := []testCase{
		{
			name:     "nil list",
			input:    nil,
			expected: nil,
		},
		{
			name:     "single node list",
			input:    utils.NewList([]int{1}),
			expected: utils.NewList([]int{1}),
		},
		{
			name:     "two nodes list",
			input:    utils.NewList([]int{1, 2}),
			expected: utils.NewList([]int{2, 1}),
		},
		{
			name:     "multiple nodes list (odd)",
			input:    utils.NewList([]int{1, 2, 3, 4, 5}),
			expected: utils.NewList([]int{5, 4, 3, 2, 1}),
		},
		{
			name:     "multiple nodes list (even)",
			input:    utils.NewList([]int{1, 2, 3, 4}),
			expected: utils.NewList([]int{4, 3, 2, 1}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := reverseList(tc.input)
			assert.Equal(t, utils.ToSlice(tc.expected), utils.ToSlice(actual))
		})
	}
}
