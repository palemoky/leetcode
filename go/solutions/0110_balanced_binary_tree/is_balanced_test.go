package balanced_binary_tree

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func TestIsBalanced(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []any
		want  bool
	}{
		{
			name:  "Empty tree",
			input: []any{},
			want:  true,
		},
		{
			name:  "Single node",
			input: []any{1},
			want:  true,
		},
		{
			name:  "Balanced tree",
			input: []any{3, 9, 20, nil, nil, 15, 7},
			want:  true,
		},
		{
			name:  "Unbalanced tree (left skewed)",
			input: []any{1, 2, nil, 3},
			want:  false,
		},
		{
			name:  "Unbalanced tree (right skewed)",
			input: []any{1, nil, 2, nil, nil, nil, 3},
			want:  false,
		},
		{
			name:  "Perfect binary tree",
			input: []any{1, 2, 3, 4, 5, 6, 7},
			want:  true,
		},
		{
			name:  "Subtree unbalanced",
			input: []any{1, 2, 2, 3, 3, nil, nil, 4, 4},
			want:  false,
		},
		{
			name:  "Left balanced, right unbalanced",
			input: []any{1, 2, 3, nil, nil, 4, nil, nil, nil, nil, nil, 5},
			want:  false,
		},
		{
			name:  "Root unbalanced (height diff = 2)",
			input: []any{1, 2, 3, 4, 5, nil, nil, 6},
			want:  false,
		},
	}

	funcsToTest := map[string]func(*utils.TreeNode) bool{
		"TopDown":  isBalancedTopDown,
		"BottomUp": isBalanced,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					root := utils.BuildTree(tc.input)
					got := fn(root)
					assert.Equal(t, tc.want, got, "Input: %v", tc.input)
				})
			}
		})
	}
}
