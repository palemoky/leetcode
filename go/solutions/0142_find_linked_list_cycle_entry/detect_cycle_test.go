package find_linked_list_cycle_entry

import (
	"leetcode/go/solutions/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCycleEntry(t *testing.T) {
	// --- 准备测试数据 ---
	// 对于查找入口的测试，我们需要获得确切的节点指针作为期望结果

	// Case 1: 无环
	noCycleList := utils.NewCycleList([]int{1, 2, 3}, -1)

	// Case 2: 环指向头部
	cycleToHeadList := utils.NewCycleList([]int{1, 2, 3}, 0)
	expectedHeadEntry := cycleToHeadList // 入口就是头节点

	// Case 3: 环指向中间
	// 我们需要手动构建以获取中间节点的指针
	n1 := &utils.ListNode{Val: 1}
	n2 := &utils.ListNode{Val: 2}
	n3 := &utils.ListNode{Val: 3} // 期望的入口节点
	n4 := &utils.ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n3 // 环指向 n3
	cycleToMiddleList := n1
	expectedMiddleEntry := n3

	testCases := []struct {
		name     string
		input    *utils.ListNode
		expected *utils.ListNode
	}{
		{
			name:     "nil list",
			input:    nil,
			expected: nil,
		},
		{
			name:     "no cycle list",
			input:    noCycleList,
			expected: nil,
		},
		{
			name:     "cycle to head",
			input:    cycleToHeadList,
			expected: expectedHeadEntry,
		},
		{
			name:     "cycle to middle",
			input:    cycleToMiddleList,
			expected: expectedMiddleEntry,
		},
	}

	// 将待测试的函数放入 map 中
	functionsToTest := map[string]func(*utils.ListNode) *utils.ListNode{
		"HashMapMethod":         findCycleEnteryHashMap,
		"MathTwoPointersMethod": findCycleEnteryMathTwoPointers,
	}

	for funcName, findEntryFunc := range functionsToTest {
		t.Run(funcName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					actual := findEntryFunc(tc.input)
					// 检查返回的指针是否与期望的指针指向同一个内存地址
					assert.Same(t, tc.expected, actual)
				})
			}
		})
	}
}
