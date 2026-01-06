package invert_binary_tree

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func levelOrder(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := []int{}
	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}

		nums = append(nums, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return nums
}

func TestInvertTree(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []any
		want  []int
	}{
		{"Empty", []any{}, []int{}},
		{"Single", []any{1}, []int{1}},
		{"Example", []any{4, 2, 7, 1, 3, 6, 9}, []int{4, 7, 2, 9, 6, 3, 1}},
		{"Left skewed", []any{1, 2, nil, 3}, []int{1, 2, 3}},
		{"Right skewed", []any{1, nil, 2, nil, nil, nil, 3}, []int{1, 2, 3}},
	}

	funcsToTest := map[string]func(root *utils.TreeNode) *utils.TreeNode{
		"Postorder": invertTreePostorder,
		"Preorder":  invertTreePreorder,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					root := utils.BuildTree(tc.input)
					got := fn(root)
					assert.Equal(t, tc.want, levelOrder(got))
				})
			}
		})
	}
}
