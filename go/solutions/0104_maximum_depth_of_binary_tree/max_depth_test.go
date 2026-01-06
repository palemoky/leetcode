package maximum_depth_of_binary_tree

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input []any
		want  int
	}{
		{
			name:  "Empty tree",
			input: []any{},
			want:  0,
		},
		{
			name:  "Single node",
			input: []any{1},
			want:  1,
		},
		{
			name:  "Full tree",
			input: []any{3, 9, 20, nil, nil, 15, 7},
			want:  3,
		},
		{
			name:  "Left skewed",
			input: []any{1, 2, nil, 3, nil, nil, nil, 4},
			want:  4,
		},
		{
			name:  "Right skewed",
			input: []any{1, nil, 2, nil, nil, nil, 3},
			want:  3,
		},
		{
			name:  "Balanced tree",
			input: []any{1, 2, 3, 4, 5, 6, 7},
			want:  3,
		},
	}

	funcsToTest := map[string]func(root *utils.TreeNode) int{
		"BFS": maxDepthBFS,
		"DFS": maxDepthDFS,
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
