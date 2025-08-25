package binary_tree_preorder_traversal

import (
	"leetcode/go/solutions/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreorderTraversal(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		vals     []any
		expected []int
	}{
		{"Empty tree", []any{}, []int{}},
		{"Single node", []any{1}, []int{1}},
		{"Full tree", []any{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 4, 5, 3, 6, 7}},
		{"Tree with nils", []any{1, 2, 3, nil, 4}, []int{1, 2, 4, 3}},
	}

	funcsToTest := map[string]func(root *utils.TreeNode) []int{
		"Recursive":          preorderRecursive,
		"IterativeStack":     preorderIterativeStack,
		"Iterative":          preorderIterative,
		"IterativeWithColor": preorderIterativeWithColor,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					root := utils.BuildTree(tc.vals)
					got := fn(root)
					assert.Equal(t, tc.expected, got)
				})
			}
		})
	}
}
