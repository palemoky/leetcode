package binary_tree_level_order_traversal_ii

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrderBottom(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []any
		expected [][]int
	}{
		{
			name:     "示例1：完全二叉树",
			input:    []any{3, 9, 20, nil, nil, 15, 7},
			expected: [][]int{{15, 7}, {9, 20}, {3}},
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
			name:     "左偏树",
			input:    []any{1, 2, nil, 3},
			expected: [][]int{{3}, {2}, {1}},
		},
		{
			name:     "右偏树",
			input:    []any{1, nil, 2, nil, nil, nil, 3},
			expected: [][]int{{3}, {2}, {1}},
		},
		{
			name:     "复杂树",
			input:    []any{1, 2, 3, 4, 5, 6, 7},
			expected: [][]int{{4, 5, 6, 7}, {2, 3}, {1}},
		},
	}

	funcsToTest := map[string]func(*utils.TreeNode) [][]int{
		"levelOrderBottom": levelOrderBottom,
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
