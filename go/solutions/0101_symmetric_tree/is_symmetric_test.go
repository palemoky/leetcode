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
		{"Empty tree", []any{}, true},
		{"Single node", []any{1}, true},
		{"Symmetric tree", []any{1, 2, 2, 3, 4, 4, 3}, true},
		{"Asymmetric tree", []any{1, 2, 2, nil, 3, nil, 3}, false},
		{"Left only", []any{1, 2, nil}, false},
		{"Right only", []any{1, nil, 2}, false},
	}

	funcsToTest := map[string]func(root *utils.TreeNode) bool{
		"MirrorRecursive": isSymmetricMirrorRecursive,
		"LevelOrder":      isSymmetricLevelOrder,
		"TwoQueues":       isSymmetricTwoQueues,
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
