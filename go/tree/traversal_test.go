package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		vals     []any
		expected [][]int
	}{
		{"Empty tree", []any{}, [][]int{}},
		{"Single node", []any{1}, [][]int{{1}}},
		{"Full tree", []any{1, 2, 3, 4, 5, 6, 7}, [][]int{{1}, {2, 3}, {4, 5, 6, 7}}},
		{"Tree with nils", []any{1, 2, 3, nil, 4}, [][]int{{1}, {2, 3}, {4}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			root := buildTree(tc.vals)
			got := levelOrder(root)
			assert.Equal(t, tc.expected, got)
		})
	}
}

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

	funcsToTest := map[string]func(root *TreeNode) []int{
		"IterativeStack": preorderIterativeStack,
		"Iterative":      preorderIterative,
		"Recursive":      preorderRecursive,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					root := buildTree(tc.vals)
					got := fn(root)
					assert.Equal(t, tc.expected, got)
				})
			}
		})
	}
}

func TestInorderTraversal(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		vals     []any
		expected []int
	}{
		{"Empty tree", []any{}, []int{}},
		{"Single node", []any{1}, []int{1}},
		{"Full tree", []any{1, 2, 3, 4, 5, 6, 7}, []int{4, 2, 5, 1, 6, 3, 7}},
		{"Tree with nils", []any{1, 2, 3, nil, 4}, []int{2, 4, 1, 3}},
	}

	funcsToTest := map[string]func(root *TreeNode) []int{
		"Iterative":          inorderIterative,
		"IterativeWithColor": inorderIterativeWithColor,
		"Recursive":          inorderRecursive,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					root := buildTree(tc.vals)
					got := fn(root)
					assert.Equal(t, tc.expected, got)
				})
			}
		})
	}
}

func TestPostorderTraversal(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		vals     []any
		expected []int
	}{
		{"Empty tree", []any{}, []int{}},
		{"Single node", []any{1}, []int{1}},
		{"Full tree", []any{1, 2, 3, 4, 5, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}},
		{"Tree with nils", []any{1, 2, 3, nil, 4}, []int{4, 2, 3, 1}},
	}

	funcsToTest := map[string]func(root *TreeNode) []int{
		"Iterative":              postorderIterative,
		"Iterative with Reverse": postorderTraversalWithReverse,
		"Recursive":              postorderRecursive,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					root := buildTree(tc.vals)
					got := fn(root)
					assert.Equal(t, tc.expected, got)
				})
			}
		})
	}
}
