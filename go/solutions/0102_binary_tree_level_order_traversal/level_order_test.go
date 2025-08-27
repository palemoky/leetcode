package binary_tree_level_order_traversal

import (
	"leetcode/go/solutions/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []any
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
			root := utils.BuildTree(tc.input)
			got := levelOrder(root)
			assert.Equal(t, tc.expected, got)
		})
	}
}
