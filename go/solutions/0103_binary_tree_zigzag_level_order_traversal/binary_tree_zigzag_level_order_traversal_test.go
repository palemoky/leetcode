package binary_tree_zigzag_level_order_traversal

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func TestZigzagLevelOrder(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []any
		expected [][]int
	}{
		{
			name:     "示例1：完全二叉树",
			input:    []any{3, 9, 20, nil, nil, 15, 7},
			expected: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			name:     "示例2：单节点",
			input:    []any{1},
			expected: [][]int{{1}},
		},
		{
			name:     "空树",
			input:    []any{},
			expected: [][]int{},
		},
		{
			name:     "两层树",
			input:    []any{1, 2, 3},
			expected: [][]int{{1}, {3, 2}},
		},
		{
			name:     "三层完全二叉树",
			input:    []any{1, 2, 3, 4, 5, 6, 7},
			expected: [][]int{{1}, {3, 2}, {4, 5, 6, 7}},
		},
		{
			name:     "左偏树",
			input:    []any{1, 2, nil, 3},
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "四层树",
			input:    []any{1, 2, 3, 4, nil, nil, 5},
			expected: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}

	funcsToTest := map[string]func(*utils.TreeNode) [][]int{
		"zigzagLevelOrder": zigzagLevelOrder,
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
