package validate_binary_search_tree

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func TestIsValidBST(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name  string
		input []any
		want  bool
	}{
		{"Empty tree", []any{}, true},
		{"Single node", []any{1}, true},
		{"Valid BST", []any{2, 1, 3}, true},
		{"Invalid BST", []any{5, 1, 4, nil, nil, 3, 6}, false},
		{"Left skewed valid", []any{3, 2, nil, 1}, true},
		{"Left skewed invalid", []any{3, 2, nil, 9}, false},
		{"Right skewed valid", []any{1, nil, 2, nil, nil, nil, 3}, true},
		{"Right skewed invalid", []any{1, nil, 2, nil, nil, nil, 0}, false},
	}

	funcsToTest := map[string]func(root *utils.TreeNode) bool{
		"Closure": isValidBSTClosure,
		"Pointer": isValidBSTPointer,
		"DFS":     isValidBSTDFS,
		"Array":   isValidBSTArray,
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
