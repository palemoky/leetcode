package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ===== 辅助函数 =====
// 为了方便测试，我们创建一些辅助函数

// createLinkedList 从一个整数切片创建链表
func createLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head
	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}

	return head
}

// toSlice 将链表转换回整数切片，方便断言比较
func toSlice(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var res []int
	current := head
	for current != nil {
		res = append(res, current.Val)
		current = current.Next
	}
	return res
}

// ===== 测试函数 =====

// TestReverse 测试反转链表函数
func TestReverse(t *testing.T) {
	// 定义测试用例的结构体
	type testCase struct {
		name     string    // 测试用例名称
		input    *ListNode // 输入链表
		expected *ListNode // 期望的输出链表
	}

	// 定义测试用例表
	testCases := []testCase{
		{
			name:     "nil list",
			input:    nil,
			expected: nil,
		},
		{
			name:     "single node list",
			input:    createLinkedList([]int{1}),
			expected: createLinkedList([]int{1}),
		},
		{
			name:     "two nodes list",
			input:    createLinkedList([]int{1, 2}),
			expected: createLinkedList([]int{2, 1}),
		},
		{
			name:     "multiple nodes list (odd)",
			input:    createLinkedList([]int{1, 2, 3, 4, 5}),
			expected: createLinkedList([]int{5, 4, 3, 2, 1}),
		},
		{
			name:     "multiple nodes list (even)",
			input:    createLinkedList([]int{1, 2, 3, 4}),
			expected: createLinkedList([]int{4, 3, 2, 1}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := reverse(tc.input)
			assert.Equal(t, toSlice(tc.expected), toSlice(actual))
		})
	}
}

// TestFindMiddle 统一测试两个查找中间节点的函数
// 因为它们的行为和预期结果应该完全一样
func TestFindMiddle(t *testing.T) {
	// 为测试用例构建链表，并提前获取中间节点的指针
	list1 := createLinkedList([]int{1, 2, 3, 4, 5}) // 中间节点是 3 (l.Next.Next)
	list2 := createLinkedList([]int{1, 2, 3, 4})    // 中间节点是 3 (l.Next.Next) - 按实现逻辑，偶数取第二个
	list3 := createLinkedList([]int{1})             // 中间节点是 1 (l)
	list4 := createLinkedList([]int{1, 2})          // 中间节点是 2 (l.Next)

	// 定义测试用例的结构体
	type testCase struct {
		name     string    // 测试用例名称
		input    *ListNode // 输入链表
		expected *ListNode // 期望的中间节点指针
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
	functionsToTest := map[string]func(*ListNode) *ListNode{
		"ArrayMethod":       findMiddleArray,
		"TwoPointersMethod": findMiddleTwoPointers,
	}

	for funcName, findMiddleFunc := range functionsToTest {
		t.Run(funcName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
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

func TestMergeSortedTwoLists(t *testing.T) {
	// 定义测试用例表，覆盖各种情况
	testCases := []struct {
		name     string
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{
			name:     "both lists are nil",
			l1:       nil,
			l2:       nil,
			expected: nil,
		},
		{
			name:     "l1 is nil, l2 is not",
			l1:       nil,
			l2:       createLinkedList([]int{1, 2, 3}),
			expected: createLinkedList([]int{1, 2, 3}),
		},
		{
			name:     "l2 is nil, l1 is not",
			l1:       createLinkedList([]int{4, 5, 6}),
			l2:       nil,
			expected: createLinkedList([]int{4, 5, 6}),
		},
		{
			name:     "standard merge with interleaved values",
			l1:       createLinkedList([]int{1, 3, 5}),
			l2:       createLinkedList([]int{2, 4, 6}),
			expected: createLinkedList([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name:     "merge with different lengths (l1 shorter)",
			l1:       createLinkedList([]int{1, 5}),
			l2:       createLinkedList([]int{2, 3, 4}),
			expected: createLinkedList([]int{1, 2, 3, 4, 5}),
		},
		{
			name:     "merge with different lengths (l2 shorter)",
			l1:       createLinkedList([]int{1, 2, 10}),
			l2:       createLinkedList([]int{5, 6}),
			expected: createLinkedList([]int{1, 2, 5, 6, 10}),
		},
		{
			name:     "l1 values are all smaller than l2",
			l1:       createLinkedList([]int{1, 2, 3}),
			l2:       createLinkedList([]int{4, 5, 6}),
			expected: createLinkedList([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name:     "merge with duplicate values",
			l1:       createLinkedList([]int{1, 2, 2, 5}),
			l2:       createLinkedList([]int{1, 3, 6}),
			expected: createLinkedList([]int{1, 1, 2, 2, 3, 5, 6}),
		},
		{
			name:     "merge with single node lists",
			l1:       createLinkedList([]int{10}),
			l2:       createLinkedList([]int{4}),
			expected: createLinkedList([]int{4, 10}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := mergeSortedTwoLists(tc.l1, tc.l2)
			assert.Equal(t, toSlice(tc.expected), toSlice(actual))
		})
	}
}
