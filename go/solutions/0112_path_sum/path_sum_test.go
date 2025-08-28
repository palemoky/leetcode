package path_sum

import (
	"leetcode/go/solutions/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPathSum(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name   string
		input  []any
		target int
		want   bool
	}{
		{"Typical path sum exists", []any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}, 22, true},
		{"No path sum", []any{1, 2, 3}, 5, false},
		{"Single node equals target", []any{1}, 1, true},
		{"Single node not equals target", []any{1}, 2, false},
		{"Empty tree", []any{}, 0, false},
		{"Negative numbers", []any{-2, nil, -3}, -5, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := utils.BuildTree(tc.input)
			got := hasPathSum(root, tc.target)
			assert.Equal(t, tc.want, got, "Input: %v, target=%d", tc.input, tc.target)
		})
	}
}
