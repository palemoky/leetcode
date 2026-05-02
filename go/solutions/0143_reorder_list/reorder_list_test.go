package reorder_list

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func TestReorderList(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "two elements",
			input:    []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "even length",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 4, 2, 3},
		},
		{
			name:     "odd length",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 5, 2, 4, 3},
		},
		{
			name:     "three elements",
			input:    []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			name:     "all equal",
			input:    []int{1, 1, 1, 1},
			expected: []int{1, 1, 1, 1},
		},
	}

	funcsToTest := map[string]func(*utils.ListNode){
		"ReverseHalf": reorderList,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					head := utils.NewList(tc.input)
					fn(head)
					assert.Equal(t, tc.expected, utils.ToSlice(head))
				})
			}
		})
	}
}
