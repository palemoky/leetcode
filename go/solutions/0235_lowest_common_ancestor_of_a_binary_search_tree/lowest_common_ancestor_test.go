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
		{"Ancestor is root", []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 2, 8, 6},
		{"Ancestor is left child", []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 2, 4, 2},
		{"Ancestor is right child", []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 7, 9, 8},
		{"Ancestor is root (distant leaves)", []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 0, 9, 6},
		{"Root is ancestor", []any{2, 1}, 2, 1, 2},
		{"Both nodes are same", []any{2, 1}, 1, 1, 1},
		{"Leftmost and rightmost", []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}, 0, 9, 6},
	}

	funcsToTest := map[string]func(*utils.TreeNode, *utils.TreeNode, *utils.TreeNode) *utils.TreeNode{
		"Iterative": lowestCommonAncestorIterative,
		"Recursive": lowestCommonAncestorRecursive,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					root := utils.BuildTree(tc.input)
					p := findNode(root, tc.p)
					q := findNode(root, tc.q)
					got := fn(root, p, q)
					assert.Equal(t, tc.want, got.Val, "Input: %v, p=%d, q=%d", tc.input, tc.p, tc.q)
				})
			}
		})
	}
}
