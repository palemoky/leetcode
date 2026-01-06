package symmetric_tree

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func TestIsSymmetric(t *testing.T) {
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
			name:  "Symmetric tree",
			input: []any{1, 2, 2, 3, 4, 4, 3},
			want:  true,
		},
		{
			name:  "Asymmetric tree",
			input: []any{1, 2, 2, nil, 3, nil, 3},
			want:  false,
		},
		{
			name:  "Left only",
			input: []any{1, 2, nil},
			want:  false,
		},
		{
			name:  "Right only",
			input: []any{1, nil, 2},
			want:  false,
		},
	}

	funcsToTest := map[string]func(root *utils.TreeNode) bool{
		"LevelOrder": isSymmetricLevelOrder,
		"TwoQueues":  isSymmetricTwoQueues,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					root := utils.BuildTree(tc.input)
					got := fn(root)
					assert.Equal(t, tc.want, got, "Input: %v", tc.input)
				})
			}
		})
	}
}
