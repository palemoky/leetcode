package binary_tree_maximum_path_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"leetcode/go/solutions/utils"
)

func TestMaxPathSum(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *utils.TreeNode
		expected int
	}{
		{
			name: "简单路径",
			input: &utils.TreeNode{
				Val: 1,
				Left: &utils.TreeNode{
					Val: 2,
				},
				Right: &utils.TreeNode{
					Val: 3,
				},
			},
			expected: 6, // 2 -> 1 -> 3
		},
		{
			name: "复杂路径",
			input: &utils.TreeNode{
				Val: -10,
				Left: &utils.TreeNode{
					Val: 9,
				},
				Right: &utils.TreeNode{
					Val: 20,
					Left: &utils.TreeNode{
						Val: 15,
					},
					Right: &utils.TreeNode{
						Val: 7,
					},
				},
			},
			expected: 42, // 15 -> 20 -> 7
		},
		{
			name: "单节点",
			input: &utils.TreeNode{
				Val: -3,
			},
			expected: -3,
		},
		{
			name: "全负数",
			input: &utils.TreeNode{
				Val: -2,
				Left: &utils.TreeNode{
					Val: -1,
				},
			},
			expected: -1, // 只选择 -1
		},
		{
			name: "左子树路径最大",
			input: &utils.TreeNode{
				Val: 5,
				Left: &utils.TreeNode{
					Val: 4,
					Left: &utils.TreeNode{
						Val: 11,
						Left: &utils.TreeNode{
							Val: 7,
						},
						Right: &utils.TreeNode{
							Val: 2,
						},
					},
				},
				Right: &utils.TreeNode{
					Val: 8,
					Left: &utils.TreeNode{
						Val: 13,
					},
					Right: &utils.TreeNode{
						Val: 4,
						Right: &utils.TreeNode{
							Val: 1,
						},
					},
				},
			},
			expected: 48, // 7 -> 11 -> 4 -> 5 -> 8 -> 13
		},
	}

	funcsToTest := map[string]func(*utils.TreeNode) int{
		"maxPathSum": maxPathSum,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					result := fn(tc.input)
					assert.Equal(t, tc.expected, result)
				})
			}
		})
	}
}
