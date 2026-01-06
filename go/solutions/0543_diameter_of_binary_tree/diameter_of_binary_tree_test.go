package diameter_of_binary_tree

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name  string
		input []any
		want  int
	}{
		{"LeetCode Example 1", []any{1, 2, 3, 4, 5}, 3},
		{"Single node", []any{1}, 0},
		{"Empty tree", []any{}, 0},
		{"Left skewed", []any{1, 2, nil, 3, nil, nil, nil, 4}, 3},
		{"Right skewed", []any{1, nil, 2, nil, nil, nil, 3}, 2},
		{"Balanced tree", []any{1, 2, 3, 4, 5, 6, 7}, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := utils.BuildTree(tc.input)
			got := diameterOfBinaryTree(root)
			assert.Equal(t, tc.want, got, "Input: %v", tc.input)
		})
	}
}
