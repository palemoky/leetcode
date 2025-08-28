package lowest_common_ancestor_of_a_binary_search_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"leetcode/go/solutions/utils"
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
		{"LeetCode Example 1", []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 2, 8, 6},
		{"LeetCode Example 2", []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 2, 4, 2},
		{"Root is ancestor", []any{2, 1}, 2, 1, 2},
		{"Both nodes are same", []any{2, 1}, 1, 1, 1},
		{"Leftmost and rightmost", []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 0, 9, 6},
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
