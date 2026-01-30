package reverse_nodes_in_k_group

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func TestReverseKGroup(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		head     []int
		k        int
		expected []int
	}{
		{name: "Example 1: k=2", head: []int{1, 2, 3, 4, 5}, k: 2, expected: []int{2, 1, 4, 3, 5}},
		{name: "Example 2: k=3", head: []int{1, 2, 3, 4, 5}, k: 3, expected: []int{3, 2, 1, 4, 5}},
		{name: "k=1 不反转", head: []int{1, 2, 3, 4, 5}, k: 1, expected: []int{1, 2, 3, 4, 5}},
		{name: "k 等于链表长度", head: []int{1, 2, 3, 4, 5}, k: 5, expected: []int{5, 4, 3, 2, 1}},
		{name: "k 大于链表长度", head: []int{1, 2, 3}, k: 5, expected: []int{1, 2, 3}},
		{name: "单节点链表", head: []int{1}, k: 1, expected: []int{1}},
		{name: "两个节点 k=2", head: []int{1, 2}, k: 2, expected: []int{2, 1}},
		{name: "完整分组", head: []int{1, 2, 3, 4, 5, 6}, k: 3, expected: []int{3, 2, 1, 6, 5, 4}},
		{name: "最后一组不足 k 个", head: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, expected: []int{3, 2, 1, 6, 5, 4, 7}},
	}

	funcsToTest := map[string]func(*utils.ListNode, int) *utils.ListNode{
		"reverseKGroup (迭代)":            reverseKGroup,
		"reverseKGroupByRecursion (递归)": reverseKGroupByRecursion,
	}

	for fnName, fn := range funcsToTest {
		t.Run(fnName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					head := utils.NewList(tc.head)
					result := fn(head, tc.k)
					resultArray := utils.ToSlice(result)
					assert.Equal(t, tc.expected, resultArray)
				})
			}
		})
	}
}
