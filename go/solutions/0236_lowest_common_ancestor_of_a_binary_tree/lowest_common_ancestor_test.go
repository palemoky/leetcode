package lowest_common_ancestor_of_a_binary_tree

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func findNode(root *utils.TreeNode, val int) *utils.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := findNode(root.Left, val); left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func TestLowestCommonAncestor(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name  string
		input []any
		p, q  int
		want  int
	}{
		{"LeetCode Example 1", []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}, 5, 1, 3},
		{"LeetCode Example 2", []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}, 5, 4, 5},
		{"Root is ancestor", []any{1, 2}, 1, 2, 1},
		{"Both nodes are same", []any{1, 2}, 2, 2, 2},
		{"Leftmost and rightmost", []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}, 6, 8, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := utils.BuildTree(tc.input)
			p := findNode(root, tc.p)
			q := findNode(root, tc.q)
			got := lowestCommonAncestor(root, p, q)
			assert.Equal(t, tc.want, got.Val, "Input: %v, p=%d, q=%d", tc.input, tc.p, tc.q)
		})
	}
}
