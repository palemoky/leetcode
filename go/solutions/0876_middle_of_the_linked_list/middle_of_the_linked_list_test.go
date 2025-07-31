package middle_of_the_linked_list

import (
	"leetcode/go/solutions/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFindMiddle 统一测试两个查找中间节点的函数
// 因为它们的行为和预期结果应该完全一样
func TestFindMiddle(t *testing.T) {
	t.Parallel()

	// 为测试用例构建链表，并提前获取中间节点的指针
	list1 := utils.NewList([]int{1, 2, 3, 4, 5}) // 中间节点是 3 (l.Next.Next)
	list2 := utils.NewList([]int{1, 2, 3, 4})    // 中间节点是 3 (l.Next.Next) - 按实现逻辑，偶数取第二个
	list3 := utils.NewList([]int{1})             // 中间节点是 1 (l)
	list4 := utils.NewList([]int{1, 2})          // 中间节点是 2 (l.Next)

	type testCase struct {
		name     string          // 测试用例名称
		input    *utils.ListNode // 输入链表
		expected *utils.ListNode // 期望的中间节点指针
	}

	testCases := []testCase{
		{
			name:     "nil list",
			input:    nil,
			expected: nil,
		},
		{
			name:     "single node list",
			input:    list3,
			expected: list3, // 期望返回节点1本身
		},
		{
			name:     "odd length list",
			input:    list1,
			expected: list1.Next.Next, // 期望返回节点3
		},
		{
			name:     "even length list",
			input:    list2,
			expected: list2.Next.Next, // 期望返回节点3
		},
		{
			name:     "two nodes list",
			input:    list4,
			expected: list4.Next, // 期望返回节点2
		},
	}

	// 将待测试的函数放入 map 中，方便统一调用
	functionsToTest := map[string]func(*utils.ListNode) *utils.ListNode{
		"ArrayMethod":       findMiddleArray,
		"TwoPointersMethod": findMiddleTwoPointers,
	}

	for funcName, findMiddleFunc := range functionsToTest {
		t.Run(funcName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()
					actual := findMiddleFunc(tc.input)

					// 对于查找中间节点，我们期望返回的是原链表中的某个节点
					// 因此，我们应该比较它们的指针地址是否相同
					// assert.Same() 用于检查两个指针是否指向同一个对象
					assert.Same(t, tc.expected, actual)
				})
			}
		})
	}
}
