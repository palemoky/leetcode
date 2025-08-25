package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name  string
		input []any
		want  []int // 层序遍历结果
	}{
		{
			name:  "Empty input",
			input: []any{},
			want:  []int{},
		},
		{
			name:  "Single node",
			input: []any{1},
			want:  []int{1},
		},
		{
			name:  "Full tree",
			input: []any{1, 2, 3, 4, 5, 6, 7},
			want:  []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:  "Tree with nils",
			input: []any{1, nil, 2, nil, nil, 3},
			want:  []int{1, 2, 3}, // 层序遍历只收集非 nil 节点
		},
		{
			name:  "Left skewed",
			input: []any{1, 2, nil, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "Right skewed",
			input: []any{1, nil, 2, nil, nil, nil, 3},
			want:  []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := levelOrder(tc.input)
			assert.Equal(t, tc.want, got, "Input: %v", tc.input)
		})
	}
}

// 层序遍历辅助函数，只收集非 nil 节点值
func levelOrder(vals []any) []int {
	root := BuildTree(vals)
	if root == nil {
		return []int{}
	}
	res := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}
