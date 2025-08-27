package binary_tree_inorder_traversal

import (
	"leetcode/go/solutions/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []any
		expected []int
	}{
		{"Empty tree", []any{}, []int{}},
		{"Single node", []any{1}, []int{1}},
		{"Full tree", []any{1, 2, 3, 4, 5, 6, 7}, []int{4, 2, 5, 1, 6, 3, 7}},
		{"Tree with nils", []any{1, 2, 3, nil, 4}, []int{2, 4, 1, 3}},
	}

	funcsToTest := map[string]func(root *utils.TreeNode) []int{
		"Iterative":          inorderIterative,
		"IterativeWithColor": inorderIterativeWithColor,
		"Recursive":          inorderRecursive,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					root := utils.BuildTree(tc.input)
					got := fn(root)
					assert.Equal(t, tc.expected, got)
				})
			}
		})
	}
}
