package binary_tree_right_side_view

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func TestRightSideView(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []any
		expected []int
	}{
		{
			name:     "示例1：完全二叉树",
			input:    []any{1, 2, 3, nil, 5, nil, 4},
			expected: []int{1, 3, 4},
		},
		{
			name:     "示例2：右偏树",
			input:    []any{1, nil, 3},
			expected: []int{1, 3},
		},
		{
			name:     "空树",
			input:    []any{},
			expected: []int{},
		},
		{
			name:     "单节点",
			input:    []any{1},
			expected: []int{1},
		},
		{
			name:     "左偏树",
			input:    []any{1, 2, nil, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "完全二叉树",
			input:    []any{1, 2, 3, 4, 5, 6, 7},
			expected: []int{1, 3, 7},
		},
		{
			name:     "不完全树",
			input:    []any{1, 2, 3, nil, 5, nil, 4, nil, nil, 6},
			expected: []int{1, 3, 4, 6},
		},
	}

	funcsToTest := map[string]func(*utils.TreeNode) []int{
		"rightSideView": rightSideView,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					root := utils.BuildTree(tc.input)
					result := fn(root)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
