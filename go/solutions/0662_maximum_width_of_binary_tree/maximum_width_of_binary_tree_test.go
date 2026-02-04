package maximum_width_of_binary_tree

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func TestWidthOfBinaryTree(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []any
		expected int
	}{
		{
			name:     "示例1：包含空节点",
			input:    []any{1, 3, 2, 5, 3, nil, 9},
			expected: 4,
		},
		{
			name:     "示例2：只有左子树",
			input:    []any{1, 3, 2, 5, nil, nil, 9, 6, nil, 7},
			expected: 4, // 最后一层: [6, 7]，索引 0 到 3，宽度 = 4
		},
		{
			name:     "示例3：只有右子树",
			input:    []any{1, 3, 2, 5},
			expected: 2,
		},
		{
			name:     "空树",
			input:    []any{},
			expected: 0,
		},
		{
			name:     "单节点",
			input:    []any{1},
			expected: 1,
		},
		{
			name:     "完全二叉树",
			input:    []any{1, 2, 3, 4, 5, 6, 7},
			expected: 4,
		},
		{
			name:     "左偏树",
			input:    []any{1, 2, nil, 3},
			expected: 1, // 每层只有一个节点
		},
		{
			name:     "右偏树",
			input:    []any{1, nil, 2, nil, nil, nil, 3},
			expected: 1, // 每层只有一个节点
		},
	}

	funcsToTest := map[string]func(*utils.TreeNode) int{
		"widthOfBinaryTree": widthOfBinaryTree,
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
