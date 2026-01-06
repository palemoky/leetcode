package remove_nth_node_from_end_of_list

import (
	"testing"

	"leetcode/go/solutions/utils"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEndTwoPointers(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string // 测试用例的描述
		inputList    []int  // 输入的链表（用切片表示）
		n            int    // 要移除的倒数第 n 个节点
		expectedList []int  // 期望的结果链表（用切片表示）
	}{
		{
			name:         "Remove from the middle",
			inputList:    []int{1, 2, 3, 4, 5},
			n:            2, // 移除倒数第2个节点 (4)
			expectedList: []int{1, 2, 3, 5},
		},
		{
			name:         "Remove the head of the list",
			inputList:    []int{1, 2, 3, 4, 5},
			n:            5, // 移除倒数第5个节点 (1)
			expectedList: []int{2, 3, 4, 5},
		},
		{
			name:         "Remove the tail of the list",
			inputList:    []int{1, 2, 3, 4, 5},
			n:            1, // 移除倒数第1个节点 (5)
			expectedList: []int{1, 2, 3, 4},
		},
		{
			name:         "List with only one node",
			inputList:    []int{1},
			n:            1,       // 移除倒数第1个节点 (1)
			expectedList: []int{}, // 结果为空链表
		},
		{
			name:         "List with two nodes, remove head",
			inputList:    []int{1, 2},
			n:            2, // 移除倒数第2个节点 (1)
			expectedList: []int{2},
		},
		{
			name:         "List with two nodes, remove tail",
			inputList:    []int{1, 2},
			n:            1, // 移除倒数第1个节点 (2)
			expectedList: []int{1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			head := utils.NewList(tc.inputList)
			resultHead := removeNthFromEndTwoPointers(head, tc.n)
			resultSlice := utils.ToSlice(resultHead)
			assert.Equal(t, tc.expectedList, resultSlice, "The resulting linked list should match the expected one.")
		})
	}
}
