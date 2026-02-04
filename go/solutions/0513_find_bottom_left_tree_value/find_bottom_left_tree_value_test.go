package find_bottom_left_tree_value

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func TestFindBottomLeftValue(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []any
		expected int
	}{
		{
			name:     "示例1：完全二叉树",
			input:    []any{2, 1, 3},
			expected: 1,
		},
		{
			name:     "示例2：不完全树",
			input:    []any{1, 2, 3, 4, nil, 5, 6, 7},
			expected: 7,
		},
		{
			name:     "单节点",
			input:    []any{1},
			expected: 1,
		},
		{
			name:     "左偏树",
			input:    []any{1, 2, nil, 3},
			expected: 3,
		},
		{
			name:     "右偏树",
			input:    []any{1, nil, 2, nil, nil, nil, 3},
			expected: 3,
		},
		{
			name:     "完全二叉树",
			input:    []any{1, 2, 3, 4, 5, 6, 7},
			expected: 4,
		},
		{
			name:     "最后一层只有一个节点",
			input:    []any{1, 2, 3, 4},
			expected: 4,
		},
	}

	funcsToTest := map[string]func(*utils.TreeNode) int{
		"findBottomLeftValue": findBottomLeftValue,
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
